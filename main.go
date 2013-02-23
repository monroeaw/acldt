package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Command struct {
	// args does not include the command name
	Run  func(cmd *Command, args []string)
	Flag flag.FlagSet

	Usage string // first word is the command name
	Short string // `acldt help` output
	Long  string // `acldt help <cmd>` output
}

func (c *Command) printUsage() {
	if c.Runnable() {
		fmt.Printf("Usage: acldt %s\n\n", c.Usage)
	}
	fmt.Println(strings.TrimSpace(c.Long))
}

func (c *Command) Name() string {
	name := c.Usage
	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}
	return name
}

func (c *Command) Runnable() bool {
	return c.Run != nil
}

func (c *Command) ShowUsage() bool {
	return c.Short != ""
}

var commands = []*Command{
	cmdGitRmerge,
	cmdGitDbranch,
	cmdVersion,
	cmdHelp,
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		usage()
	}

	for _, cmd := range commands {
		if cmd.Name() == args[0] && cmd.Run != nil {
			cmd.Flag.Usage = usage
			cmd.Flag.Parse(args[1:])
			cmd.Run(cmd, cmd.Flag.Args())
			return
		}
	}

	fmt.Fprintf(os.Stderr, "Unknown command: %s\n", args[0])
	usage()
}
