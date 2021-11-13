package toy_web

type Server interface {
	Start(addr string) error
	Router
}

type ServerBuilder interface {
	UseMiddleware(middleware Middleware) ServerBuilder
	UseRoute(pattern, method string, action Action) ServerBuilder
	Build(name string) (Server, error)
}
