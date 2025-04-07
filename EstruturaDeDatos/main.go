package main

import (
	lista "EstruturaDeDatos/Lista"
	"fmt"
)

func main() {

	nuevaLista := &lista.Tipo_Lista{}

	lista.Insertar(20, nuevaLista)
	lista.Insertar(10, nuevaLista)
	lista.Insertar(40, nuevaLista.Sig)
	// data := lista.Fin(nuevaLista)

	lista.ImprimirLista(nuevaLista)

	loca := lista.Localiza(40, nuevaLista)

	fmt.Println(loca)

	nuevaLista = lista.Anular(nuevaLista)
	fmt.Println(nuevaLista)

	// lista.ImprimirLista(nuevaLista)

}
