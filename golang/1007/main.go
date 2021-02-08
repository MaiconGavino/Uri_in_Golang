package main

import (
	"fmt"
)

func main() {
	var ent1, ent2, ent3, ent4 int
	fmt.Scanln(&ent1)
	fmt.Scanln(&ent2)
	fmt.Scanln(&ent3)
	fmt.Scanln(&ent4)
	saida := (ent1 * ent2) - (ent3 * ent4)
	fmt.Println("DIFERENCA =", saida)
}
