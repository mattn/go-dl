package dl

import "C"
import "syscall"

func Open(filename string) (syscall.Handle, error) {
	return syscall.LoadLibrary(filename)
}

func Sym(handle syscall.Handle, symbol string) (uintptr, error) {
	return syscall.GetProcAddress(handle, symbol)
}

func Close(handle syscall.Handle) error {
	return syscall.FreeLibrary(handle)
}

func Call(p, a1, a2, a3 uintptr) uintptr {
	ret, _, _ := syscall.Syscall(p,
		uintptr(p),
		uintptr(a1),
		uintptr(a2),
		uintptr(a3))
	return uintptr(ret)
}
