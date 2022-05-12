package main

import (
	"fmt"
	"golang/src/mypackage"
)

func main() {
	fmt.Printf("Area trapecio: %v\n", mypackage.T.TrapecioArea(3.5, 9.5, 4.0))
	fmt.Printf("Area circulo: %v\n", mypackage.C.CirculoArea(4.5))

}
