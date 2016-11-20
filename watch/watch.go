package watch

import (
	"log"

	fsnotify "gopkg.in/fsnotify.v1"
)

type Handler func(string) error

func Watch(handler Handler, targets ...string) error {
	watch, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watch.Close()

	for _, target := range targets {
		log.Printf("adding target %s", target)
		if err := watch.Add(target); err != nil {
			return err
		}
	}

	for {
		select {
		case err := <-watch.Errors:
			log.Printf("watcher error: %s", err)
		case evt := <-watch.Events:
			log.Printf("watcher event: %s", evt)
			if evt.Op&fsnotify.Create == fsnotify.Create ||
				evt.Op&fsnotify.Write == fsnotify.Write ||
				evt.Op&fsnotify.Remove == fsnotify.Remove ||
				evt.Op&fsnotify.Rename == fsnotify.Rename {

				if err := handler(evt.Name); err != nil {
					log.Printf("handler error: %s", err)
				}
			}
		}
	}
}
