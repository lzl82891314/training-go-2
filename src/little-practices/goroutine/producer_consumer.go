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
)

var waitGroup = sync.WaitGroup{}

// 创建 3 个生产者，5个消费者模型
// 每个消费者需要消耗满 100个产品才能满足

func producer(n int, ch chan<- int) {
	defer waitGroup.Done()
	rand.Seed(time.Now().UnixNano())
	times := rand.Intn(10)
	asleepTimes := 0
	for true {
		p := rand.Intn(10)
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
				if asleepTimes == 3 {
					fmt.Printf("Producer %d need to go home\n", n)
					return
				}
			}
		}
	}
}

func consumer(n int, ch chan int) {
	waitGroup.Done()
	s := make([]int, 0, FullCount)
	rand.Seed(time.Now().UnixNano())
	times := rand.Intn(10)
	for len(s) < FullCount {
		select {
		case c := <-ch:
			{
				s = append(s, c)
				t := time.Duration(times) * time.Second
				fmt.Printf("Consumer: %d consume a %d, remains %d, then will sleep %d s\n", n, c, FullCount-len(s), times)
				time.Sleep(t)
			}
		default:
			{
				fmt.Println("producer need heer up")
				time.Sleep(time.Second * 3)
			}
		}
	}
	fmt.Printf("Consumer: %d has full\n", n)
}

func pc() {
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
