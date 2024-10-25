package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	m sync.Mutex
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(100)
	go worker(1, ch, &wg)
	go worker(2, ch, &wg)
	go worker(3, ch, &wg)
	go worker(4, ch, &wg)
	go worker(5, ch, &wg)
	go worker(6, ch, &wg)
	go worker(7, ch, &wg)
	go worker(8, ch, &wg)
	go worker(9, ch, &wg)
	go worker(10, ch, &wg)
	go worker(11, ch, &wg)
	go worker(12, ch, &wg)
	go worker(13, ch, &wg)
	go worker(14, ch, &wg)
	go worker(15, ch, &wg)
	go worker(16, ch, &wg)
	go worker(17, ch, &wg)
	go worker(18, ch, &wg)
	go worker(19, ch, &wg)
	go worker(20, ch, &wg)
	for i := 1; i <= 100; i++ {
		ch <- i
	}
	wg.Wait()
	close(ch)
}

func worker(workerId int, ch chan int, wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	for x := range ch {
		m.Lock()
		fmt.Printf("Worker %d received %d \n", workerId, x)
		m.Unlock()
		wg.Done()
	}
}
