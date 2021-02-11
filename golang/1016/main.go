package main

import (
	"fmt"
)

func main() {
	var km int
	fmt.Scanln(&km)
	distancia := km * 2
	fmt.Printf("%d minutos\n", distancia)
}
