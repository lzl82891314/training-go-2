package node

import (
	"sort"
)

type INodePriority interface {
	GetMost(candidates []INode) (INode, bool)
	GetMostByValue(candidates []int) (int, bool)
	Sort(candidates []INode)
	SortByValue(candidates []int)
	RemoveMost(candidates *[]INode) bool
	RemoveMostByValue(candidates *[]int) bool
}

type ByValue struct{}

func (p *ByValue) GetMost(candidates []INode) (INode, bool) {
	if l := len(candidates); l != 0 {
		// 按Node权重获取优先级
		// 可以先Sort从小到大排序，然后选最后一个，就是优先级最高的node
		p.Sort(candidates)
		return candidates[l-1], true
	}
	return nil, false
}

func (p *ByValue) GetMostByValue(candidates []int) (int, bool) {
	if l := len(candidates); l != 0 {
		p.SortByValue(candidates)
		return candidates[l-1], true
	}
	return 0, false
}

func (p *ByValue) Sort(candidates []INode) {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].GetValue() < candidates[j].GetValue()
	})
}

func (p *ByValue) SortByValue(candidates []int) {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
}

func (p *ByValue) RemoveMost(candidates *[]INode) bool {
	if candidates == nil {
		return false
	}

	l := len(*candidates)
	if l == 0 {
		return false
	}
	if l == 1 {
		can := make([]INode, 0)
		candidates = &can
		return true
	}

	p.Sort(*candidates)
	*candidates = (*candidates)[0 : l-1]
	return true
}

func (p *ByValue) RemoveMostByValue(candidates *[]int) bool {
	if candidates == nil {
		return false
	}

	l := len(*candidates)
	if l == 0 {
		return false
	}
	if l == 1 {
		can := make([]int, 0)
		candidates = &can
		return true
	}

	p.SortByValue(*candidates)
	*candidates = (*candidates)[0 : l-1]
	return true
}
