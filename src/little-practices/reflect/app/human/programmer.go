package human

import "fmt"

type Programmer struct {
}

func (p *Programmer) SayHello() {
	fmt.Println("hello world")
}

func (p *Programmer) Move() {
	fmt.Println("walk by fingers")
}
