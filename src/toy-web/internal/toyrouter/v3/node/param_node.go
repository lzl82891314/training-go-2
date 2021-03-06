package node

import (
	"log"
	"strings"
	tw "toy-web"
)

func init() {
	err := Register(Param, newParamNode, isParamNode)
	if err != nil {
		log.Fatalln(err)
	}
}

type paramNode struct {
	segment  string
	children []INode
	handlers map[string]tw.Action
	value    int
}

func newParamNode(segment string) INode {
	return &paramNode{
		segment:  segment,
		children: make([]INode, 0, 3),
		handlers: make(map[string]tw.Action, 2),
		value:    Param,
	}
}

func isParamNode(segment string) bool {
	segment = strings.Trim(segment, " ")
	if segment == "" {
		return false
	}
	return strings.HasPrefix(segment, ParamSymbol)
}

func (n *paramNode) GetSegment() string {
	return n.segment
}

func (n *paramNode) GetValue() int {
	return n.value
}

func (n *paramNode) GetChildren() []INode {
	return n.children
}

func (n *paramNode) SetChild(child INode) {
	for _, v := range n.children {
		if v.GetSegment() == child.GetSegment() {
			return
		}
	}
	n.children = append(n.children, child)
}

func (n *paramNode) GetAction(method string) (tw.Action, bool) {
	action, ok := n.handlers[strings.ToUpper(method)]
	return action, ok
}

func (n *paramNode) SetAction(method string, action tw.Action) {
	n.handlers[strings.ToUpper(method)] = action
}

func (n *paramNode) MatchSegment(segment string) bool {
	return isParamNode(segment)
}

func (n *paramNode) Match(segment string, ctx tw.IContext) bool {
	if segment == WildcardSymbol {
		// 通配符无法匹配路径参数
		return false
	}
	p := n.segment[1:]
	if ctx != nil {
		ctx.SetPathParam(p, segment)
	}
	return true
}
