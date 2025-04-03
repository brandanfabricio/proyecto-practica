package lista

import (
	"fmt"
)

type TipoElemento int
type ListaPuntero struct {
	elementos TipoElemento
	sig       *ListaPuntero
}

func NewListPuntero() *ListaPuntero {
	return &ListaPuntero{}
}

func (l *ListaPuntero) Fin() *ListaPuntero {

	q := l
	for l.sig != nil {

		q = q.sig
	}
	return q
}

func (l *ListaPuntero) Insertar(x TipoElemento, posicion int) {

	actual := l

	for actual.sig != nil {
		actual = actual.sig
		fmt.Println(actual.elementos)
	}

	fmt.Println("##############")
	nuevoNodo := NewListPuntero()
	nuevoNodo.elementos = x

	// fmt.Println("%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%")
	// fmt.Println(nuevoNodo.elementos)
	// fmt.Println(nuevoNodo.sig)
	// fmt.Println("%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%")

	actual.sig = nuevoNodo

}

func (l *ListaPuntero) ImprimirList() {
	actual := l
	for l != nil {
		for actual != nil {
			fmt.Print(actual.elementos, " -> ")
			actual = actual.sig
		}
	}
	fmt.Println("nil")
}
