package toyserver

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	tw "toy-web"
	tc "toy-web/internal/toycontext"
	"toy-web/internal/toyrouter/factory"
	_ "toy-web/internal/toyrouter/v1"
	_ "toy-web/internal/toyrouter/v2"
)

type ToyServer struct {
	Name   string
	mid    []tw.Middleware
	Router tw.IRouter
}

func New(name string) (*ToyServer, error) {
	v := loadRouter()
	log.Printf("current router version is: %s", v)

	router, err := factory.New(v)
	if err != nil {
		return nil, err
	}

	return &ToyServer{
		Name:   name,
		Router: router,
	}, nil
}

func (ts *ToyServer) Start(addr string) error {
	return http.ListenAndServe(addr, ts)
}

func (ts *ToyServer) Use(middleware tw.Middleware) {
	ts.mid = append(ts.mid, middleware)
}

func (ts *ToyServer) Map(pattern, method string, action tw.Action) error {
	return ts.Router.Map(pattern, method, action)
}

func (ts *ToyServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := tc.New(w, req)
	handler, ok := ts.Router.Match(req.URL.Path, req.Method, ctx)
	if !ok {
		if err := ctx.NotFound(fmt.Sprintf("route handler was not registed: %s", req.URL.Path)); err != nil {
			panic(err)
		}
		return
	}

	handler = ts.bindMiddleware(handler)
	handler(ctx)
}

func (ts *ToyServer) bindMiddleware(handler tw.Action) tw.Action {
	if length := len(ts.mid); length != 0 {
		for i := length - 1; i >= 0; i-- {
			m := ts.mid[i]
			handler = m(handler)
		}
	}
	return handler
}

func loadRouter() string {
	v := "v1"
	args := os.Args
	for _, val := range args {
		if strings.HasPrefix(val, "-router=") {
			v = strings.TrimPrefix(val, "-router=")
			break
		}
	}
	return v
}
