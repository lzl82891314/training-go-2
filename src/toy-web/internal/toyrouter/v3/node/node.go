package node

import (
	"fmt"
	"sort"
	tw "toy-web"
)

const (
	// Root 根节点
	Root = iota

	// Wildcard 通配符 *
	Wildcard

	// Param 路径参数 :
	Param

	// Reg 正则表达式
	Reg

	// Static 静态路由
	Static
)

const (
	RootSymbol     = "/"
	WildcardSymbol = "*"
	ParamSymbol    = ":"
)

type INode interface {
	GetSegment() string
	GetValue() int

	GetChildren() []INode
	SetChild(child INode)

	GetAction(method string) tw.Action
	SetAction(method string, action tw.Action)

	MatchSegment(segment string) bool
	Match(segment string, ctx tw.IContext) bool
}

func NewNode(value int, segment string) (INode, error) {
	switch value {
	case Static:
		return newStaticNode(segment), nil
	case Reg:
		return newRegNode(segment), nil
	case Param:
		return newParamNode(segment), nil
	case Wildcard:
		return newWildcardNode(), nil
	case Root:
		return newRootNode(), nil
	}
	return nil, fmt.Errorf("node value %d not implemented", value)
}

func NewNodeBySegment(segment string) (INode, error) {
	if isStaticNode(segment) {
		return NewNode(Static, segment)
	} else if isRegNode(segment) {
		return NewNode(Reg, segment)
	} else if isParamNode(segment) {
		return NewNode(Param, segment)
	} else if isWildcardNode(segment) {
		return NewNode(Wildcard, segment)
	} else if isRootNode(segment) {
		return NewNode(Root, segment)
	}
	return nil, fmt.Errorf("segment [%s] has none matched node type", segment)
}

func Sort(candidates []INode) {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].GetValue() < candidates[j].GetValue()
	})
}
