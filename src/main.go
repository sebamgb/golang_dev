package main

import "fmt"

func trapecioArea(b, B, h float64) float64 {
	altura := h
	sum := b + B
	mitad := sum / 2
	return mitad * altura
}

func circuloArea(a float64) float64 {
	pi := 3.144142
	radio := a / 2
	cuadrado_radio := radio * radio
	area := pi * cuadrado_radio
	return area
}

func main() {
	fmt.Printf("Area trapecio: %v\n", trapecioArea(3.5, 9.5, 4.1))
	fmt.Printf("Area circulo: %v\n", circuloArea(4.5))

}
