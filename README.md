##acldt

ACL Development Tools.

### Overview

A set of development tools for ACL developers.

### Usage

```plain
NAME:
   acldt - ACL Development Tools

USAGE:
   acldt [global options] command [command options] [arguments...]

VERSION:
   dev

COMMANDS:
   git-rmerge, gm       Runs Git rebase and Git merge with --no-ff
against current branch
   git-dbranch, gd      Deletes local and remote branches
   help, h              Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --version    print the version
   --help, -h   show help
```

```plain
$ acldt help git-rmerge
NAME:
   git-rmerge - Runs Git rebase and Git merge with --no-ff against
current branch

USAGE:
   command git-rmerge [command options] [arguments...]

DESCRIPTION:

Run Git rebase on a branch and then run Git merge with no fast forward
(git merge --no-ff).

As an example, assuming current branch is master, running this command
rebases a list of topic branches on top of master and then merge them
into master with no fast forward.

  $ acldt git-rmerge topic1 topic2 ...

OPTIONS:

```

### Installation

acldt is in beta but you are welcome to try it out. You'll need to
build it with [Go](http://code.google.com/p/go/) yourself for now.

```plain
$ go build
$ cp acldt /usr/local/bin/
$ chmod +x /usr/local/bin/acldt
```
