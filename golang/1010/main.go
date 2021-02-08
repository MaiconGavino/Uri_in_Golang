package main

import (
	"fmt"
)

func main() {
	var code1, qtd1, code2, qtd2 int
	var value1, value2 float64
	fmt.Scanln(&code1, &qtd1, &value1)
	fmt.Scanln(&code2, &qtd2, &value2)
	valor := (float64(qtd1) * value1) + (float64(qtd2) * value2)
	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", valor)
}
