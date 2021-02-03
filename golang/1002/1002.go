package main

import "fmt"

func main() {
	var raio float64
	var result float64
	const pi = 3.14159
	fmt.Scanln(&raio)
	result = pi * (raio * raio)
	fmt.Printf("A=%.4f\n", result)
}
