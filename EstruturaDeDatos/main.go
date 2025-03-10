package main

import (
	lista "EstruturaDeDatos/Lista"
	"fmt"
)

func main() {
	list := lista.NewListArray()

	list.Inserta(2, 0)
	list.Inserta(5, 1)
	list.Inserta(7, 2)
	list.Inserta(6, 3)
	list.Inserta(8, 2)

	// data := []interface{}{}
	fmt.Println(list)
	list.Suprime(0)

	fmt.Println(list)
}
