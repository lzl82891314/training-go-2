package main

import "fmt"

func main() {
	p := &Parent{}

	s := &Son{}

	p.ShowName()
	// 坑：如果子类没有这个方法，则go runtime会自动找到父类执行
	// 并且此时，如果调用指针已经指向了父类，那么之后的所有逻辑都会沿着父类的实现完成
	// 因此此处会输出两个 my name is: daddy
	s.ShowName()
}

type Parent struct{}

func (p *Parent) ShowName() {
	fmt.Printf("my name is: %s\n", p.GetName())
}

func (p *Parent) GetName() string {
	return "daddy"
}

type Son struct {
	Parent
}

func (s *Son) GetName() string {
	return "son"
}
