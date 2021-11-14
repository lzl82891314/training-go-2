package node

import (
	"log"
	"strings"
	tw "toy-web"
)

func init() {
	err := Register(Root, newRootNode, isRootNode)
	if err != nil {
		log.Fatalln(err)
	}
}

type RootNode struct {
	segment  string
	children []INode
	handlers map[string]tw.Action
	value    int
}

func newRootNode(segment string) INode {
	return &RootNode{
		segment:  RootSymbol,
		children: make([]INode, 0, 3),
		handlers: make(map[string]tw.Action, 2),
		value:    Root,
	}
}

func isRootNode(segment string) bool {
	segment = strings.Trim(segment, " ")
	if segment == "" {
		return false
	}
	return segment == RootSymbol
}

func (n *RootNode) GetSegment() string {
	return n.segment
}

func (n *RootNode) GetValue() int {
	return n.value
}

func (n *RootNode) GetChildren() []INode {
	return n.children
}

func (n *RootNode) SetChild(child INode) {
	for _, v := range n.children {
		if v.GetSegment() == child.GetSegment() {
			return
		}
	}
	n.children = append(n.children, child)
}

func (n *RootNode) GetAction(method string) (tw.Action, bool) {
	action, ok := n.handlers[strings.ToUpper(method)]
	return action, ok
}

func (n *RootNode) SetAction(method string, action tw.Action) {
	n.handlers[strings.ToUpper(method)] = action
}

func (n *RootNode) MatchSegment(segment string) bool {
	return isRootNode(segment)
}

func (n *RootNode) Match(segment string, ctx tw.IContext) bool {
	return n.segment == segment
}
