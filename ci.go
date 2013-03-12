package main

import (
	"github.com/howeyc/fsnotify"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var cmdCiWatch = &Command{
	Usage: "ci:watch [<dir>]",
	Short: "watch for CI builds",
	Long: `
Watch for CI builds for a list of Git repositories.
Currently, only Semaphore CI is supported. For example,

  $ acldt ci:watch dir1, dir2 ...
`,
}

func init() {
	cmdCiWatch.Run = runCiWatch
}

func runCiWatch(cmd *Command, args []string) {
	if len(args) == 0 {
		cmd.printUsage()
		return
	}

	verifyGitDirs(args)
	watcher := createWatcher()
	go watchForGitPush(watcher)
	addWatchers(watcher, args)
	defer watcher.Close()

	select {}
}

func verifyGitDirs(dirs []string) {
	for _, dir := range dirs {
		if b, _ := fileExists(filepath.Join(dir, ".git")); !b {
			log.Fatal(dir + " is not a git repository")
		}
	}
}

func createWatcher() *fsnotify.Watcher {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	return watcher
}

func addWatchers(watcher *fsnotify.Watcher, dirs []string) {
	for _, dir := range dirs {
		gitRemoteDir := filepath.Join(dir, ".git", "refs", "remotes", "origin")
		err := watcher.Watch(gitRemoteDir)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func watchForGitPush(watcher *fsnotify.Watcher) {
	for {
		select {
		case ev := <-watcher.Event:
			if ev.IsCreate() {
				_, file := filepath.Split(ev.Name)
				if len(filepath.Ext(file)) == 0 {
					sha, err := ioutil.ReadFile(file)
					if err != nil {
						log.Println(sha)
						log.Fatal(err)
					}
					// TODO
				}
			}
		case err := <-watcher.Error:
			log.Println("Erroror:", err)
		}
	}
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
