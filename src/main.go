package main

import "fmt"

func  main()  {
	// Retos:
	//     - area trapecio y circulo

	// Area trapecio

	base_menor := 3.5
	Base_mayor := 9.5
	altura := 4.0

	paso1:= base_menor+Base_mayor
	paso2:= paso1 / 2
	area := paso2 * altura

	fmt.Println("Area trapecio:", area)


	// Area circulo

	pi := 3.144142
	diametro := 8.0
	radio := diametro / 2

	radio_cuadrado := radio *radio
	area = pi * radio_cuadrado

	fmt.Println("Area circulo:", area, "^2")

}