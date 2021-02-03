package main
import "fmt"
func main(){
	aux := 0
	for true{
		var entrada, cont int
		fmt.Scanln(&entrada)
		if entrada == -1{
			break
		}
		aux += 1
		cont = 0
		for true{
			if entrada <= 1{
				break
			}
			cont += 1
			entrada -= 2
		}
		fmt.Printf("Experiment %d: %d full cycle(s)\n", aux, cont)
	}
}