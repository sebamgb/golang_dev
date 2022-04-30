package main

import (
	"fmt"
)

// Reto condicionales:
//authentication
func auth(u, p string) {
	user := "seba"
	password := "1234"
	if u == user && p == password {
		fmt.Println("access")
	} else {
		fmt.Println("access denied")
	}
}

func par(a uint) {
	fmt.Println(a)
	if a%2 == 0 {
		fmt.Println("es par")
	} else {
		fmt.Print("no es par")
	}
}
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
	//value, err := strconv.Atoi("ahcagxh")
	// nil: no existe error
	//if err != nil {
	//	// log fatal luego de mostrar error termina el programa
	//	log.Fatal(err)
	//} else {
	//	fmt.Println("value:", value)
	//}
	fmt.Println("par?:")
	par(1)
	fmt.Println()
	auth("seba", "1324")
}
