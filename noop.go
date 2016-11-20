package main

import (
	"log"

	"github.com/roboll/watcher/handlers"
)

func init() {
	handlers.Register("noop", Noop)
}

func Noop(path string) error {
	log.Printf("noop: %s", path)
	return nil
}
