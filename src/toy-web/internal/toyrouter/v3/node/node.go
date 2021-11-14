package node

import (
	"fmt"
	"sync"
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

var (
	nodeValues     = make([]int, 5, 5)
	nodeGenerators = make(map[int]NewNodeFunc, 5)
	nodeMatchers   = make(map[int]MatchNodeFunc, 5)
	priority       = &ByValue{}
	mutex          sync.Mutex
)

type INode interface {
	GetSegment() string
	GetValue() int

	GetChildren() []INode
	SetChild(child INode)

	GetAction(method string) (tw.Action, bool)
	SetAction(method string, action tw.Action)

	MatchSegment(segment string) bool
	Match(segment string, ctx tw.IContext) bool
}

type NewNodeFunc func(segment string) INode
type MatchNodeFunc func(segment string) bool

func Register(value int, nodeFunc NewNodeFunc, matchFunc MatchNodeFunc) error {
	mutex.Lock()
	defer mutex.Unlock()
	if _, ok := nodeGenerators[value]; ok {
		return fmt.Errorf("this value [%d] of node has been registered, please change the node value", value)
	}
	nodeValues = append(nodeValues, value)
	nodeGenerators[value] = nodeFunc
	nodeMatchers[value] = matchFunc
	return nil
}

func New(value int, segment string) (INode, error) {
	nodeFunc, ok := nodeGenerators[value]
	if !ok {
		return nil, fmt.Errorf("node value %d not implemented", value)
	}
	return nodeFunc(segment), nil
}

func NewBySegment(segment string) (INode, error) {
	candidates := make([]int, len(nodeValues))
	copy(candidates, nodeValues)
	for len(candidates) > 0 {
		n, ok := priority.GetMostByValue(candidates)
		if !ok {
			return nil, fmt.Errorf("segment [%s] has no matched route strategy", segment)
		}
		matcher, ok := nodeMatchers[n]
		if !ok {
			return nil, fmt.Errorf("segment [%s] has no matched route strategy", segment)
		}
		if matcher(segment) {
			return nodeGenerators[n](segment), nil
		}
		priority.RemoveMostByValue(&candidates)
	}
	return nil, fmt.Errorf("segment [%s] has none matched node type", segment)
}
