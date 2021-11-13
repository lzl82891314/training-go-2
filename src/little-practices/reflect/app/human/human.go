package human

import "fmt"

type Human struct {
}

func (h *Human) SayHello() {
	fmt.Println("hello")
}

func (h *Human) Move() {
	fmt.Println("walk by feet")
}
