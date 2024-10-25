package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 1

	var p *int = nil
	p = &i // pegando o endereço da variável
	*p++   //desreferenciando (pegando o valor)
	i++

	// Go não tem aritmétrica de ponteiro
	// p++

	fmt.Println(p, *p, i, &i)

	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(&i))
	fmt.Println(reflect.TypeOf(p))
	fmt.Println(reflect.TypeOf(*p))
}
