package mypackage

import "fmt"

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
