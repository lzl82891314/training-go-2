package toy_web

type Router interface {
	Map(pattern, method string, action Action) error
	Match(path, method string) (Action, bool)
}
