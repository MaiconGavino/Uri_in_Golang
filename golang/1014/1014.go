package main
import "fmt"
func main(){
	var distancia int
	var combustivel, consumo  float64
	fmt.Scanln( &distancia)
	fmt.Scanln( &combustivel)
	consumo = float64(distancia) / combustivel
	fmt.Printf("%.3f km/l\n", consumo)
}