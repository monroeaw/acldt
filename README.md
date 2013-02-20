##acldt

ACL Development Tools.

### Overview

A set of development tools for ACL developers.

### Usage

```plain
$ acldt help
Usage: acldt [command] [options] [arguments]

Supported commands are:

  git:rmerge run Git rebase and Git merge with --no-ff
  version  show acldt version
  help     show help

See 'acldt help [command]' for more information about a command.
```

```plain
$ acldt help git:rmerge
Usage: acldt git:rmerge [<branch>]

Run Git rebase on a branch and then run Git merge with no fast forward
(git merge --no-ff).

As an example, assuming current branch is master, running this command
rebases a list of topic branches on top of master and then merge them
into master with no fast forward.

  $ acldt git:rmerge topic1 topic2
```

### Installation

acldt is in beta but you are welcome to try it out. You'll need to
build it with [Go](http://code.google.com/p/go/) yourself for now.

```plain
$ go build
$ cp acldt /usr/local/bin/
$ chmod +x /usr/local/bin/acldt
```
