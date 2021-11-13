package node

import (
	"strings"
	tw "toy-web"
)

type StaticNode struct {
	segment  string
	children []INode
	handlers map[string]tw.Action
	value    int
}

func newStaticNode(segment string) INode {
	return &StaticNode{
		segment:  segment,
		children: make([]INode, 0, 3),
		handlers: make(map[string]tw.Action, 2),
		value:    Static,
	}
}

func isStaticNode(segment string) bool {
	segment = strings.Trim(segment, " ")
	if segment == "" {
		return false
	}
	return segment != RootSymbol && segment != WildcardSymbol && strings.HasPrefix(segment, ParamSymbol)
}

func (n *StaticNode) GetSegment() string {
	return n.segment
}

func (n *StaticNode) GetValue() int {
	return n.value
}

func (n *StaticNode) GetChildren() []INode {
	return n.children
}

func (n *StaticNode) SetChild(child INode) {
	for _, v := range n.children {
		if v.GetSegment() == child.GetSegment() {
			return
		}
	}
	n.children = append(n.children, child)
}

func (n *StaticNode) GetAction(method string) tw.Action {
	return n.handlers[method]
}

func (n *StaticNode) SetAction(method string, action tw.Action) {
	n.handlers[method] = action
}

func (n *StaticNode) MatchSegment(segment string) bool {
	return isStaticNode(segment)
}

func (n *StaticNode) Match(segment string, ctx tw.IContext) bool {
	return n.segment == segment && segment != WildcardSymbol
}
