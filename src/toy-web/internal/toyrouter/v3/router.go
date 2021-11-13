package v3

import (
	"errors"
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

func init() {
	n, _ := node.NewNode(node.Root, "/")
	factory.Register("v3", &ToyRouter{
		node: n,
	})
}

type ToyRouter struct {
	node node.INode
}

var _ tw.IRouter = &ToyRouter{}

func (t *ToyRouter) Map(pattern, method string, action tw.Action) error {
	if pattern == "/" {
		t.node.SetAction(strings.ToUpper(method), action)
		return nil
	}

	segments, err := splitPattern(pattern)
	if err != nil {
		return err
	}
	n, i := findLast(1, t.node.GetChildren(), segments)
	if n == nil {
		n, i = t.node, 1
	}
	n, err = generateNode(i, n, segments)
	if err != nil {
		return err
	}
	n.SetAction(method, action)
	return nil
}

func splitPattern(pattern string) ([]string, error) {
	pos := strings.Index(pattern, node.WildcardSymbol)
	if pos > 0 {
		// 通配符必须是最后一个
		if pos != len(pattern)-1 {
			return nil, errors.New("invalid route pattern")
		}
		if pattern[pos-1] != '/' {
			return nil, errors.New("invalid route pattern")
		}
	}
	do := strings.Trim(pattern, "/")
	return strings.Split(do, "/"), nil
}

func findLast(i int, nodes []node.INode, segments []string) (node.INode, int) {
	if i >= len(segments) {
		return nil, i
	}
	s := segments[i]
	var ans node.INode = nil
	var index = i
	for _, n := range nodes {
		if n.GetSegment() == s {
			tmp, tmpI := findLast(i+1, n.GetChildren(), segments)
			if tmp != nil {
				ans, index = tmp, tmpI
			}
		}
	}
	return ans, index
}

func generateNode(i int, n node.INode, segments []string) (node.INode, error) {
	if i >= len(segments) {
		return n, nil
	}
	s := segments[i]
	ch, err := node.NewNodeBySegment(s)
	if err != nil {
		return nil, err
	}
	n.SetChild(ch)
	return generateNode(i+1, ch, segments)
}

func (t *ToyRouter) Match(path, method string) (tw.Action, bool) {
	panic("implement me")
}
