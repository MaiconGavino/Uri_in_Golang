package main

import (
	"fmt"
	"math"
)

func main() {
	var entA, entB, entC float64
	fmt.Scanln(&entA, &entB, &entC)
	areaTri := entA * entC / 2.0
	areaCir := 3.14159 * math.Pow(entC, 2)
	areaTra := (entA + entB) * entC / 2.0
	areaQua := math.Pow(entB, 2)
	areaRet := entA * entB

	fmt.Printf("TRIANGULO: %.3f\n", areaTri)
	fmt.Printf("CIRCULO: %.3f\n", areaCir)
	fmt.Printf("TRAPEZIO: %.3f\n", areaTra)
	fmt.Printf("QUADRADO: %.3f\n", areaQua)
	fmt.Printf("RETANGULO: %.3f\n", areaRet)
}
