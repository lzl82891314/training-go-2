package main

import (
	"fmt"
	"sync"
)

var (
	wg   sync.WaitGroup
	once sync.Once
)

func f1(chan1 chan<- int) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		chan1 <- i
	}
	close(chan1)
}

func f2(chan1 <-chan int, chan2 chan<- int) {
	defer wg.Done()
	for {
		i, ok := <-chan1
		if !ok {
			break
		}
		chan2 <- i * i
	}
	once.Do(func() { close(chan2) })
}

func chanComm() {
	chan1 := make(chan int, 100)
	chan2 := make(chan int, 100)
	wg.Add(2)
	go f1(chan1)
	go f2(chan1, chan2)
	wg.Wait()
	for result := range chan2 {
		fmt.Println(result)
	}
}
