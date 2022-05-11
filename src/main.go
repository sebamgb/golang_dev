package main

import (
	"fmt"
)

type pc struct {
	ram   int
	disk  int
	brand string
}

// para acceder a la lectura de los valores en un struct antes del nombre de la función se le agreda un atributo con el tipo de dato del struct al que se quiere entrar
func (my_pc pc) ping() {
	fmt.Println(my_pc.brand, "pong")
}

// para modificar valores de un struct, con una función para optimizar codigo, se debe acceder al valor del struct delimitandolo como puntero con el asterisco
func (my_pc *pc) duplicateRam() {
	my_pc.ram = my_pc.ram * 2

}

func main() {
	// declaración normal de varible
	a := 50
	// b apunta a la dirección de memoria de a osea es el puntero de a
	b := &a

	fmt.Println(b)
	fmt.Println(*b) // con el '*' apuntamos a el valor de a

	*b = 100
	fmt.Println(a) // al madificar el puntero del volor de a tambien modificamos a

	my_pc := pc{ram: 16, disk: 500, brand: "asus"}
	fmt.Println(my_pc)

	// llamamos a la funcion que se mete al struct
	my_pc.ping()

	fmt.Println(my_pc)
	my_pc.duplicateRam()

	fmt.Println(my_pc)
}
