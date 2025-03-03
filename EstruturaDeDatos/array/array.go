package array

import (
	"fmt"
	"slices"
)

func NewArray() {

	fmt.Println("array")
}

var datas []interface{}

// recorrido de un aray, s

func Reading() {
	// limite inferior es 0
	// limite superior  total del array
	index := 0
	for {
		if len(datas) == index {
			break
		}
		fmt.Printf("posicion %d -> %v \n", index, datas[index])
		index++
	}
}

// aladir elementos al array se debe ralizar varios pasos
// tanto si es para agregarlo al princio o al final
func Add(newData int) {

	datas = append(datas, newData)

}

// si es string ordenar de forma alphabetic
func Clean(cleanData int) {

	for index, data := range datas {
		if data == cleanData {
			inicio := datas[:index]
			fin := datas[index+1:]
			datas = slices.Concat(inicio, fin)
		}
	}

}

// booble sort
func BobbleSort() {

}
