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
)

func init() {
	if err := Register("toy", newToyServer); err != nil {
		log.Fatalln(err)
	}
}

type toyServer struct {
	mid    []tw.Middleware
	router tw.IRouter
}

var _ tw.IServer = &toyServer{}

func newToyServer() (tw.IServer, error) {
	v := loadRouter()
	log.Printf("current router version is: %s", v)

	router, err := factory.New(v)
	if err != nil {
		return nil, err
	}

	return &toyServer{
		router: router,
	}, nil
}

func (ts *toyServer) Start(addr string) error {
	return http.ListenAndServe(addr, ts)
}

func (ts *toyServer) Use(middleware tw.Middleware) {
	ts.mid = append(ts.mid, middleware)
}

func (ts *toyServer) Map(pattern, method string, action tw.Action) error {
	return ts.router.Map(pattern, method, action)
}

func (ts *toyServer) Match(path, method string, ctx tw.IContext) (tw.Action, bool) {
	return ts.router.Match(path, method, ctx)
}

func (ts *toyServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := tc.New(w, req)
	handler, ok := ts.router.Match(req.URL.Path, req.Method, ctx)
	if !ok {
		if err := ctx.NotFound(fmt.Sprintf("route handler was not registed: %s", req.URL.Path)); err != nil {
			panic(err)
		}
		return
	}

	handler = ts.bindMiddleware(handler)
	handler(ctx)
}

func (ts *toyServer) bindMiddleware(handler tw.Action) tw.Action {
	if length := len(ts.mid); length != 0 {
		for i := length - 1; i >= 0; i-- {
			m := ts.mid[i]
			handler = m(handler)
		}
	}
	return handler
}

func loadRouter() string {
	v := "v3"
	args := os.Args
	for _, val := range args {
		if strings.HasPrefix(val, "-router=") {
			v = strings.TrimPrefix(val, "-router=")
			break
		}
	}
	return v
}
