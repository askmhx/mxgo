package mxgo

import (
	"github.com/howeyc/fsnotify"
)

type WatchAction func(filePath string,event *fsnotify.FileEvent)


func WatchDoTask(callback WatchAction,filePaths ...string){
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		mxLog.Error(err)
	}

	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				callback(ev.Name,ev)
			case err := <-watcher.Error:
				mxLog.Error("error:", err)
			}
		}
	}()
    for i := range filePaths{
		err = watcher.Watch(filePaths[i])
		if err != nil {
			mxLog.Error(err)
		}
	}

}
