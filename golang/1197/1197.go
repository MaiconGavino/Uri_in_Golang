package main
import (
		"fmt"
		"io"
)
func main(){
	for{
		var veloc, temp, result int
		_,err:=fmt.Scanln(&veloc, &temp)
		if err == io.EOF{
			break
		}
		result = veloc * (temp * 2)
		fmt.Printf("%d\n",result)
	}
}