package toy_web

type Middleware func(next HandlerFunc) HandlerFunc
