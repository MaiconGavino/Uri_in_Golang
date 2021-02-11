package main

import (
	"fmt"
)

func main() {
	var result, value, parcelas int
	fmt.Scanln(&value)
	fmt.Scanln(&parcelas)

	if value%parcelas == 0 {
		result = value / parcelas
		for i := 0; i < parcelas; i++ {
			fmt.Printf("%d\n", result)
		}
	} else {
		aux := value % parcelas
		result = value / parcelas
		for i := 0; i < aux; i++ {
			fmt.Printf("%d\n", result+1)
		}
		for i := 0; i < (parcelas - aux); i++ {
			fmt.Printf("%d\n", result)
		}
	}

}
