package lista

import "fmt"

type tipo_elemento int
type Tipo_Lista struct {
	Elemento tipo_elemento
	Sig      *Tipo_Lista
}

func Fin(L *Tipo_Lista) *Tipo_Lista {
	// devuele el ultimo elemento de la lista
	posicion := L

	for posicion.Sig != nil {
		posicion = posicion.Sig
	}
	return posicion

}

func Insertar(x tipo_elemento, L *Tipo_Lista) {

	temp := L.Sig
	L.Sig = &Tipo_Lista{}
	L.Sig.Elemento = x
	L.Sig.Sig = temp

}

func ImprimirLista(cabeza *Tipo_Lista) {
	actual := cabeza.Sig
	for actual != nil {
		fmt.Printf("%d -> ", actual.Elemento)
		actual = actual.Sig
	}
	fmt.Println("nil")
}

func Suprime(L *Tipo_Lista) {
	L.Sig = L.Sig.Sig
}

func Localiza(x tipo_elemento, L *Tipo_Lista) *Tipo_Lista {
	lista := L
	for lista.Sig != nil {
		if lista.Sig.Elemento == x {
			return lista.Sig
		} else {
			lista = lista.Sig
		}

	}

	return lista
}
func Anular(L *Tipo_Lista) *Tipo_Lista {
	newList := &Tipo_Lista{}
	newList.Sig = nil

	return newList
}
