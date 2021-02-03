package main
import "fmt"
func main(){
	var entA, entB, prod int
	fmt.Scanln(&entA)
	fmt.Scanln(&entB)
	prod = entA*entB
	fmt.Printf("PROD = %d\n",prod)
}