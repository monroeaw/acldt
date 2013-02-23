package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"
)

var cmdGitRmerge = &Command{
	Usage: "git:rmerge [<branch>]",
	Short: "run Git rebase and Git merge with --no-ff against current branch",
	Long: `
Run Git rebase on a branch and then run Git merge with no fast forward
(git merge --no-ff).

As an example, assuming current branch is master, running this command
rebases a list of topic branches on top of master and then merge them
into master with no fast forward.

  $ acldt git:rmerge topic1 topic2 ...
`,
}

var cmdGitDbranch = &Command{
	Usage: "git:dbranch [<branch>]",
	Short: "delete local and remote branches",
	Long: `
Delete local and remote branches. For example,

  $ acldt git:dbranch branch1 branch2 ...
`,
}

func init() {
	cmdGitRmerge.Run = runGitRmerge
	cmdGitDbranch.Run = runGitDbranch
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
		if hasRemoteBranch(baseBranch) {
			execCmd("git pull origin " + baseBranch)
		}
		execCmd("git merge " + topicBranch + " --no-ff")
		execCmd("git push origin HEAD")

		deleteBranch(topicBranch)
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

func deleteBranch(branch string) {
	execCmd("git branch -D " + branch)
	execCmd("git push origin :" + branch)
}

func runGitDbranch(cmd *Command, args []string) {
	if len(args) == 0 {
		cmd.printUsage()
		return
	}

	for _, arg := range args {
		branch := strings.TrimSpace(arg)
		deleteBranch(branch)
	}
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
