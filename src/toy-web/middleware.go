package toy_web

type Middleware func(next Action) Action
