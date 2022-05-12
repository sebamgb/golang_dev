package mypackage

import (
	"fmt"
)

// CarPublic struct de de acceso publico por la mayúscula
type CarPublic struct {
	Brand string
	Year  int
}

// structs y funciones de acceso privado por comenzar con minúscula
type carPrivate struct {
	brand string
	year  int
}

// PrintMessage imprime un mensaje con acceso público
func PrintMessge(text string) {
	fmt.Println(text)
}

func printMessge(text string) {
	fmt.Println(text)
}

// trapecio para reto

type Trapecio struct {
	altura     float64
	base_menor float64
	base_mayor float64
}

// sacar area para el trapecio
//-declaración de variable para función
var T Trapecio

func (T *Trapecio) TrapecioArea(b, B, h float64) float64 {

	T.base_menor = b
	T.base_mayor = B
	T.altura = h
	sum := b + B
	mitad := sum / 2
	return mitad * h
}

// circulo para el reto

type Circulo struct {
	pi       float64
	diametro float64
}

// sacar area al circulo
//-declaración de variable usada en la función
var C Circulo

func (C *Circulo) CirculoArea(d float64) float64 {

	C.pi = 3.144142
	C.diametro = d
	radio := d / 2
	cuadrado_radio := radio * radio
	area := C.pi * cuadrado_radio
	return area
}

// ===== paquete para ejemplificar stringers =====
type Pc struct {
	Ram   int
	Disk  int
	Brand string
}

var My_pc Pc

func (My_pc Pc) String() string {
	return fmt.Sprintf("tengo %dGB de ram, %dGB de espacio en disco y es marca %s \n", My_pc.Ram, My_pc.Disk, My_pc.Brand)
}
