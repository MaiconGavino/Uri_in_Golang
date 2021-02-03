package main

import "fmt"

func main() {
	var entA, entB int
	var soma int
	fmt.Scanln(&entA)
	fmt.Scanln(&entB)
	soma = entA + entB
	fmt.Printf("SOMA = %d\n", soma)
}
