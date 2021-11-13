package toy_web

type Server interface {
	Use(middleware Middleware)
	Router
	Start(addr string) error
}
