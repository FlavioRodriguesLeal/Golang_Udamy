package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	ch <- 1           //enviando dados  para o canal (escrita)
	ch <- 2           //enviando dados  para o canal (escrita)
	ch <- 3           //enviando dados  para o canal (escrita)
	fmt.Println(<-ch) //recebendo dados do canal (leitura)
	fmt.Println(<-ch) //recebendo dados do canal (leitura)
	fmt.Println(<-ch) //recebendo dados do canal (leitura)

	ch <- 2
	fmt.Println(<-ch)
}
