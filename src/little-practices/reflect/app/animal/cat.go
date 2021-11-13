package animal

import "fmt"

type Cat struct {
}

func (c *Cat) SayHello() {
	fmt.Println("meow meow meow")
}

func (c *Cat) Move() {
	fmt.Println("run by cat paws")
}
