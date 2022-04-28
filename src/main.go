package main

import "fmt"

func main() {
	// Declaración de variables
	string_hello := "Hello"
	string_world := "world"

	// Println
	fmt.Println(string_hello, string_world)
	fmt.Println(string_hello, string_world)

	// Printf
	nombre := "Platzi"
	cursos := 500

	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos)

	//Sprintf
	message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos)
	fmt.Println(message)

	// Tipos de datos
	fmt.Printf("string_hello %T\n", string_hello)
	fmt.Printf("cursos %T\n", cursos)
}
