package main

import "fmt"

func main() {
	samplePractice()
}

func samplePractice() {
	fmt.Println("hello")

	go func() {
		fmt.Println("world")
	}()
}
