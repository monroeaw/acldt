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

### Installation

acldt is in beta but you are welcome to try it out. You'll need to
build it with [Go](http://code.google.com/p/go/) yourself for now.

```plain
$ go build
$ cp acldt /usr/local/bin
$ chmod +x /usr/local/bin/acldt
```
