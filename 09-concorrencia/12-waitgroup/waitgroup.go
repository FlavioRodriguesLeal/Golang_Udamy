package main

import (
	"fmt"
	"sync"
	"time"
)

func callDatabase(wg *sync.WaitGroup) {
	fmt.Println("startCallDatabase")
	time.Sleep(time.Second)
	fmt.Println("finalCallDatabase")
	wg.Done()
}

func callAPI(wg *sync.WaitGroup) {
	fmt.Println("startCallAPI")
	time.Sleep(time.Second * 2)
	fmt.Println("finalCallAPI")
	wg.Done()
}

func processInternal(wg *sync.WaitGroup) {
	fmt.Println("startProcessInternal")
	time.Sleep(time.Second)
	fmt.Println("finalProcessInternal")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	callDatabase(&wg)
	callAPI(&wg)
	processInternal(&wg)
	wg.Wait()
}
