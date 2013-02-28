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

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func getWatchDirs(paths []string) []string {
	dirsSet := make(map[string]struct{})
	for _, path := range paths {
		log.Println(filepath.Glob(path + "/.git/refs/remotes/origin/*"))

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

	for _, path := range args {
		if b, _ := fileExists(filepath.Join(path, ".git")); !b {
			log.Fatal(path + " is not a git repository")
		}
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

	for _, path := range args {
		gitRemoteDir := filepath.Join(path, ".git", "refs", "remotes", "origin")
		log.Println(gitRemoteDir)
		err = watcher.Watch(gitRemoteDir)
		if err != nil {
			log.Fatal(err)
		}
	}

	defer watcher.Close()

	select {}
}
