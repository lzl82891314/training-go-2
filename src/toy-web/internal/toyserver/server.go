package toyserver

import (
	"fmt"
	"net/http"
	tw "toy-web"
)

type ToyServer struct {
	Name       string
	Middleware tw.HandlerFunc
	tw.Router
}

func (ts *ToyServer) Start(addr string) error {
	return http.ListenAndServe(addr, ts)
}

func (ts *ToyServer) Map(pattern, method string, handlerFunc tw.HandlerFunc) error {
	return ts.Router.Map(pattern, method, handlerFunc)
}

func (ts *ToyServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := tw.New(w, req)
	ts.Middleware(ctx)
	handler, ok := ts.Router.Match(req.URL.Path, req.Method)
	if !ok {
		if err := ctx.NotFound(fmt.Sprintf("route handler was not registed: %s", req.URL.Path)); err != nil {
			panic(err)
		}
		return
	}
	handler(ctx)
}
