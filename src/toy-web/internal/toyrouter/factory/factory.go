package factory

import (
	"fmt"
	"sync"
	tw "toy-web"
)

var (
	routers = make(map[string]tw.Router, 3)
	mutex   = sync.RWMutex{}
)

func Register(name string, router tw.Router) {
	mutex.Lock()
	defer mutex.Unlock()
	routers[name] = router
}

func New(name string) (tw.Router, error) {
	router, ok := routers[name]
	if !ok {
		return nil, fmt.Errorf("router [%s] not exists", name)
	}
	return router, nil
}
