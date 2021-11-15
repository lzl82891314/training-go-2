package main

import (
	"fmt"
	"time"
)

func makeNum(ch chan<- int) {
	time.Sleep(5 * time.Second)
	ch <- 10
}

func timeout(ch chan<- int) {
	time.Sleep(3 * time.Second)
	ch <- 0
}

func chanBlock() {
	ch := make(chan int, 1)
	timeoutCh := make(chan int, 1)
	go makeNum(ch)
	go timeout(timeoutCh)
	select {
	case <-ch:
		fmt.Println(ch)
	case <-timeoutCh:
		fmt.Println("timeout")
	}
}

func simpleBlock() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	ch <- 3
}

func bufferBlock() {
	ch := make(chan int, 100)
	go func(ch chan<- int) {
		for true {
			ch <- time.Now().Nanosecond()
		}
	}(ch)
	time.Sleep(time.Second)
	for true {
		fmt.Println(<-ch)
		time.Sleep(100 * time.Millisecond)
	}
}
