package main

import (
	"fmt"
)

func main() {
	// condiciones anidadas con switch
	switch mod := 2 % 2; mod {
	case 0:
		fmt.Println("es par")
	default:
		fmt.Println("no es par")
	}
	//sin condicion
	value := 100
	switch {
	case value < 100:
		fmt.Println("es menor a 100")
	case value > 100:
		fmt.Println("es mayor a 100")
	default:
		fmt.Println("es 100")
	}
}
