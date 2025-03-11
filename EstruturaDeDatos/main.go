package main

import (
	lista "EstruturaDeDatos/Lista"
	"fmt"
)

func main() {
	Newlista := lista.NewListPuntero()

	Newlista.Insertar(10, 1)

	fmt.Println(Newlista)

}
