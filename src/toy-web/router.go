package toy_web

type Router interface {
	Route(pattern, method string, handleFunc HandlerFunc) error
	Find(path, method string) (HandlerFunc, bool)
}
