package v2

import tw "toy-web"

type TreeNode struct {
	segment  string
	children []*TreeNode
	handlers map[string]tw.Action
}

func NewNode(segment string) *TreeNode {
	return &TreeNode{
		segment:  segment,
		children: make([]*TreeNode, 0, 3),
		handlers: make(map[string]tw.Action, 2),
	}
}
