package toy_web

type Server interface {
	Start(addr string) error
	Shutdown() error
}

type ServerBuilder interface {
	UseMiddleware(middleware Middleware) ServerBuilder
	UseRoute(pattern, method string, handlerFunc HandlerFunc) ServerBuilder
	Build(name string) (Server, error)
}
