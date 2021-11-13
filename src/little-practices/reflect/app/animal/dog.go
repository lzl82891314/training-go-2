package animal

import "fmt"

type Dog struct {
}

func (d *Dog) SayHello() {
	fmt.Println("weo weo weo")
}

func (d *Dog) Move() {
	fmt.Println("run by dog paws")
}
