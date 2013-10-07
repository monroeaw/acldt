package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"regexp"
	"strconv"
	"time"
)

var gitTagbuild = cli.Command{
	Name:      "git-tagbuild",
	ShortName: "gt",
	Usage:     "Runs Git tag in the format of YYYY-MM-DD-NN",
	Description: `Run Git Git tag in the format of YYYY-MM-DD-NN, where NN is an
incremental number for the build of the day. For example,

  $ acldt git-tagbuild

generates a tag named 2013-10-06-01 with commit message "Release 2013-10-06-01".
The tag will be pushed to remote as well`,
	Action: gitTagbuildAction,
}

func gitTagbuildAction(c *cli.Context) {
	execCmd("git fetch --tags")

	tagName := buildTagName()
	tagCmd := fmt.Sprintf(`git tag -a %s -m "Release %s"`, tagName, tagName)
	pushTagCmd := fmt.Sprintf(`git push origin %s`, tagName)
	execCmd(tagCmd)
	execCmd(pushTagCmd)
}

func buildTagName() string {
	date := getDateString()
	tags := execCmdOutput("git tag")
	buildNumber := increaseBuildNumber(tags, date)
	return fmt.Sprintf("%s-%s", date, formatNumber(buildNumber))
}

func getDateString() string {
	const layout = "2006-01-02"
	t := time.Now()
	return t.Format(layout)
}

func increaseBuildNumber(tags []string, date string) int {
	rs := fmt.Sprintf("%s-(\\d+)", date)
	r := regexp.MustCompile(rs)
	var buildNumber int
	for _, tag := range tags {
		if r.MatchString(tag) {
			n, _ := strconv.Atoi(r.FindStringSubmatch(tag)[1])
			if n > buildNumber {
				buildNumber = n
			}
		}
	}

	return buildNumber + 1
}

func formatNumber(n int) string {
	result := strconv.Itoa(n)
	if n < 10 && n > 0 {
		result = "0" + result
	}

	return result
}
