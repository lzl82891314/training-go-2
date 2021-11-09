package toy_web

type Handler interface {
	Handle(ctx *Context)
}

type HandlerFunc func(ctx *Context)
