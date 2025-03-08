package main

import (
	lista "EstruturaDeDatos/Lista"
	"fmt"
)

func main() {
	list := lista.NewListArray()

	list.Inserta(2, 0)
	list.Inserta(5, 0)
	// list.Inserta(7, 0)
	// list.Inserta(6, 3, *list)
	// list.Inserta(8, 2, *list)

	// data := []interface{}{}

	fmt.Println(list)
}
