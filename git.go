package main

var cmdGit = &Command{
	Run:   runGit,
	Usage: "git command [arguments]",
	Short: "run a Git command",
	Long:  `Run a Git command`,
}

func runGit(cmd *Command, args []string) {
}
