package builder

import (
	"os"
	"strings"
	"sync"
	tw "toy-web"
	"toy-web/internal/toyserver"
)

var (
	routerMap = make(map[string]tw.Router, 3)
	mutex     = sync.RWMutex{}
	mws       = make([]tw.Middleware, 0, 5)
)

func Register(name string, router tw.Router) {
	mutex.Lock()
	defer mutex.Unlock()
	routerMap[name] = router
}

func Use(middleware tw.Middleware) {
	mws = append(mws, middleware)
}

func Build(name string) (tw.Server, error) {
	v := "v1"
	args := os.Args
	for _, val := range args {
		if strings.HasPrefix(val, "-router=") {
			v = strings.TrimPrefix(val, "-router=")
			break
		}
	}

	var root = func(ctx *tw.Context) {}
	if length := len(mws); length != 0 {
		for i := length - 1; i >= 0; i-- {
			m := mws[i]
			root = m(root)
		}
	}

	server := &toyserver.ToyServer{
		Name:       name,
		Router:     routerMap[v],
		Middleware: root,
	}
	return server, nil
}
