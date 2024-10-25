package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//numerico inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))
	fmt.Println("\n")

	//sem sinal(só positivos)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))
	fmt.Println("\n")

	//com sinal... uint8 uint16 uint32 uint64
	i1 := int64(math.MaxInt64)
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))
	fmt.Println("\n")

	var i2 rune = 'a' //representação um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune e", reflect.TypeOf(i2))
	fmt.Println(i2)
	fmt.Println("\n")

	// números reais (float32, flout64)
	//var x float32 = 49.99
	var x = float32(49.99)
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99)) // float64
	fmt.Println("\n")

	//boolean
	bo := true
	fmt.Println("O tipo do bo é", reflect.TypeOf(bo))
	fmt.Println(bo)
	fmt.Println(!bo)
	fmt.Println("\n")

	//String
	s1 := "Olá meu nome é Flávio"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da String é", len(s1))
	fmt.Println("\n")

	//String com multiplas linhas
	s2 := `Olá
	meu
	nome
	é
	Flávio`
	fmt.Println("O tamanho da String é", len(s2))
	fmt.Println(s2)
	fmt.Println("\n")

	//char
	// char := 'a' //não exite
	// fmt.Println("O tipo de char é", reflect.TypeOf(char))//int32
	// fmt.Println("O tipo de char é", char) //97
}
