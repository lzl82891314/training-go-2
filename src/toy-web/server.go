package toy_web

type IServer interface {
	Use(middleware Middleware)
	IRouter
	Start(addr string) error
}
