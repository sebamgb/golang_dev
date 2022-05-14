package main

import (
	"fmt"
)

// ejemplificando channels otra forma más usada de organizar las goroutines

//-creamos una función
//-con un parametro para un channel que recibe como entrada strings
//-con afán de hacer más eficiente el código le indicamos si el channel va a ser de entrada o salida
//entrada: chan<- typedata
//salida: <-chan typedata
func say(t string, c chan<- string) {
	//ingresando string t en canal
	//salida: t = <- chan
	c <- t
}

func main() {
	//--declaramos un channel--
	//-channel que maneja tipos de datos string y recibe una goroutine a la vez, si no se le indica el número de goroutines que maneje va a elegir uno dinamicamente
	c := make(chan string, 1)
	fmt.Println("hello")
	go say("bye", c)
	//-imprimimos la salida del dato en el canal
	fmt.Println(<-c)
}

// primitivos o wait group para una optimizacion en el codigo
// chanels más livianos de codigo para tareas que no necesiten tanta optimización
