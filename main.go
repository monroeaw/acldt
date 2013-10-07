package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "acldt"
	app.Usage = "ACL Development Tools"
	app.Version = Version
	app.Commands = []cli.Command{
		gitRmergeCmd,
		gitDbranchCmd,
		gitTagbuild,
		eyForeachCmd,
	}

	app.Run(os.Args)
}
