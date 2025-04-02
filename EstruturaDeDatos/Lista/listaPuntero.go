package lista

import "fmt"

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

func (l *ListaPuntero) Insertar(x TipoElemento, p int) {
	fmt.Println(x)
	fmt.Println(p)

}
