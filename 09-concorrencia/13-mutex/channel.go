package main

import (
	"fmt"
	"sync"
)

func somar(c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	i := <-c
	i++
	c <- i
}

func main() {
	var wg sync.WaitGroup
	c := make(chan int, 1)
	c <- 0
	for x := 0; x < 10000000; x++ {
		wg.Add(1)
		go somar(c, &wg)
		// go func() {
		// 	defer wg.Done()
		// 	i := <-c
		// 	i++
		// 	c <- i
		// }()
	}
	wg.Wait()
	close(c)
	fmt.Println(<-c)
}
