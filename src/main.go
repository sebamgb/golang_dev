package main

import "fmt"

type car struct {
	brand string
	year  int
}

func main() {
	// Instanciado
	my_car := car{brand: "ford", year: 2022}
	fmt.Println(my_car)

	// Otra manera
	var another_car car
	another_car.brand = "ferrari"
	fmt.Println(another_car) //al no especificar todo el contenido del struct imrime los zero values del tipo de dato
}
