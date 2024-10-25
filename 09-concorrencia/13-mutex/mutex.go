package main

import (
	"fmt"
	"sync"
	_ "sync"
)

var (
	m sync.Mutex
)

func somar(i *int, wg *sync.WaitGroup) {
	defer wg.Done()
	m.Lock()
	*i++
	m.Unlock()
}

func main() {
	i := 0
	var wg sync.WaitGroup
	for x := 0; i < 1000000; x++ {
		wg.Add(1)
		somar(&i, &wg)
	}
	wg.Wait()
	fmt.Println(i)
}
