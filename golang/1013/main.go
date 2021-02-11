package main

import (
	"fmt"
	"math"
)

func main() {
	var entA, entB, entC float64
	fmt.Scanln(&entA, &entB, &entC)
	maiorAB := (int(entA) + int(entB) + int(math.Abs(entA-entB))) / 2
	maiorABC := (int(maiorAB) + int(entC) + int(math.Abs(float64(maiorAB)-entC))) / 2
	fmt.Printf("%d eh maior", maiorABC)
}
