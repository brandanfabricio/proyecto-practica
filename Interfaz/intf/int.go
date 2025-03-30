package intf

import (
	"syscall"
	"unsafe"
)

var (
	user32   = syscall.NewLazyDLL("user32.dll")
	kernel32 = syscall.NewLazyDLL("kernel32.dll")

	registerClassEx  = user32.NewProc("RegisterClassExW")
	createWindowEx   = user32.NewProc("CreateWindowExW")
	defWindowProc    = user32.NewProc("DefWindowProcW")
	showWindow       = user32.NewProc("ShowWindow")
	getMessage       = user32.NewProc("GetMessageW")
	translateMessage = user32.NewProc("TranslateMessage")
	dispatchMessage  = user32.NewProc("DispatchMessageW")
	postQuitMessage  = user32.NewProc("PostQuitMessage")

	loadCursor      = user32.NewProc("LoadCursorW")
	makeIntResource = uintptr(32512) // ID del cursor por defecto
	getModuleHandle = kernel32.NewProc("GetModuleHandleW")
)

const (
	WS_OVERLAPPEDWINDOW = 0x00CF0000
	WS_VISIBLE          = 0x10000000
	CW_USEDEFAULT       = ^0x80000000
	WM_DESTROY          = 0x0002
	SW_SHOW             = 5
)

// Definimos la estructura WNDCLASSEX
type WndClassEx struct {
	Size       uint32
	Style      uint32
	WndProc    uintptr
	ClsExtra   int32
	WndExtra   int32
	Instance   uintptr
	Icon       uintptr
	Cursor     uintptr
	Background uintptr
	MenuName   *uint16
	ClassName  *uint16
	IconSm     uintptr
}

// Función para manejar eventos de la ventana
func wndProc(hwnd uintptr, msg uint32, wparam, lparam uintptr) uintptr {
	switch msg {
	case WM_DESTROY:
		postQuitMessage.Call(0)
		return 0
	}
	ret, _, _ := defWindowProc.Call(hwnd, uintptr(msg), wparam, lparam)
	return ret
}

func test() {
	className, _ := syscall.UTF16PtrFromString("MiVentana")

	// Obtener el handle del módulo
	hInstance, _, _ := getModuleHandle.Call(0)

	// Crear la estructura de la ventana
	wndClass := WndClassEx{
		Size:      uint32(unsafe.Sizeof(WndClassEx{})),
		WndProc:   syscall.NewCallback(wndProc),
		Instance:  hInstance,
		ClassName: className,
	}

	// Cargar el cursor
	wndClass.Cursor, _, _ = loadCursor.Call(0, makeIntResource)

	// Registrar la clase de la ventana
	registerClassEx.Call(uintptr(unsafe.Pointer(&wndClass)))

	// Crear la ventana
	hwnd, _, _ := createWindowEx.Call(
		0,
		uintptr(unsafe.Pointer(className)),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("Mi App en Go"))),
		WS_OVERLAPPEDWINDOW|WS_VISIBLE,
		uintptr(CW_USEDEFAULT&0xFFFFFFFF), uintptr(CW_USEDEFAULT&0xFFFFFFFF), 800, 600,
		0, 0, hInstance, 0,
	)

	// Mostrar la ventana
	showWindow.Call(hwnd, SW_SHOW)

	// Bucle de mensajes
	var msg struct {
		HWnd    uintptr
		Message uint32
		WParam  uintptr
		LParam  uintptr
		Time    uint32
		Pt      struct {
			X, Y int32
		}
	}

	for {
		ret, _, _ := getMessage.Call(uintptr(unsafe.Pointer(&msg)), 0, 0, 0)
		if ret == 0 {
			break
		}
		translateMessage.Call(uintptr(unsafe.Pointer(&msg)))
		dispatchMessage.Call(uintptr(unsafe.Pointer(&msg)))
	}
}
