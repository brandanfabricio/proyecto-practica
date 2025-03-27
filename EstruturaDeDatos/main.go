package main

import (
	lista "EstruturaDeDatos/Lista"
	"fmt"
)

func main() {

	list := lista.NewListPuntero()
	fmt.Println(list)
	fmt.Println(list.Fin())

}
