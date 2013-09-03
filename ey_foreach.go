package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"regexp"
	"strings"
)

var eyForeachCmd = cli.Command{
	Name:      "ey-foreach",
	ShortName: "ef",
	Usage:     "Applies action for each EY environment",
	Description: `Applies action for each Engineyard environment. For example,
to upload recipes for each production environment for the app Projects:

  $ acldt ey-foreach -a projects -e production recipes upload`,
	Flags: []cli.Flag{
		cli.StringFlag{Name: "a", Usage: "app name on EY, e.g., projects"},
		cli.StringFlag{Name: "e", Usage: "env name on EY, e.g., production"},
	},
	Action: eyForeachAction,
}

func eyForeachAction(c *cli.Context) {
	if c.String("a") == "" {
		fmt.Println("Missing app name")
		cli.ShowCommandHelp(c, "ef")
		return
	}

	if c.String("e") == "" {
		fmt.Println("Missing env name")
		cli.ShowCommandHelp(c, "ef")
		return
	}

	listEnvsCmd := fmt.Sprintf("ey environments --all --account=ACL --simple")
	envs := execCmdOutput(listEnvsCmd)
	filteredEnvs := filterEnvs(envs, c.String("a"), c.String("e"))

	for _, env := range filteredEnvs {
		eyCmd := fmt.Sprintf("ey %s -e %s", strings.Join(c.Args(), " "), env)
		fmt.Printf("Running: %s\n", eyCmd)
		execCmd(eyCmd)
	}
}

func filterEnvs(envs []string, appName, envName string) []string {
	regexpString := fmt.Sprintf("%s.*_%s", appName, envName)
	r, err := regexp.Compile(regexpString)
	if err != nil {
		log.Fatal(err)
	}

	filteredEnvs := []string{}
	for _, env := range envs {
		if r.MatchString(env) {
			filteredEnvs = append(filteredEnvs, env)
		}
	}

	return filteredEnvs
}
