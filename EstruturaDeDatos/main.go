package main

import (
	lista "EstruturaDeDatos/Lista"
	"fmt"
)

func main() {
	Newlista := lista.NewListArray()
	fmt.Println("elmenteo siguiente ", Newlista.Fin())

	for i := 0; i <= 10; i++ {
		numbre := i * 2
		Newlista.Inserta(numbre, i)
	}

	Newlista.Inserta(31, 3)
	Newlista.Inserta(32, 8)
	fmt.Println("Enctrado -> ", Newlista.Localizar(20))
	fmt.Println("Enctrado -> ", Newlista.Localizar(14))
	ultimo := Newlista.Len()
	fmt.Println("tamaño ", ultimo)

	fmt.Println("elmenteo siguiente ", Newlista.Fin())
	fmt.Println("elmenteo  ", Newlista.Recupera(3))

	fmt.Println(Newlista)
	Newlista.Imprimir()
	Newlista.Anular()
	fmt.Println(Newlista)

}
