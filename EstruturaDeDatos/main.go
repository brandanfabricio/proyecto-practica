package main

import (
	lista "EstruturaDeDatos/Lista"
	"fmt"
)

func main() {

	list := lista.NewListPuntero()
	list.Insertar(10, 1)
	list.Insertar(14, 1)
	list.Insertar(13, 1)
	list.Insertar(14, 1)
	list.Insertar(20, 1)
	list.Insertar(100, 1)
	fmt.Println("&&&&&&&&&&&&&&&&&&&&&&")
	fmt.Println(list)
	// list.ImprimirList()

}
