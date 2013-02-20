package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"
)

var cmdGitRmerge = &Command{
	Usage: "git:rmerge <branch>",
	Short: "run Git rebase and Git merge with --no-ff",
	Long: `
Run Git rebase on a branch and then run Git merge with no fast forward
(git merge --no-ff).

As an example, assuming current branch is master, running this command
rebases a list of topic branches on top of master and then merge them
into master with no fast forward.

  $ acldt git:rmerge topic1 topic2
`,
}

func init() {
	cmdGitRmerge.Run = runGitRmerge
}

func runGitRmerge(cmd *Command, args []string) {
	if len(args) == 0 {
		cmd.printUsage()
		return
	}

	baseBranch := getBaseBranch()
	for _, arg := range args {
		topicBranch := strings.TrimSpace(arg)

		execCmd("git fetch")

		execCmd("git checkout " + topicBranch)
		if hasRemoteBranch(topicBranch) {
			execCmd("git pull origin " + topicBranch)
		}
		execCmd("git rebase -i origin/" + baseBranch)
		execCmd("git push origin HEAD -f")

		execCmd("git checkout " + baseBranch)
		execCmd("git pull origin " + baseBranch)
		execCmd("git merge " + topicBranch + " --no-ff")
		execCmd("git push origin HEAD")

		execCmd("git branch -d " + topicBranch)
		execCmd("git push origin :" + topicBranch)
	}
}

func getBaseBranch() string {
	out, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(string(out))
}

func hasRemoteBranch(branch string) bool {
	out, err := exec.Command("git", "branch", "-r").Output()
	if err != nil {
		return false
	}
	for _, line := range bytes.Split(out, []byte{'\n'}) {
		if strings.TrimSpace(string(line)) == ("origin/" + branch) {
			return true
		}
	}

	return false
}

func execCmd(input string) {
	inputs := strings.Split(input, " ")
	name := inputs[0]
	args := inputs[1:]

	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
