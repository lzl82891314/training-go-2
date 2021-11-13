package builder

import (
	"log"
	"os"
	"strings"
	tw "toy-web"
	"toy-web/internal/toyrouter/factory"
	_ "toy-web/internal/toyrouter/v1"
	_ "toy-web/internal/toyrouter/v2"
	"toy-web/internal/toyserver"
)

var (
	mws = make([]tw.Middleware, 0, 5)
)

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

	log.Printf("current router version is: %s", v)
	router, err := factory.New(v)
	if err != nil {
		return nil, err
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
		Router:     router,
		Middleware: root,
	}
	return server, nil
}
