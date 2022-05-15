package main

import (
	"fmt"
)

func m(t string, c chan string) {
	c <- t
}
func main() {
	c := make(chan string, 2)
	c <- "mensaje1"
	c <- "mensaje2"
	//==range y close==
	//cerramos el channel con close cuando esta lleno siendo una buena practica que optimiza más
	close(c)
	fmt.Println(len(c), cap(c))
	//-una vez cerrado no se pueden agregar más datos aunque le sobre espacio
	//c<- "mensaje3"

	//recorremos el channel en busca de los datos
	//-al no cerrar el channel antes de recorrer se siguen esperando datos
	for message := range c {
		fmt.Println(message)
	}

	//==select==
	email1 := make(chan string)
	email2 := make(chan string)
	go m("mensaje1", email1)
	go m("mensaje2", email2)

	//es bueno saber cuantos canales estas usando para la condicional del for
	for i := 0; i < 2; i++ {
		//select muestra dos channels completamente independientes
		//no hay certesa de cual va a mostrar, pero si enumeramos los nombres de los channels va a tomar los channels de forma descendente, es decir los numeros mas altos primero
		select {
		case m1 := <-email1:
			fmt.Println("email recibido de:", m1)
		case m2 := <-email2:
			fmt.Println("email recibido de:", m2)
		}
	}
}
