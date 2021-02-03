package main

import "fmt"

func main() {
	var entA, entB, media float64
	fmt.Scanln(&entA)
	fmt.Scanln(&entB)
	media = ((entA * 3.5) + (entB * 7.5)) / 11.0
	fmt.Printf("MEDIA = %.5f\n", media)
}
