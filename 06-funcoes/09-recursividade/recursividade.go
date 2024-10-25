package main

import "fmt"

func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("Número inválido: %d", n)
	case n == 0:
		return 1, nil

	default:
		fatorialAnterior, _ := fatorial(n - 1)
		fmt.Printf("Valor de n: %d, valor fatorialAnterior: %d, Resultado: %d \n", n, fatorialAnterior, (n * fatorialAnterior))
		return n * fatorialAnterior, nil
	}
}

func main() {
	resultado, _ := fatorial(5)
	fmt.Println(resultado)

	_, erro := fatorial(-4)
	if erro != nil {
		fmt.Println(erro)
	}
}
