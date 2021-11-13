package v2

import (
	"errors"
	"strings"
	tw "toy-web"
	"toy-web/internal/toyrouter/factory"
)

// 第一种路由的问题是映射过于强硬，只能强路由映射每一个地址
// 但是正常的使用情况可能是可以映射比如 Home/*这种类型的留有规则以满足：
// 如果有存在 Home/Index 则返回Index的handler，如果没有则返回Home/*的handler
// 上述这种需求使用Map就很难实现了，因此进一步优化路由的数据结构，由单一Map改为路由树
// 每段路径的segment就是树上的各个节点

func init() {
	factory.Register("v2", &ToyRouter{
		router: CreateTreeNode("/"),
	})
}

type ToyRouter struct {
	router *TreeNode
}

func (t *ToyRouter) Map(pattern, method string, handleFunc tw.HandlerFunc) error {
	if pattern == "/" {
		t.router.handlers[strings.ToUpper(method)] = handleFunc
		return nil
	}
	if err := pathValidator(pattern); err != nil {
		return err
	}
	pattern = strings.Trim(pattern, "/")
	segments := strings.Split(pattern, "/")
	if len(segments) == 0 {
		return errors.New("router path can not be null or empty")
	}

	p, cur := 0, t.router
	for len(cur.children) > 0 {
		if p == len(segments) {
			// 说明已经找到，直接返回
			break
		}
		found, segment := false, segments[p]
		for _, ch := range cur.children {
			if ch.segment == segment {
				cur, p, found = ch, p+1, true
				break
			}
		}
		if !found {
			// 没有匹配到就直接返回生成子树
			break
		}
	}
	// 生成子树，如果完美匹配则在函数内直接退出
	cur = treeGenerator(cur, p, segments)
	cur.handlers[strings.ToUpper(method)] = handleFunc
	return nil
}

func pathValidator(pattern string) error {
	pos := strings.Index(pattern, "*")
	if pos > 0 {
		// 通配符必须是最后一个
		if pos != len(pattern)-1 {
			return errors.New("invalid route pattern")
		}
		if pattern[pos-1] != '/' {
			return errors.New("invalid route pattern")
		}
	}
	return nil
}

func treeGenerator(node *TreeNode, p int, segments []string) *TreeNode {
	if p == len(segments) {
		return node
	}
	segment := segments[p]
	cur := CreateTreeNode(segment)
	node.children = append(node.children, cur)
	return treeGenerator(cur, p+1, segments)
}

func (t *ToyRouter) Find(path, method string) (tw.HandlerFunc, bool) {
	if path == "/" {
		return t.router.handlers[method], true
	}
	path = strings.Trim(path, "/")
	segments := strings.Split(path, "/")
	return doFind(0, method, t.router, segments, nil)
}

func doFind(p int, method string, cur *TreeNode, segments []string, wildcard *TreeNode) (tw.HandlerFunc, bool) {
	if p == len(segments) {
		handlerFunc, ok := cur.handlers[method]
		return handlerFunc, ok
	}
	segment := segments[p]
	for _, node := range cur.children {
		if node.segment == segment && node.segment != "*" {
			find, ok := doFind(p+1, method, node, segments, wildcard)
			if ok {
				return find, ok
			}
		} else if node.segment == "*" {
			// 命中通配符
			// 将通配符的绑定记录，作为后续查询的备选
			wildcard = node
		}
	}
	if wildcard == nil {
		return nil, false
	}
	handlerFunc, ok := wildcard.handlers[strings.ToUpper(method)]
	return handlerFunc, ok
}
