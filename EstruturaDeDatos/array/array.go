package array

import "fmt"

func NewArray() {

	fmt.Println("array")
}

// recorrido de un aray, s

func Reading(array []interface{}) {
	// limite inferior es 0
	// limite superior  total del array
	index := 0
	for {
		if len(array) == index {
			break
		}
		fmt.Printf("posicion %d -> %v \n", index, array[index])
		index++
	}
}

// aladir elementos al array se debe ralizar varios pasos
// tanto si es para agregarlo al princio o al final
//
