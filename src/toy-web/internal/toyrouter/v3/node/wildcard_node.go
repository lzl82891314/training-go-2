package node

import (
	"strings"
	tw "toy-web"
)

type WildcardNode struct {
	segment  string
	children []INode
	handlers map[string]tw.Action
	value    int
}

func newWildcardNode() INode {
	return &WildcardNode{
		segment:  WildcardSymbol,
		children: nil, // 通配符之后不需要子路径，因此不需要初始化
		handlers: make(map[string]tw.Action, 2),
		value:    Wildcard,
	}
}

func isWildcardNode(segment string) bool {
	return segment == WildcardSymbol
}

func (n *WildcardNode) GetSegment() string {
	return n.segment
}

func (n *WildcardNode) GetValue() int {
	return n.value
}

func (n *WildcardNode) GetChildren() []INode {
	return n.children
}

func (n *WildcardNode) SetChild(child INode) {
	for _, v := range n.children {
		if v.GetSegment() == child.GetSegment() {
			return
		}
	}
	n.children = append(n.children, child)
}

func (n *WildcardNode) GetAction(method string) (tw.Action, bool) {
	action, ok := n.handlers[strings.ToUpper(method)]
	return action, ok
}

func (n *WildcardNode) SetAction(method string, action tw.Action) {
	n.handlers[strings.ToUpper(method)] = action
}

func (n *WildcardNode) MatchSegment(segment string) bool {
	return isWildcardNode(segment)
}

func (n *WildcardNode) Match(segment string, ctx tw.IContext) bool {
	// 通配符完全匹配
	return true
}
