package semaphoreci

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type Commit struct {
	Id  string
	Url string
}

type Build struct {
	Build_Url    string
	Build_Number int
	Result       string
	Started_At   string
	Finihshed_At string
	Commit       Commit
}

func (b *Branch) getBuild(commitId string) (*Build, error) {
	for _, build := range b.Builds {
		if build.Commit.Id == commitId {
			return &build, nil
		}
	}

	return nil, errors.New(commitId + " can't be found")
}

type Branch struct {
	Branch_Name        string
	Result             string
	Branch_History_Url string
	Builds             []Build
}

type Project struct {
	Name     string
	Branches []Branch
}

type Semaphoreci struct {
  AuthToken string
}

func (s *Semaphoreci) projectURL() string {
  return "https://semaphoreapp.com/api/v1/projects?auth_token=" + s.AuthToken
}

func (s *Semaphoreci) GetProject(projectName string) (*Project, error) {
	resp, err := http.Get(s.projectURL())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	var projects []Project
	for {
		if err := dec.Decode(&projects); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}

	for _, project := range projects {
		if project.Name == projectName {
			return &project, nil
		}
	}

	return nil, errors.New(projectName + " can't be found")
}

func (s *Semaphoreci) GetBranch(projectName, branchName string) (*Branch, error) {
	project, err := s.GetProject(projectName)
	if err != nil {
		log.Fatal(err)
	}

	for _, branch := range project.Branches {
		if branch.Branch_Name == branchName {
			resp, err := http.Get(branch.Branch_History_Url)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()

			dec := json.NewDecoder(resp.Body)
			var branchDetail Branch
			for {
				if err := dec.Decode(&branchDetail); err == io.EOF {
					break
				} else if err != nil {
					log.Fatal(err)
				}
			}

			return &branchDetail, nil
		}
	}

	return nil, errors.New(branchName + " can't be found")
}
