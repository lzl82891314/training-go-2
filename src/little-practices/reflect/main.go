package main

import (
	"fmt"
	"reflect"
	"reflect/app"
	"reflect/app/human"
)

func main() {
	hello("jeffery", 28, 100.12, human.Human{}, &human.Programmer{})

	types := reflect.TypeOf(hello)
	values := reflect.ValueOf(hello)
	in := types.NumIn()
	params := make(map[reflect.Value]reflect.Type, in)
	for i := 0; i < in; i++ {
		t := types.In(i)
		v := values.Index(i)
		params[v] = t
	}
	for _, t := range params {
		fmt.Println(t.Name())
	}
}

func hello(name string, age int, weight float32, h human.Human, an app.Animal) {
	fmt.Printf("name: %s, age: %d\n", name, age)
	fmt.Printf("name: %s, age: %d, weight: %f, human: %v, animal: %v\n", name, age, weight, h, an)
}
