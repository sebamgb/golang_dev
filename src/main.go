package main

import "fmt"

func  main()  {
	// declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.14
	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// declaracion de variables enteras
	// base:= 14
	// var altura int = 13
	// var area int

	// zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Area cuadrado
	var base_cuadrado int
	area := base_cuadrado * base_cuadrado

	fmt.Println(area)

	x:= 50
	y:=10

	// suma
	result :=x+y

	fmt.Println("suma:", result)

	// resta
	result = x - y
	
	fmt.Println("resta:", result)

	// multiplicación
	result = x*y
	
	fmt.Println("multiplicación:", result)

	// división
	result =x /y

	fmt.Println("división:", result)

	// modulo
	result = x % y

	fmt.Println("modulo:", result)

	// incremental
	x++

	fmt.Println("incremental:", x)

	// decremetal
	x--

	fmt.Println("decremental:", x)

	// Retos:
	//     - area trapecio y circulo
}