package main

import "fmt"

func main() {
	var entA, entB, entC, media float64
	fmt.Scanf("%f", &entA)
	fmt.Scanf("%f", &entB)
	fmt.Scanf("%f", &entC)
	media = ((entA * 2) + (entB * 3) + (entC * 5)) / 10
	fmt.Printf("MEDIA = %.1f\n", media)
}
