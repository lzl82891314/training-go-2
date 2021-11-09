package toyserver

import (
	"fmt"
	"net/http"
	tw "toy-web"
)

type ToyServer struct {
	Name string
	tw.Router
	middleware tw.HandlerFunc
}

func (ts *ToyServer) Start(addr string) error {
	return http.ListenAndServe(addr, ts)
}

func (ts *ToyServer) Shutdown() error {
	panic("implement me")
}

func (ts *ToyServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := tw.CreateContext(w, req)
	handler, ok := ts.Router.Find(ctx.Req.URL.Path, ctx.Req.Method)
	if !ok {
		ctx.NotFoundResponse(fmt.Sprintf("route handler was not registed: %s", ctx.Req.URL.Path))
		return
	}
	handler(ctx)
}
