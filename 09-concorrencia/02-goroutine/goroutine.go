package main

import (
	"fmt"
	"time"
)

func fala(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (interação %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	//fala("Maria", "Porque você não fala comigo?", 3)
	//fala("João", "Só posso falar depois de você!", 1)

	//go fala("Maria", "Porque você não fala comigo?", 500)
	//go fala("João", "Só posso falar depois de você!", 500)

	//time.Sleep(time.Second * 5)
	//fmt.Println("Fim!")

	go fala("Maria", "Entendi!!!", 10)
	fala("João", "Parabéns!", 5)
}
