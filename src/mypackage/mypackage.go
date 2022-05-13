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
	Altura     float64
	Base_menor float64
	Base_mayor float64
	//area
	sum   float64
	mitad float64
}

// sacar area para el trapecio
//-declaración de variable para función
var T Trapecio

func (T *Trapecio) area() float64 {
	T.sum = T.Base_menor + T.Base_mayor
	T.mitad = T.sum / 2
	return T.mitad * T.Altura
}

//-striger para trapecio
func (T Trapecio) string() string {
	return fmt.Sprintf("El trapecio con altura %v, base menor %v y base mayor %v tiene area ", T.Altura, T.Base_menor, T.Base_mayor)
}

// circulo para el reto

type Circulo struct {
	Diametro float64
	//area
	radio          float64
	cuadrado_radio float64
	pi             float64
}

// sacar area al circulo
//-declaración de variable usada en la función
var C Circulo

func (C *Circulo) area() float64 {
	C.pi = 3.144142
	C.radio = C.Diametro / 2
	C.cuadrado_radio = C.radio * C.radio
	return C.pi * C.cuadrado_radio
}

//-stringer para circulo
func (C Circulo) string() string {
	return fmt.Sprintf("El circulo de diametro %v tiene area ", C.Diametro)
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

// ---o---

// paquete para ejemplificar interfaces

type figuras2d interface {
	area() float64
}

type Cuadrado struct {
	Base float64
}

type Rectangulo struct {
	Base   float64
	Altura float64
}

func (C Cuadrado) area() float64 {
	return C.Base * C.Base
}

func (R Rectangulo) area() float64 {
	return R.Base * R.Altura
}

func Calcular(f figuras2d) {
	fmt.Println("Area:", f.area())
}

// ---o---
// ===== interfaces al reto de las areas =====

type figuras interface {
	area() float64
	string() string
}

func Cal(f figuras) {
	fmt.Print(f.string())
	fmt.Println(f.area())
}
