package semaphoreci

import (
	"github.com/bmizerany/assert"
	"testing"
)

var authToken = "H3ESLzEwhpdgjqv9RQjW"
var semaphoreci = Semaphoreci{authToken}

func TestBranchGetBuild(t *testing.T) {
	commit := Commit{Id: "sha"}
	build := Build{Commit: commit, Build_Number: 123}
	branch := Branch{Builds: []Build{build}}

	result, _ := branch.getBuild("sha")
	assert.Equal(t, result.Build_Number, build.Build_Number)
}

func TestFetchProjectByName(t *testing.T) {
	project, _ := semaphoreci.GetProject("accounts")
	assert.Equal(t, project.Name, "accounts")
	assert.T(t, len(project.Branches) > 0)

	_, err := semaphoreci.GetProject("accounts1")
	assert.NotEqual(t, err, nil)
}

func TestFetchBranchByName(t *testing.T) {
	branch, _ := semaphoreci.GetBranch("accounts", "master")
	assert.Equal(t, branch.Branch_Name, "master")
	assert.T(t, len(branch.Builds) > 0)

	_, err := semaphoreci.GetBranch("accounts", "master1")
	assert.NotEqual(t, err, nil)
}
