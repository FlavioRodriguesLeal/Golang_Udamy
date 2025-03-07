package main

import (
	"fmt"
	"reflect"
)

type curso struct {
	nome string
}

func main() {
	var coisa interface{}
	fmt.Println(coisa)
	fmt.Println(reflect.TypeOf(coisa))

	coisa = 3
	fmt.Println(coisa)
	fmt.Println(reflect.TypeOf(coisa))

	type dinamico interface{}

	var coisa2 dinamico = "Opa"
	fmt.Println(coisa2)
	fmt.Println(reflect.TypeOf(coisa2))

	coisa2 = true
	fmt.Println(coisa2)
	fmt.Println(reflect.TypeOf(coisa2))

	coisa2 = curso{"Golang: Explorando a Linguagem do Google"}
	fmt.Println(coisa2)
	fmt.Println(reflect.TypeOf(coisa2))
}
