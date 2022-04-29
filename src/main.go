package main

import "fmt"

func main() {
	// For condicional
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// For while
	counter := 1
	for counter <= 10 {
		fmt.Println("\n", counter)
		counter++
	}

	// For forever
	//counter_forever := 1
	//for {
	//	fmt.Println("\n\n", counter_forever)
	//	counter_forever++
	//}

	// Reto con for condicional

	for c := 10; c >= 1; c-- {
		fmt.Println(c)
	}
}
