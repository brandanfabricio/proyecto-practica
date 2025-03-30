package main

import (
	"syscall"
	"unsafe"
)

var (
	user32           = syscall.NewLazyDLL("user32.dll")
	gdi32            = syscall.NewLazyDLL("gdi32.dll")
	registerClassEx  = user32.NewProc("RegisterClassExW")
	createWindowEx   = user32.NewProc("CreateWindowExW")
	defWindowProc    = user32.NewProc("DefWindowProcW")
	dispatchMessage  = user32.NewProc("DispatchMessageW")
	getMessage       = user32.NewProc("GetMessageW")
	translateMessage = user32.NewProc("TranslateMessage")
	postQuitMessage  = user32.NewProc("PostQuitMessage")
	loadCursor       = user32.NewProc("LoadCursorW")
	showWindow       = user32.NewProc("ShowWindow")
	updateWindow     = user32.NewProc("UpdateWindow")
	beginPaint       = user32.NewProc("BeginPaint")
	endPaint         = user32.NewProc("EndPaint")
	textOut          = gdi32.NewProc("TextOutW")
)

const (
	WS_OVERLAPPEDWINDOW = 0x00CF0000
	WS_VISIBLE          = 0x10000000
	CW_USEDEFAULT       = 0x80000000
	WM_DESTROY          = 0x0002
	WM_PAINT            = 0x000F
	SW_SHOW             = 5
)

type WNDCLASSEX struct {
	Size       uint32
	Style      uint32
	WndProc    uintptr
	ClassExtra int32
	WndExtra   int32
	Instance   uintptr
	Icon       uintptr
	Cursor     uintptr
	Background uintptr
	MenuName   *uint16
	ClassName  *uint16
	IconSm     uintptr
}

type PAINTSTRUCT struct {
	Hdc       uintptr
	FErase    int32
	RcPaint   [4]int32
	Restore   int32
	IncUpdate int32
	Reserved  [32]byte
}

type MSG struct {
	Hwnd    uintptr
	Message uint32
	WParam  uintptr
	LParam  uintptr
	Time    uint32
	Pt      struct {
		X, Y int32
	}
}

func main() {
	cr, _, _ := loadCursor.Call(0, 32512)
	// 1️⃣ Definir la clase de ventana
	className := syscall.StringToUTF16Ptr("MiVentana")
	wndClass := WNDCLASSEX{
		Size:      uint32(unsafe.Sizeof(WNDCLASSEX{})),
		WndProc:   syscall.NewCallback(windowProc), // Manejador de eventos de la ventana
		ClassName: className,
		Cursor:    cr, // Cursor por defecto
	}

	// 2️⃣ Registrar la clase de la ventana
	registerClassEx.Call(uintptr(unsafe.Pointer(&wndClass)))

	// 3️⃣ Crear la ventana
	hwnd, _, _ := createWindowEx.Call(
		0,
		uintptr(unsafe.Pointer(className)),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("Ventana en Go"))),
		WS_OVERLAPPEDWINDOW|WS_VISIBLE,
		CW_USEDEFAULT, CW_USEDEFAULT, 500, 400,
		0, 0, 0, 0,
	)

	// 4️⃣ Mostrar la ventana
	showWindow.Call(hwnd, SW_SHOW)
	updateWindow.Call(hwnd)

	// 5️⃣ Bucle de mensajes (para que la ventana siga funcionando)
	var msg MSG
	for {
		ret, _, _ := getMessage.Call(uintptr(unsafe.Pointer(&msg)), 0, 0, 0)
		if ret == 0 {
			break
		}
		translateMessage.Call(uintptr(unsafe.Pointer(&msg)))
		dispatchMessage.Call(uintptr(unsafe.Pointer(&msg)))
	}
}

// 🔹 Manejador de eventos de la ventana
func windowProc(hwnd uintptr, msg uint32, wparam, lparam uintptr) uintptr {
	switch msg {
	case WM_PAINT: // 📌 Se activa cuando la ventana necesita redibujarse
		var ps PAINTSTRUCT
		hdc, _, _ := beginPaint.Call(hwnd, uintptr(unsafe.Pointer(&ps)))

		// Dibujar el texto "Hola Mundo" en la ventana
		text := syscall.StringToUTF16Ptr("¡Hola Mundo desde Go!")
		textOut.Call(hdc, 50, 50, uintptr(unsafe.Pointer(text)), uintptr(len("¡Hola Mundo desde Go!")))

		endPaint.Call(hwnd, uintptr(unsafe.Pointer(&ps)))
		return 0

	case WM_DESTROY: // 📌 Se activa cuando la ventana se cierra
		postQuitMessage.Call(0)
		return 0
	}

	ret, _, _ := defWindowProc.Call(hwnd, uintptr(msg), wparam, lparam)
	return ret
}
