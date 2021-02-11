package main

import (
	"fmt"
	"math"
)

func main() {
	var entX1, entX2, entY1, entY2 float64
	fmt.Scanln(&entX1, &entY1)
	fmt.Scanln(&entX2, &entY2)
	distancia := math.Sqrt(math.Pow((entX2-entX1), 2) + math.Pow((entY2-entY1), 2))
	fmt.Printf("%.4f", distancia)
}
