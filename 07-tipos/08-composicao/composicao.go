package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
	//pode adicionar mais méstodos
}

type bmw7 struct {
}

func (b bmw7) fazerBaliza() {
	fmt.Println("Fazer ...")
}

func (b bmw7) ligarTurbo() {
	fmt.Println("Turbo ...")
}

func main() {
	var b esportivoLuxuoso = bmw7{}
	b.fazerBaliza()
	b.ligarTurbo()
}
