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

type rootNode struct {
	segment  string
	children []INode
	handlers map[string]tw.Action
	value    int
}

func newRootNode(segment string) INode {
	return &rootNode{
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

func (n *rootNode) GetSegment() string {
	return n.segment
}

func (n *rootNode) GetValue() int {
	return n.value
}

func (n *rootNode) GetChildren() []INode {
	return n.children
}

func (n *rootNode) SetChild(child INode) {
	for _, v := range n.children {
		if v.GetSegment() == child.GetSegment() {
			return
		}
	}
	n.children = append(n.children, child)
}

func (n *rootNode) GetAction(method string) (tw.Action, bool) {
	action, ok := n.handlers[strings.ToUpper(method)]
	return action, ok
}

func (n *rootNode) SetAction(method string, action tw.Action) {
	n.handlers[strings.ToUpper(method)] = action
}

func (n *rootNode) MatchSegment(segment string) bool {
	return isRootNode(segment)
}

func (n *rootNode) Match(segment string, ctx tw.IContext) bool {
	return n.segment == segment
}
