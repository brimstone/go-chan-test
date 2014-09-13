package dir

import (
	"github.com/brimstone/go-chan-test/channel"
	"gopkg.in/fsnotify.v1"
	"log"
)

type Dir struct {
	directory string
	watcher   *fsnotify.Watcher
}

func (dir *Dir) Init(directory string) error {
	dir.directory = directory
	var err error
	dir.watcher, err = fsnotify.NewWatcher()
	return err
}

func (dir *Dir) Sync(cf chan channel.File) {
	defer dir.watcher.Close()

	err := dir.watcher.Add(dir.directory)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case event := <-dir.watcher.Events:
			if event.Op&fsnotify.Write == fsnotify.Write {
				cf <- channel.File{Filename: event.Name}
			} else if event.Op&fsnotify.Remove == fsnotify.Remove {
				cf <- channel.File{Filename: event.Name}
			}
		case err := <-dir.watcher.Errors:
			log.Println("error:", err)
		case file := <-cf:
			log.Println("Got notification about:", file.Filename)
		}
	}
}

func New(directory string) (*Dir, error) {
	dir := new(Dir)
	err := dir.Init(directory)
	if err != nil {
		return nil, err
	}
	return dir, nil
}
