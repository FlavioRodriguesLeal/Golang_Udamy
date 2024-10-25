package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	//fmt.Println(x / y) //invalid operation: x / y (mismatched types float64 and int)
	fmt.Println(int(x) / y)
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado...
	fmt.Println("Teste " + string(123))

	//int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	//string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true") // OU 1 para true
	if b {
		fmt.Println("Verdadeiro")
	}

	//String pra float
	valor := "3.29"
	res, _ := strconv.ParseFloat(valor, 64)
	fmt.Println(res)

	//Float to string
	var floatNumber float64 = 27.156633168032640

	fmt.Println("as float32 with 'E' (decimal exponent) :", strconv.FormatFloat(floatNumber, 'E', -1, 32))
	fmt.Println("as float64 with 'E' (decimal exponent) :", strconv.FormatFloat(floatNumber, 'E', -1, 64))
	fmt.Println("as float32 with 'f' (no exponent) :", strconv.FormatFloat(floatNumber, 'f', -1, 32))
	fmt.Println("as float64 with 'f' (no exponent) :", strconv.FormatFloat(floatNumber, 'f', -1, 64))
	fmt.Println("with fmt.Sprint :", fmt.Sprint(floatNumber))
	fmt.Println("with fmt.Sprintf :", fmt.Sprintf("%.2f", floatNumber))

}
