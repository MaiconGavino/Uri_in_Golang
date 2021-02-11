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

	fmt.Printf("TRIANGULO: %.3f\nCIRCULO: %.3f\nTRAPEZIO: %.3f\nQUADRADO: %.3f\nRETANGULO: %.3f\n", areaTri, areaCir, areaTra, areaQua, areaRet)
}
