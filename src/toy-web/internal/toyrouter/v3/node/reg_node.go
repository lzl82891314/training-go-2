package node

import (
	"log"
	"regexp"
	"strings"
	tw "toy-web"
)

const (
	RegPrefix = "["
	RegSuffix = "]"
)

type RegNode struct {
	segment  string
	children []INode
	handlers map[string]tw.Action
	value    int
}

func newRegNode(pattern string) INode {
	return &RegNode{
		segment:  pattern,
		children: make([]INode, 0, 3),
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

func (n *RegNode) GetSegment() string {
	return n.segment
}

func (n *RegNode) GetValue() int {
	return n.value
}

func (n *RegNode) GetChildren() []INode {
	return n.children
}

func (n *RegNode) SetChild(child INode) {
	for _, v := range n.children {
		if v.GetSegment() == child.GetSegment() {
			return
		}
	}
	n.children = append(n.children, child)
}

func (n *RegNode) GetAction(method string) tw.Action {
	return n.handlers[method]
}

func (n *RegNode) SetAction(method string, action tw.Action) {
	n.handlers[method] = action
}

func (n *RegNode) MatchSegment(segment string) bool {
	return isRegNode(segment)
}

func (n *RegNode) Match(segment string, ctx tw.IContext) bool {
	matched, err := regexp.MatchString(n.segment, segment)
	if err != nil {
		log.Printf("router segment [%s] matching reg [%s] error [%s]", segment, n.segment, err.Error())
		return false
	}
	return matched
}
