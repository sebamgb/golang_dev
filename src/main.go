package main

import (
	"fmt"
)

type trapecio struct {
	altura     float64
	base_menor float64
	base_mayor float64
}

func (t *trapecio) trapecioArea() float64 {
	sum := t.base_menor + t.base_mayor
	mitad := sum / 2
	return mitad * t.altura
}

type circulo struct {
	pi       float64
	diametro float64
}

func (c *circulo) circuloArea() float64 {
	radio := c.diametro / 2
	cuadrado_radio := radio * radio
	area := c.pi * cuadrado_radio
	return area
}

func main() {
	c := circulo{pi: 3.144142, diametro: 4.5}
	t := trapecio{altura: 4.0, base_menor: 3.5, base_mayor: 9.5}
	fmt.Printf("Area trapecio: %v\n", t.trapecioArea())
	fmt.Printf("Area circulo: %v\n", c.circuloArea())

	fmt.Println(my_pc)
}
