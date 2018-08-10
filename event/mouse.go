package event

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
	//"yaqiang/recoil"
	"yaqiang/global"
	"yaqiang/recoil"
)

// String returns a human-friendly display name of the hotkey
// such as "Hotkey[Id: 1, Alt+Ctrl+O]"
var (
	user32 = windows.NewLazySystemDLL("user32.dll")
	procSetWindowsHookEx = user32.NewProc("SetWindowsHookExA")
	procCallNextHookEx = user32.NewProc("CallNextHookEx")
	procUnhookWindowsHookEx = user32.NewProc("UnhookWindowsHookEx")
	procGetMessage = user32.NewProc("GetMessageW")
	keyboardHook HHOOK
)
const (
	WH_MOUSE_LL    = 14
	WM_LBUTTONDOWN = 513
	WM_LBUTTONUP   = 514
	WM_RBUTTONDOWN = 516
	WM_RBUTTONUP   = 517
)

type (
	DWORD uint32
	WPARAM uintptr
	LPARAM uintptr
	LRESULT uintptr
	HANDLE uintptr
	HINSTANCE HANDLE
	HHOOK HANDLE
	HWND HANDLE
)

type HOOKPROC func(int, WPARAM, LPARAM) LRESULT


// http://msdn.microsoft.com/en-us/library/windows/desktop/dd162805.aspx
type POINT struct {
	X, Y int32
}

// http://msdn.microsoft.com/en-us/library/windows/desktop/ms644958.aspx
type MSG struct {
	Hwnd HWND
	Message uint32
	WParam uintptr
	LParam uintptr
	Time uint32
	Pt POINT
}


func SetWindowsHookEx(idHook int, lpfn HOOKPROC, hMod HINSTANCE, dwThreadId DWORD) HHOOK {
	ret, _, _ := procSetWindowsHookEx.Call(
		uintptr(idHook),
		uintptr(syscall.NewCallback(lpfn)),
		uintptr(hMod),
		uintptr(dwThreadId),
	)
	return HHOOK(ret)
}

func CallNextHookEx(hhk HHOOK, nCode int, wParam WPARAM, lParam LPARAM) LRESULT {
	ret, _, _ := procCallNextHookEx.Call(
		uintptr(hhk),
		uintptr(nCode),
		uintptr(wParam),
		uintptr(lParam),
	)
	return LRESULT(ret)
}

func UnhookWindowsHookEx(hhk HHOOK) bool {
	ret, _, _ := procUnhookWindowsHookEx.Call(
		uintptr(hhk),
	)
	return ret != 0
}

func GetMessage(msg *MSG, hwnd HWND, msgFilterMin uint32, msgFilterMax uint32) int {
	ret, _, _ := procGetMessage.Call(
		uintptr(unsafe.Pointer(msg)),
		uintptr(hwnd),
		uintptr(msgFilterMin),
		uintptr(msgFilterMax))
	return int(ret)
}

func Start() {
	// defer user32.Release()
	keyboardHook = SetWindowsHookEx(WH_MOUSE_LL,
		(HOOKPROC)(func(nCode int, wparam WPARAM, lparam LPARAM) LRESULT {
			if nCode == 0 && wparam == WM_LBUTTONDOWN {
				//print("按下鼠标左键")
				global.MouseLeftDown = true
				go recoil.Start()
			}
			if nCode == 0 && wparam == WM_LBUTTONUP {
				//print("松开鼠标左键")
				global.MouseLeftDown = false
			}
			return CallNextHookEx(keyboardHook, nCode, wparam, lparam)
		}), 0, 0)
	var msg MSG
	for GetMessage(&msg, 0, 0, 0) != 0 {

	}
	UnhookWindowsHookEx(keyboardHook)
	keyboardHook = 0
}