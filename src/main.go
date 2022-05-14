package main

import (
	"fmt"
	"sync"
	"time"
)

// ejemplificando las gp goroutines
func say(t string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(t)
}

func main() {
	//-declaramos un wait group para acumular go routines
	var wg sync.WaitGroup
	fmt.Println("hello")
	wg.Add(1)
	//-empleando la cuncurrencia
	go say("world", &wg)
	//-slución a que se imprima el código con concurrencia no recomendable
	//time.Sleep(time.Second * 1)
	wg.Wait()

	//--funciones anonimas y goroutines--

	go func(t string) {
		fmt.Println(t)
	}("adios")
	//-lo mejor es agragar esta rutina al wait group pero:
	time.Sleep(time.Second * 1)
}
