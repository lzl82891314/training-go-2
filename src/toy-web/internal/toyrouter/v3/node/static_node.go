package node

import (
	"log"
	"strings"
	tw "toy-web"
)

func init() {
	err := Register(Static, newStaticNode, isStaticNode)
	if err != nil {
		log.Fatalln(err)
	}
}

type staticNode struct {
	segment  string
	children []INode
	handlers map[string]tw.Action
	value    int
}

func newStaticNode(segment string) INode {
	return &staticNode{
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
	return segment != RootSymbol && segment != WildcardSymbol && !strings.HasPrefix(segment, ParamSymbol)
}

func (n *staticNode) GetSegment() string {
	return n.segment
}

func (n *staticNode) GetValue() int {
	return n.value
}

func (n *staticNode) GetChildren() []INode {
	return n.children
}

func (n *staticNode) SetChild(child INode) {
	for _, v := range n.children {
		if v.GetSegment() == child.GetSegment() {
			return
		}
	}
	n.children = append(n.children, child)
}

func (n *staticNode) GetAction(method string) (tw.Action, bool) {
	action, ok := n.handlers[strings.ToUpper(method)]
	return action, ok
}

func (n *staticNode) SetAction(method string, action tw.Action) {
	n.handlers[strings.ToUpper(method)] = action
}

func (n *staticNode) MatchSegment(segment string) bool {
	return isStaticNode(segment)
}

func (n *staticNode) Match(segment string, ctx tw.IContext) bool {
	return n.segment == segment && segment != WildcardSymbol
}
