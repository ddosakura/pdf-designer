package watcher

import (
	"io/ioutil"
	"path/filepath"

	"github.com/ddosakura/gklang"
	"github.com/fsnotify/fsnotify"
)

const (
	delay = 10
)

var (
	watcher *fsnotify.Watcher
)

// Init the watcher
func Init(src string) {
	var err error
	watcher, err = fsnotify.NewWatcher()
	if err != nil {
		gklang.Er(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				eventDispatcher(event)
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				gklang.Log(gklang.LWarn, err)
			}
		}
	}()
	addWatcher(src)
	<-done
}

func addWatcher(src string) {
	dirs := make([]string, 0)
	dirs = append(dirs, src)
	dirs = appendDirWatcher(dirs, src)

	for _, dir := range dirs {
		gklang.Log(gklang.LInfo, "watcher add -> ", dir)
		err := watcher.Add(dir)
		if err != nil {
			gklang.Er(err)
		}
	}
	gklang.Log(gklang.LInfo, "ready!")
}

func appendDirWatcher(dirs []string, path string) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		gklang.Er(err)
	}
	for _, v := range files {
		if v.IsDir() {
			p := filepath.Join(path, v.Name())
			dirs = append(dirs, p)
			dirs = appendDirWatcher(dirs, p)
		}
	}
	return dirs
}

func eventDispatcher(event fsnotify.Event) {
	switch event.Op {
	case
		fsnotify.Write,
		fsnotify.Rename:
		gklang.Log(gklang.LInfo, "EVENT", event.Op.String(), event.Name)
		go callFreshWebPage(event)
	case fsnotify.Remove:
	case fsnotify.Create:
	}
}
