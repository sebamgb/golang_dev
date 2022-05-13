package main

import (
	"golang/src/mypackage"
)

func main() {
	my_trapecio := mypackage.Trapecio{
		Base_menor: 3.5,
		Base_mayor: 9.5,
		Altura:     4,
	}
	my_circulo := mypackage.Circulo{
		Diametro: 4,
	}

	mypackage.Cal(&my_trapecio)
	mypackage.Cal(&my_circulo)
}
