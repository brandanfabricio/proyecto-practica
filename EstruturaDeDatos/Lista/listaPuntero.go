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
	}
	nuevoNodo := NewListPuntero()
	nuevoNodo.elementos = x
	actual.sig = nuevoNodo
}

func (l *ListaPuntero) ImprimirList() {
	actual := l
	for l != nil {
		for actual != nil {
			fmt.Print(actual.elementos, " -> ")
			actual = actual.sig

			if actual == nil {
				return
			}
		}
	}
}

func (l *ListaPuntero) SUPRIME(posicion int) {

	actual := l
	count := 0

	for actual.sig != nil {

		fmt.Println("-> ", actual.elementos)

		fmt.Print(count)
		if count == posicion {
			fmt.Println("aki")

		}
		actual = actual.sig
		count++

	}

}
