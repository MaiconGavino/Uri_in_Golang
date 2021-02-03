package main
import "fmt"
func main(){
	var casos int
	fmt.Scanln(&casos)
	for i:= 0; i < casos; i++ {
		var entrada string
		fmt.Scanln(&entrada)
		aux := 0
		soma := 0
		for true{
			if len(entrada) == aux{
				break
			}
			if entrada[aux] == '1'{
				soma = soma + 2
			}
			if entrada[aux] == '2'{
				soma = soma + 5
			}
			if entrada[aux] == '3'{
				soma = soma + 5
			}
			if entrada[aux] == '4'{
				soma = soma + 4
			}
			if entrada[aux] == '5'{
				soma = soma + 5
			}
			if entrada[aux] == '6'{
				soma = soma + 6
			}
			if entrada[aux] == '7'{
				soma = soma + 3
			}
			if entrada[aux] == '8'{
				soma = soma + 7
			}
			if entrada[aux] == '9'{
				soma = soma + 6
			}
			if entrada[aux] == '0'{
				soma = soma + 6
			}
			aux = aux + 1;
		}
		fmt.Printf("%d leds\n",soma)
	}
}