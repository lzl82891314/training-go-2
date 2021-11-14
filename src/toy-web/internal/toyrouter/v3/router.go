package v3

import (
	"errors"
	"fmt"
	"strings"
	tw "toy-web"
	"toy-web/internal/toyrouter/factory"
	"toy-web/internal/toyrouter/v3/node"
)

// 版本二基本可以满足日常的需求了
// 但是其实还可以进一步优化，因为路由的匹配优先级就其实可以抽象出来
// 比如优先级最高的是静态路由、向上是正则表达式、通配符等等
// 因此可以看到v2的版本是没有任何扩展性的，只要后续有新的匹配规则要加入 比如正则匹配 都需要修改原始代码
// 因此v3就是对v2进行进一步的抽象，增加代码扩展性

var (
	PatternInvalidErr = errors.New("route pattern is invalid")
	MethodInvalidErr  = errors.New("route method is invalid")
)

func init() {
	root, _ := node.NewNode(node.Root, "")
	factory.Register("v3", &ToyRouter{
		node:     root,
		priority: &node.ByValue{},
	})
}

type ToyRouter struct {
	node     node.INode
	priority node.INodePriority
}

var _ tw.IRouter = &ToyRouter{}

func (t *ToyRouter) Map(pattern, method string, action tw.Action) error {
	if err := validateRoute(pattern, method); err != nil {
		return err
	}

	segments := splitPattern(pattern)
	return t.doMap(0, t.node, segments, method, action)
}

func validateRoute(pattern, method string) error {
	if pattern == "" {
		return PatternInvalidErr
	}
	pos := strings.Index(pattern, node.WildcardSymbol)
	if pos > 0 {
		// 通配符必须是最后一个
		if pos != len(pattern)-1 {
			return PatternInvalidErr
		}
		if pattern[pos-1] != '/' {
			return PatternInvalidErr
		}
	}
	if method == "" {
		return MethodInvalidErr
	}

	m := strings.ToUpper(method)
	if m != "GET" && m != "POST" && m != "PUT" && m != "DELETE" {
		return MethodInvalidErr
	}
	return nil
}

func splitPattern(pattern string) []string {
	do := strings.Trim(pattern, node.RootSymbol)
	split := strings.Split(do, node.RootSymbol)
	ans := make([]string, 1, len(split)+1)
	ans[0] = node.RootSymbol
	for _, s := range split {
		if s == "" {
			continue
		}
		ans = append(ans, s)
	}
	return ans
}

func (t *ToyRouter) doMap(i int, n node.INode, segments []string, method string, action tw.Action) error {
	if n == nil {
		return errors.New("root route didn't register")
	}
	if i == len(segments) {
		n.SetAction(method, action)
		return nil
	}

	s := segments[i]
	if n.Match(s, nil) {
		if i == len(segments)-1 {
			n.SetAction(method, action)
			return nil
		}
		get, ok := t.priority.Get(n.GetChildren())
		if ok && get.Match(segments[i+1], nil) && get.GetChildren() != nil {
			return t.doMap(i+1, get, segments, method, action)
		}
		return t.doMap(i+1, n, segments, method, action)
	}

	if n.GetChildren() == nil {
		return fmt.Errorf("route node [%s] can not register child node", n.GetSegment())
	}
	ch, err := node.NewNodeBySegment(s)
	if err != nil {
		return err
	}
	n.SetChild(ch)
	n = ch
	return t.doMap(i+1, n, segments, method, action)
}

func (t *ToyRouter) Match(path, method string, ctx tw.IContext) (tw.Action, bool) {
	if err := validateRoute(path, method); err != nil {
		return nil, false
	}

	segments := splitPattern(path)
	n, ok := t.doMatch(0, t.node, segments, ctx)
	if !ok {
		return nil, false
	}
	return n.GetAction(method)
}

func (t *ToyRouter) doMatch(i int, n node.INode, segments []string, ctx tw.IContext) (node.INode, bool) {
	if n == nil {
		return nil, false
	}
	s := segments[i]
	if ok := n.Match(s, ctx); !ok {
		return nil, false
	}
	if i == len(segments)-1 {
		return n, true
	}

	candidates := n.GetChildren()
	if candidates == nil {
		return t.doMatch(i+1, n, segments, ctx)
	}
	for len(candidates) > 0 {
		ch, ok := t.priority.Get(candidates)
		if !ok {
			return nil, false
		}
		ans, ok := t.doMatch(i+1, ch, segments, ctx)
		if ok {
			return ans, ok
		}
		t.priority.RemoveMost(&candidates)
	}
	return nil, false
}
