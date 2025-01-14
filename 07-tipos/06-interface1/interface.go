package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}
type produto struct {
	nome  string
	preco float64
}

//interface são implementadas implicitamente

func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var coisa imprimivel = pessoa{"Roberto", "Silva"}
	imprimir(coisa)

	coisa = produto{"Calça Jeans", 79.90}
	imprimir(coisa)

	p2 := produto{"Notebool", 20000.00}
	imprimir(p2)
}
