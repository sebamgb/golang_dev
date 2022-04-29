package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// Condicional if
	v := 2
	if v == 1 {
		fmt.Println("es:", v)
	} else {
		fmt.Println("distinto a:", v)
	}
	// with and
	v2 := 3
	if v == 2 && v2 != 3 {
		fmt.Println("Los dos se cumplen")
	} else {
		fmt.Println("uno o los dos no se cumplen")
	}
	// with or
	if v != 2 || v2 == 2 {
		fmt.Println("Uno o los dos se cumplen")
	} else {
		fmt.Println("No se cumple ninguno")
	}
	// ->Manejo de errores con if
	// Convertir texto a numero
	value, err := strconv.Atoi("jajjaja")
	// nil: no existe error
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("value:", value)
	}
}
