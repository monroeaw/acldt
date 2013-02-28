package main

import (
	"github.com/howeyc/fsnotify"
	"log"
	"os"
	"path/filepath"
)

var cmdSemaphoreciWatch = &Command{
	Usage: "semaphoreci:watch [<dir>]",
	Short: "watch for Semaphore CI build for a git directory",
	Long: `
Run Git rebase on a branch and then run Git merge with no fast forward
(git merge --no-ff).

As an example, assuming current branch is master, running this command
rebases a list of topic branches on top of master and then merge them
into master with no fast forward.

  $ acldt git:rmerge topic1 topic2 ...
`,
}

func init() {
	cmdSemaphoreciWatch.Run = runSemaphoreciWatch
}

func getWatchDirs(paths []string) []string {
	dirsSet := make(map[string]struct{})
	for _, path := range paths {
		filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				dirsSet[path] = struct{}{}
			}

			return nil
		})
	}

	dirs := []string{}
	for dir := range dirsSet {
		dirs = append(dirs, dir)
	}

	return dirs
}

func runSemaphoreciWatch(cmd *Command, args []string) {
	if len(args) == 0 {
		cmd.printUsage()
		return
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				log.Println("event:", ev)
			case err := <-watcher.Error:
				log.Println("Erroror:", err)
			}
		}
	}()

	for _, dir := range getWatchDirs(args) {
		log.Println(dir)
		err = watcher.Watch(dir)
		if err != nil {
			log.Fatal(err)
		}
	}

	defer watcher.Close()

	select {}
}
