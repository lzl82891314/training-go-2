package factory

import (
	"fmt"
	"sync"
	tw "toy-web"
)

var (
	routers = make(map[string]tw.IRouter, 3)
	mutex   = sync.RWMutex{}
)

func Register(name string, router tw.IRouter) {
	mutex.Lock()
	defer mutex.Unlock()
	routers[name] = router
}

func New(name string) (tw.IRouter, error) {
	router, ok := routers[name]
	if !ok {
		return nil, fmt.Errorf("router [%s] not exists", name)
	}
	return router, nil
}
