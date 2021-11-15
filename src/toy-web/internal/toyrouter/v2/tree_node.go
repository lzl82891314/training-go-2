package v2

import tw "toy-web"

type treeNode struct {
	segment  string
	children []*treeNode
	handlers map[string]tw.Action
}

func NewNode(segment string) *treeNode {
	return &treeNode{
		segment:  segment,
		children: make([]*treeNode, 0, 3),
		handlers: make(map[string]tw.Action, 2),
	}
}
