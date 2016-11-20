package handlers

import (
	"sync"

	"github.com/roboll/watcher/watch"
)

var (
	lock     = sync.RWMutex{}
	handlers = map[string]watch.Handler{}
)

func Register(name string, handler watch.Handler) {
	lock.Lock()
	defer lock.Unlock()

	handlers[name] = handler
}

func Get(name string) (handler watch.Handler, ok bool) {
	lock.RLock()
	defer lock.RUnlock()

	handler, ok = handlers[name]
	return
}
