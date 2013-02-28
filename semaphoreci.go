package main

import (
	"github.com/howeyc/fsnotify"
	"log"
	"os"
	"path/filepath"
  "net/http"
  "io/ioutil"
)

var cmdSemaphoreciWatch = &Command{
	Usage: "semaphoreci:watch [<dir>]",
	Short: "watch for Semaphore CI builds",
	Long: `
Watch for Semaphore CI builds for a list of git repositories. For example,

  $ acldt semaphoreci:watch dir1, dir2 ...
`,
}

func init() {
	cmdSemaphoreciWatch.Run = runSemaphoreciWatch
}

func runSemaphoreciWatch(cmd *Command, args []string) {
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
          go pullBuildResult(file)
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

func pullBuildResult(branch string) {
  resp, err := http.Get("https://semaphoreapp.com/api/v1/projects?auth_token=Yds3w6o26FLfJTnVK2y9")
  if err != nil {
    // handle error
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  log.Println(body)
}
