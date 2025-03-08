package lista

import "fmt"

const log_max = 100

type Lista struct {
	elementos [log_max]int
	utl       int
}

func NewListArray() *Lista {
	return &Lista{}
}

func (list *Lista) Fin() int {
	fmt.Println("Posicion")
	return list.utl + 1
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

	for q := l.utl; q > p; q-- {
		fmt.Println(p)
		fmt.Println(q)
		fmt.Println(q + 1)
		fmt.Println(x)
		fmt.Println(l.elementos)
		fmt.Println(l.elementos[q])
		l.elementos[q+1] = l.elementos[q]

	}

	// for index := range list.elementos {
	// 	list.elementos[index+1] = list.elementos[index]
	// }
	l.utl = l.utl + 1
	l.elementos[p] = x

}
