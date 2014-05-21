package mxgo

import (
	"github.com/howeyc/fsnotify"
	"path"
	"strings"
)

func WatchDoTask(callback func(filePath string, event *fsnotify.FileEvent), filePaths ...string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		mxLog.Error(err)
	}

	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				if !strings.HasPrefix(path.Base(ev.Name), ".") {
					callback(ev.Name, ev)
				}
			case <-watcher.Error:
				continue
			}
		}
	}()
	for i := range filePaths {
		err = watcher.Watch(filePaths[i])
		if err != nil {
			mxLog.Error(err)
		}
	}

}
