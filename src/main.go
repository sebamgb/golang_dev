package main

import (
	"fmt"
	pk "golang/src/mypackage"
)

func main() {
	// instanciando struct publico
	var my_car pk.CarPublic
	// struct inaccesible
	//var my_car pk.carPrivate

	fmt.Println(my_car)

	pk.PrintMessge("Hola Sebastián")

	// función inaccesible
	//pk.printMessage("hola")
}
