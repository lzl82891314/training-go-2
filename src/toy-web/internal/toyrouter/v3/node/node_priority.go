package node

import "sort"

type INodePriority interface {
	Get(candidates []INode) (INode, bool)
}

type ByValue struct{}

func (p *ByValue) Get(candidates []INode) (INode, bool) {
	if l := len(candidates); l != 0 {
		// 按Node权重获取优先级
		// 可以先Sort从小到大排序，然后选最后一个，就是优先级最高的node
		Sort(candidates)
		return candidates[l-1], true
	}
	return nil, false
}

func Sort(candidates []INode) {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].GetValue() < candidates[j].GetValue()
	})
}
