package main

import "fmt"

func normalFuncion(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func dobleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFuncion("hello world")
	tripleArgument(1, 2, "hello")
	value := returnValue(2)
	fmt.Println("Value:", value)
	value1, value2 := dobleReturn(2)
	fmt.Println(value1, value2)
	// escape de valores en funciones de retorno
	value1_2, _ := dobleReturn(2)
	fmt.Println(value1_2)

}
