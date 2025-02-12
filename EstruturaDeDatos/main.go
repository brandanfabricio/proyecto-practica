package main

import (
	"EstruturaDeDatos/array"
)

func main() {
	array.NewArray()

	data := []interface{}{100, 23, 455, 577, 12, 123, 44, 56, 9}

	array.Reading(data)
}
