package main

import (
	"fmt"
	"golang/src/mypackage"
)

func main() {
	my_pc := mypackage.Pc{
		Ram:   16,
		Disk:  500,
		Brand: "asus",
	}
	fmt.Println(my_pc.String())

}
