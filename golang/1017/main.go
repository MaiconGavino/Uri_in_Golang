package main

import (
	"fmt"
)

func main() {
	var temp, veloc int
	fmt.Scanln(&temp)
	fmt.Scanln(&veloc)
	result := float64(temp*veloc) / 12
	fmt.Printf("%.3f", result)
}
