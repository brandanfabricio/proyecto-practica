package intf

import (
	"fmt"
	"syscall"
)

var (
	kernel322    = syscall.NewLazyDLL("kernel32.dll")
	getTickCount = kernel322.NewProc("GetTickCount")
)

func timeSystem() {
	ret, _, _ := getTickCount.Call()

	fmt.Printf("Tiempo desde que se inicio el sistema %d ms\n ", ret)
}
