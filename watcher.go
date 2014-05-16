package mxgo

import (
	"log"
	"github.com/howeyc/fsnotify"
)


func Watch(callBack interface {},filePath ...string){
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
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Watch(filePath)
	if err != nil {
		log.Fatal(err)
	}
}
