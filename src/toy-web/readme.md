# 玩具web框架

由于 http包下的server能用的功能只有最小集，为了更好更方便的使用，可以手写一个小的web框架用于练手。

因此所有的功能都是基于 net.http下的基础包完成

## web框架的基础功能

1. 请求上下文 HttpContext
2. 标准的输出 Json Response
3. 请求中间件 Middlewares
4. 路由注册 Router Tree
5. 优雅退出 Graceful Shutdown

## 内部组件

### Serverable

Server服务主体，应该要满足如下的基本使用方式：

``` go
// 创建server
s := NewServer()

// 注册路由配置
s.Route("pattern", Handler)
s.Route("pattern", Handler)
s.Route("pattern", Handler)
s.Route("pattern", Handler)
s.Route("pattern", Handler)

// 启动服务
s.Start("address")

// 退出服务
s.Shutdown()
```

因此 ToyServer应该是一个满足上述几个方法的接口。大致如下：

``` go
type Serverable interface {
    Start(add string) error
    Route(pattern string, Handler) error
    Shutdown() error
}
```

此外，还有一个包方法: `NewServer() Serverable`

### Routeable

上述可以看到，路由是隔离于Server主服务的一个从属服务，Server和Route在服务主体上应该是依赖关系。

并且Route逻辑也应该可以有多个实现策略，比如最普通的基于Map强硬匹配，基于前缀树的最长前缀匹配等。

因此，抽象出一个Routeable接口用于承载路由的逻辑：

``` go
type Routeable interface {
	Route(pattern string, Handler) error
}
```

### HttpHandler

可以看到，有了Server服务和Route之后，其内都需要依赖一个Handler来处理真正的请求。因此可以对Handler进行抽象。

``` go
type Handler interface {
    HandleHttp(ctx *HttpContext)
}
```

### HttpContext

最后，HttpContext其实就是对net.http包中请求和响应的封装。

``` go
type HttpContext struct {
    Req *http.Request
    RspWriter http.ResponseWriter
}
```

至此，一个比较简单的web框架的结构就明确了。