package main

import (
	"fmt"
)

func main() {
	// Defer
	defer fmt.Println("hola")
	fmt.Println("mundo")

	// Continue y break
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 2 {
			fmt.Println("es dos")
			continue
		} else {
			if i == 8 {
				fmt.Println("break")
				break
			}
		}
	}
}
