package main

import (
	"github.com/howeyc/fsnotify"
	"log"
	"os"
	"path/filepath"
)

var cmdSemaphoreciWatch = &Command{
	Usage: "semaphoreci:watch [<dir>]",
	Short: "watch for Semaphore CI builds for git repositories",
	Long: `
Watch for Semaphore CI builds for a list of git repositories. For example, 

  $ acldt semaphoreci:watch dir1, dir2 ...
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
