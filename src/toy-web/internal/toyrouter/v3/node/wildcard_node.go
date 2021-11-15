package node

import (
	"log"
	"strings"
	tw "toy-web"
)

func init() {
	err := Register(Wildcard, newWildcardNode, isWildcardNode)
	if err != nil {
		log.Fatalln(err)
	}
}

type wildcardNode struct {
	segment  string
	children []INode
	handlers map[string]tw.Action
	value    int
}

func newWildcardNode(segment string) INode {
	return &wildcardNode{
		segment:  WildcardSymbol,
		children: nil, // 通配符之后不需要子路径，因此不需要初始化
		handlers: make(map[string]tw.Action, 2),
		value:    Wildcard,
	}
}

func isWildcardNode(segment string) bool {
	return segment == WildcardSymbol
}

func (n *wildcardNode) GetSegment() string {
	return n.segment
}

func (n *wildcardNode) GetValue() int {
	return n.value
}

func (n *wildcardNode) GetChildren() []INode {
	return n.children
}

func (n *wildcardNode) SetChild(child INode) {
	for _, v := range n.children {
		if v.GetSegment() == child.GetSegment() {
			return
		}
	}
	n.children = append(n.children, child)
}

func (n *wildcardNode) GetAction(method string) (tw.Action, bool) {
	action, ok := n.handlers[strings.ToUpper(method)]
	return action, ok
}

func (n *wildcardNode) SetAction(method string, action tw.Action) {
	n.handlers[strings.ToUpper(method)] = action
}

func (n *wildcardNode) MatchSegment(segment string) bool {
	return isWildcardNode(segment)
}

func (n *wildcardNode) Match(segment string, ctx tw.IContext) bool {
	// 通配符完全匹配
	return true
}
