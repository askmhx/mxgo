package mxgo

import (
	"github.com/howeyc/fsnotify"
)



func WatchWorker(filePaths ...string){
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		mxLog.Error(err)
	}

	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				mxLog.Info("event:", ev)
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
