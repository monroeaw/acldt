##acldt

ACL Development Tools.

### Overview

A set of development tools for ACL developers.

### Usage

#### acldt help

```plain
$ acldt help
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
   ey-foreach, ef       Applies action for each EY environment
   help, h              Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --version    print the version
   --help, -h   show help
```

#### acldt git-rmerge

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

#### acldt ey-foreach

```
$ acldt help ey-foreach
NAME:
   ey-foreach - Applies action for each EY environment

USAGE:
   command ey-foreach [command options] [arguments...]

DESCRIPTION:
   Applies action for each Engineyard environment. For example,
to upload recipes for each production environment for app Projects:

  $ acldt ey-foreach -a projects -e production recipes upload

OPTIONS:
   --a ''       app name on EY, e.g., projects
   --e ''       env name on EY, e.g., production
```

### Installation

acldt is in beta but you are welcome to try it out. You'll need to
build it with [Go](http://code.google.com/p/go/) yourself for now.

```plain
$ go build
$ cp acldt /usr/local/bin/
$ chmod +x /usr/local/bin/acldt
```
