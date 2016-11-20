package main

import (
	"flag"
	"log"
	"strings"

	"github.com/roboll/watcher/handlers"
	"github.com/roboll/watcher/watch"
)

var (
	targets     = flag.String("targets", "", "comma-separated files and/or dirs to watch")
	handlerType = flag.String("handler", "", "handler type")
)

func main() {
	flag.Parse()

	if *targets == "" {
		log.Fatal("targets is required")
	}

	handler, ok := handlers.Get(*handlerType)
	if !ok {
		log.Fatalf("invalid handler: %s", *handlerType)
	}

	watch.Watch(handler, strings.Split(*targets, ",")...)
}
