package toy_web

type IRouter interface {
	Map(pattern, method string, action Action) error
	Match(path, method string, ctx IContext) (Action, bool)
}
