package main

import (
	"fmt"
)

func rotina(c chan int) {
	fmt.Println("Executou!")
	c <- 1
	fmt.Println("c = 1")
	c <- 2
	fmt.Println("c = 2")
	c <- 3
	fmt.Println("c = 3")
	c <- 4
	fmt.Println("c = 4")
	c <- 5
	fmt.Println("c = 5")
	c <- 6
	fmt.Println("c = 6")
}

func main() {
	ch := make(chan int, 3)
	go rotina(ch)

	fmt.Println(<-ch)
}
