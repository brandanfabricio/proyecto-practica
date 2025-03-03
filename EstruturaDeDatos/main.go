package main

import (
	"EstruturaDeDatos/array"
	"fmt"
)

func main() {
	array.NewArray()

	// data := []interface{}{}
	array.Add(100)
	array.Add(23)
	array.Add(455)
	array.Add(577)
	array.Add(12)
	array.Add(123)
	array.Add(44)
	array.Add(56)
	array.Add(9)
	array.Reading()
	fmt.Println("############")
	array.Clean(12)
	array.Reading()
	fmt.Println("############")
	array.Clean(44)
	array.Add(43)
	array.Reading()

}
