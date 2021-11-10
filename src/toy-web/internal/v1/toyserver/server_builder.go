package toyserver

import (
	"log"
	tw "toy-web"
	"toy-web/internal/v1/toyrouter"
)

type ToyBuilder struct {
	middlewares []tw.Middleware
	tw.Router
}

func CreateToyBuilder() *ToyBuilder {
	return &ToyBuilder{
		middlewares: make([]tw.Middleware, 0, 5),
		Router:      toyrouter.CreateToyRouter(),
	}
}

func (t *ToyBuilder) UseMiddleware(middleware tw.Middleware) tw.ServerBuilder {
	t.middlewares = append(t.middlewares, middleware)
	return t
}

func (t *ToyBuilder) UseRoute(pattern, method string, handlerFunc tw.HandlerFunc) tw.ServerBuilder {
	err := t.Router.Route(pattern, method, handlerFunc)
	if err != nil {
		log.Println(err)
	}
	return t
}

func (t *ToyBuilder) Build(name string) (tw.Server, error) {
	server := &ToyServer{
		Name:   name,
		Router: t.Router,
	}
	if length := len(t.middlewares); length != 0 {
		var root = func(ctx *tw.Context) {}
		for i := length - 1; i >= 0; i-- {
			middleware := t.middlewares[i]
			root = middleware(root)
		}
		server.middleware = root
	}
	return server, nil
}
