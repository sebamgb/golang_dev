package main

import (
	"fmt"
)

func main() {
	// Arrays:
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))

	// Slices:
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// Metodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])  //último elemento exclusivo
	fmt.Println(slice[2:4]) //primer elemento inclusivo y último exclusivo
	fmt.Println(slice[4:])  //primer elemento inclusivo

	// array inmutables, sin embargo silces no
	// Append
	slice = append(slice, 7)
	fmt.Println(slice)

	// Append nueva lista
	new_slice := []int{8, 9, 10}
	slice = append(slice, new_slice...)
	fmt.Println(slice)
}
