package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	ProducerCount = 3
	ConsumerCount = 5
	FullCount     = 15
	TimeFactor    = 5
)

var (
	waitGroup = sync.WaitGroup{}
)

// 创建 3 个生产者，5个消费者模型
// 每个消费者需要消耗满 100个产品才能满足

func producer(n int, ch chan<- int) {
	defer waitGroup.Done()
	times := createFactor()
	asleepTimes := 0
	for true {
		p := createFactor()
		select {
		case ch <- p:
			{
				t := time.Duration(times) * time.Second
				fmt.Printf("Producer: %d produced a %d, then will sleep %d s\n", n, p, times)
				time.Sleep(t)
			}
		default:
			{
				time.Sleep(time.Second * 3)
				asleepTimes++
				fmt.Println("I need consumers!")
				if asleepTimes == 3 {
					fmt.Printf("Producer %d will go home\n", n)
					return
				}
			}
		}
	}
}

func consumer(n int, ch chan int) {
	waitGroup.Done()
	s := make([]int, 0, FullCount)
	times := createFactor()
	for len(s) < FullCount {
		select {
		case c := <-ch:
			{
				s = append(s, c)
				fmt.Printf("Consumer: %d consume a %d, remains %d, then will sleep %d s\n", n, c, FullCount-len(s), times)
				time.Sleep(time.Duration(times) * time.Second)
			}
		default:
			{
				fmt.Println("Producers need to hurry up, I'm hungry!")
				time.Sleep(time.Second * 3)
			}
		}
	}
	fmt.Printf("Consumer: %d already full\n", n)
}

func createFactor() int {
	times := 0
	for times == 0 {
		times = rand.Intn(TimeFactor)
	}
	return times
}

func pc() {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan int, FullCount)

	waitGroup.Add(ProducerCount)
	for i := 0; i < ProducerCount; i++ {
		go producer(i, ch)
	}

	waitGroup.Add(ConsumerCount)
	for i := 0; i < ConsumerCount; i++ {
		go consumer(i, ch)
	}

	waitGroup.Wait()
}
