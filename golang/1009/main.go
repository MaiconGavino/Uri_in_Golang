package main

import (
	"fmt"
)

func main() {
	var name string
	var fixo, comissao float64
	fmt.Scanln(&name)
	fmt.Scanln(&fixo)
	fmt.Scanln(&comissao)
	salary := fixo + (comissao * 0.15)
	fmt.Printf("TOTAL = R$ %.2f", salary)
}
