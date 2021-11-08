package main

import "fmt"

func main() {
	showCounter()

	fmt.Println()

	printCount()

	fmt.Println()
}

func counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func showCounter() {
	fmt.Println("showCounter(): ")
	c1 := counter()
	c2 := counter()

	for i := 0; i < 10; i++ {
		fmt.Printf("c1: %d\n", c1())
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("c2: %d\n", c2())
	}
}

func printCount() {
	fmt.Println("printCount(): ")
	funcs := make([]func(), 0, 10)
	for i := 0; i < 10; i++ {
		tmpFunc := func() {
			// 此处记录的是i的指针
			// 因此在最终执行的时候每一个i其实都是最后一个i的值
			// 即会输出10个 i=10
			fmt.Printf("cur i = %d\n", i)
		}
		funcs = append(funcs, tmpFunc)
	}

	for _, fun := range funcs {
		fun()
	}
}
