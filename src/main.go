package main

import (
	"fmt"
	"golang/src/mypackage"
)

func main() {
	my_square := mypackage.Cuadrado{
		Base: 2,
	}

	my_rectangle := mypackage.Rectangulo{
		Base:   2,
		Altura: 4,
	}
	mypackage.Calcular(my_square)

	mypackage.Calcular(my_rectangle)

	// lista de interfaces
	//-"slice" con interfaces para una lista con multiples tipos de datos
	my_interface := []interface{}{"hello", 3, 4.5, true}

	fmt.Println(my_interface...)

}
