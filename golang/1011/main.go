package main

import (
	"fmt"
	"math"
)

func main() {
	var inputRaio float64
	fmt.Scanln(&inputRaio)
	vol := (4 / 3.0) * 3.14159 * math.Pow(inputRaio, 3)
	fmt.Printf("VOLUME = %.3f\n", vol)
}
