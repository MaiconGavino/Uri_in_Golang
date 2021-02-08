package main

import (
	"fmt"
)

func main() {
	var num, hrs int
	var value float64
	fmt.Scanln(&num)
	fmt.Scanln(&hrs)
	fmt.Scanln(&value)
	salary := value * float64(hrs)
	fmt.Printf("NUMBER = %d\nSALARY = U$ %.2f\n", num, salary)
}
