package v1

import (
	"fmt"
	"strings"
	tw "toy-web"
	"toy-web/internal/toyrouter/factory"
)

// 第一种路由实现为强硬匹配，即：
// 只能路由类似于Route("home/index", handler)这类的实现
// 不支持通配符，不支持静态资源，不支持正则匹配

func init() {
	factory.Register("v1", &ToyRouter{
		node: make(map[string]tw.Action, 5),
	})
}

var _ tw.IRouter = &ToyRouter{}

type ToyRouter struct {
	node map[string]tw.Action
}

func (m *ToyRouter) Map(pattern, method string, handleFunc tw.Action) error {
	pattern = strings.Trim(pattern, "/")
	key := generateKey(pattern, method)
	_, ok := m.node[key]
	if ok {
		return fmt.Errorf("duplicated route: %s", key)
	}
	m.node[key] = handleFunc
	return nil
}

func (m *ToyRouter) Match(path, method string, ctx tw.IContext) (tw.Action, bool) {
	purePath := strings.Trim(path, "/")
	key := generateKey(purePath, method)
	load, ok := m.node[key]
	if !ok {
		return nil, ok
	}
	return load, true
}

func generateKey(pattern, method string) string {
	upper := strings.ToUpper(method)
	return fmt.Sprintf("%s$%s", upper, pattern)
}
