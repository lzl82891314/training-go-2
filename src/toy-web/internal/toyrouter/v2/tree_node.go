package v2

import tw "toy-web"

type TreeNode struct {
	segment  string
	children []*TreeNode
	handlers map[string]tw.HandlerFunc
}

func CreateTreeNode(segment string) *TreeNode {
	return &TreeNode{
		segment:  segment,
		children: make([]*TreeNode, 0, 3),
		handlers: make(map[string]tw.HandlerFunc, 2),
	}
}
