package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Inicio")
	ch := make(chan int, 10)
	fmt.Println("ch criado")
	fmt.Println("inic princ gravar")
	go gravar(ch)
	fmt.Println("fim princ gravar")
	fmt.Println("inic princ consume")
	go consume(ch)
	fmt.Println("fim princ consume")
	time.Sleep(time.Second * 20)
	fmt.Println("FIm")
}

func gravar(ch chan int) {
	fmt.Println("inicio gravar")
	for x := range 10 {
		fmt.Printf("inicio for gravar %d \n", x)
		ch <- x
		fmt.Printf("fim for gravar %d \n", x)
	}
	close(ch)
	fmt.Println("fim gravar")
}

func consume(ch chan int) {
	fmt.Println("inicio consume")
	for x := range ch {
		fmt.Printf("inicio for consume %d \n", x)
		fmt.Println(x)
		fmt.Printf("fim for consume %d \n", x)
	}
	fmt.Println("fim consume")
}
