package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	kernel32          = syscall.NewLazyDLL("kernel32.dll")
	openProcess       = kernel32.NewProc("OpenProcess")
	readProcessMemory = kernel32.NewProc("ReadProcessMemory")
	closeHandle       = kernel32.NewProc("CloseHandle")
)

func main() {
	pid := 19704 // Cambia esto por el PID real del proceso objetivo
	address := uintptr(0xc000022070)

	// Abrir el proceso con permisos de lectura
	processHandle, _, _ := openProcess.Call(0x10, 0, uintptr(pid))
	if processHandle == 0 {
		fmt.Println("No se pudo abrir el proceso")
		return
	}
	defer closeHandle.Call(processHandle)

	// Leer memoria
	var value int
	var bytesRead uintptr
	success, _, _ := readProcessMemory.Call(
		processHandle,
		address,
		uintptr(unsafe.Pointer(&value)),
		unsafe.Sizeof(value),
		uintptr(unsafe.Pointer(&bytesRead)),
	)

	if success == 0 {
		fmt.Println("Error al leer la memoria")
		return
	}

	fmt.Println("Valor leído:", value)
}
