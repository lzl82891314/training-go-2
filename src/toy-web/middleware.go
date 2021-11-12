package toy_web

type Middleware func(next HandlerFunc) HandlerFunc

type HandlerFunc func(ctx *Context)
