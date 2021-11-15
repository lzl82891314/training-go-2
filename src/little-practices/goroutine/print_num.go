package main

import (
	"fmt"
	"time"
)

func print(base, times int) {
	time.Sleep(100 * time.Millisecond)
	num := base * times
	if times == 10 {
		fmt.Printf("%d\n", num)
	} else {
		fmt.Printf("%d  ", num)
	}
}

func printNum() {
	for i := 1; i <= 10; i++ {
		for j := 1; j < 10; j++ {
			go print(j, i)
		}
		print(10, i)
	}
}
