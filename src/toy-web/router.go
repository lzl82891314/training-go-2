package toy_web

type Router interface {
	Map(pattern, method string, handleFunc HandlerFunc) error
	Match(path, method string) (HandlerFunc, bool)
}
