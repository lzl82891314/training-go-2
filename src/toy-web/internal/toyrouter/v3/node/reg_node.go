package node

import (
	"log"
	"regexp"
	"strings"
	tw "toy-web"
)

func init() {
	err := Register(Reg, newRegNode, isRegNode)
	if err != nil {
		log.Fatalln(err)
	}
}

const (
	RegPrefix = "["
	RegSuffix = "]"
)

type regNode struct {
	segment  string
	children []INode
	handlers map[string]tw.Action
	value    int
}

func newRegNode(pattern string) INode {
	return &regNode{
		segment:  pattern,
		children: nil, // 正则表达式也不应该有子节点
		handlers: make(map[string]tw.Action, 2),
		value:    Reg,
	}
}

func isRegNode(segment string) bool {
	// 我们可以定义一种规则：正则表达式是以 [ 开始以 ] 结束的
	// 这样做就不用再暴露一个函数专门注册正则表达式
	if !strings.HasPrefix(segment, RegPrefix) || !strings.HasSuffix(segment, RegSuffix) {
		return false
	}
	compile, err := regexp.Compile(segment)
	return compile != nil && err == nil
}

func (n *regNode) GetSegment() string {
	return n.segment
}

func (n *regNode) GetValue() int {
	return n.value
}

func (n *regNode) GetChildren() []INode {
	return n.children
}

func (n *regNode) SetChild(child INode) {
	for _, v := range n.children {
		if v.GetSegment() == child.GetSegment() {
			return
		}
	}
	n.children = append(n.children, child)
}

func (n *regNode) GetAction(method string) (tw.Action, bool) {
	action, ok := n.handlers[strings.ToUpper(method)]
	return action, ok
}

func (n *regNode) SetAction(method string, action tw.Action) {
	n.handlers[strings.ToUpper(method)] = action
}

func (n *regNode) MatchSegment(segment string) bool {
	return isRegNode(segment)
}

func (n *regNode) Match(segment string, ctx tw.IContext) bool {
	matched, err := regexp.MatchString(n.segment, segment)
	if err != nil {
		log.Printf("router segment [%s] matching reg [%s] error [%s]", segment, n.segment, err.Error())
		return false
	}
	return matched
}
