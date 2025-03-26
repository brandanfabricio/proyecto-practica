package lista

import (
	"fmt"
)

const log_max = 100

type Lista struct {
	elementos [log_max]int
	utl       int
}

func NewListArray() *Lista {

	return &Lista{}
}

//	func (list *Lista) Fin(p int) int {
//		if p > log_max {
//			fmt.Println("Fuera de rango")
//			return -1
//		}
//		return list.elementos[p-1]
//	}
func (list *Lista) Fin() int {
	return list.utl + 1
}
func (list *Lista) Len() int {
	return list.utl
}

func (l *Lista) Inserta(x int, p int) {

	if l.utl >= log_max {
		fmt.Println("Lista esta llena")
		return
	}
	if p > l.utl || p < 0 {
		fmt.Println("La posicion no existe")
		return
	}

	for q := l.utl; q >= p; q-- {
		l.elementos[q+1] = l.elementos[q]
	}
	l.utl = l.utl + 1
	l.elementos[p] = x
}

func (l *Lista) Suprime(p int) {
	if p > l.utl || p < 0 {
		fmt.Println("La posicion no existe")
		return
	}
	l.utl = l.utl - 1
	for q := p; q <= l.utl; q++ {
		l.elementos[q] = l.elementos[q+1]
	}
}

func (l *Lista) Localizar(x int) int {
	for q := 0; q <= l.utl; q++ {
		if x == l.elementos[q] {
			return q
		}
	}
	return -1
}
func (l *Lista) Recupera(p int) int {

	if p >= l.Fin() {
		fmt.Println("Dato no existe ")
		return -1
	}
	return l.elementos[p]
}
func (l *Lista) Siguiente(p int) int {
	if p >= l.Len()-1 {
		return -1
	}
	return p + 1
}
func (l *Lista) Anterior(p int) int {
	if p <= 0 {
		return -1
	}
	return p - 1
}

func (l *Lista) Anular() {
	l.utl = 0
	l.elementos = [log_max]int{}
}

func (l *Lista) Imprimir() {
	for index, elemento := range l.elementos {
		fmt.Printf("%d -> %d \n", index, elemento)
		if l.Fin()-1 == index {
			break
		}
	}
}
