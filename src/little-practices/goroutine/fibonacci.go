package main

import (
	"fmt"
	"time"
)

func fib(n int, ch chan<- uint64) {
	x, y := uint64(0), uint64(1)
	for i := 0; i < n; i++ {
		x, y = y, x+y
		ch <- x
		time.Sleep(time.Millisecond * 300)
	}
}

func fibonacci() {
	ch := make(chan uint64, 1000)
	go fib(1000, ch)
	c := 0
	for n := range ch {
		fmt.Printf("%d ", n)
		c++
		if c%10 == 0 {
			fmt.Println()
		}
	}
}

func fibonacci2() {
	con, quit := make(chan int), make(chan int)
	go doFib(con, quit)
	go func(quit chan int) {
		time.Sleep(time.Second * 5)
		quit <- 1
	}(quit)

	n := 0
	for true {
		c, ok := <-con
		if !ok {
			break
		}
		fmt.Printf("%d  ", c)
		n++
		if n%10 == 0 {
			fmt.Println()
		}
	}
}

func doFib(con, quit chan int) {
	x, y := 1, 1
	for true {
		select {
		case con <- x:
			{
				x, y = y, x+y
				time.Sleep(time.Millisecond * 500)
			}
		case <-quit:
			{
				fmt.Println("quit")
				close(con)
				return
			}
		}
	}
}
