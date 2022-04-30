package main

import (
	"fmt"
)

func main() {
	// Map
	n := make(map[string]int)
	n["jose"] = 14
	n["pepito"] = 13

	fmt.Println(n)

	// Recorrido de map
	for i, v := range n {
		fmt.Println(i, v)
	}

	// Encontrar valor
	value := n["jose"]
	fmt.Println(value)

	value = n["josep"]
	fmt.Println(value) //si la llave no existe imprime el zero value del tipo de dato del valor

	// existe?
	value, ok := n["jose"]
	fmt.Println(value, ok) //imprime el valor(value) y el booleano correspondiente a la existencia
}
