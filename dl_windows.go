package dl

import "C"
import "errors"
import "syscall"

func Open(filename string) (uintptr, error) {
	ret, err := syscall.LoadLibrary(filename)
	if err == 0 {
		return uintptr(ret), nil
	}
	return uintptr(ret), errors.New(syscall.Errstr(err))
}

func Sym(handle uintptr, symbol string) (uintptr, error) {
	ret, err := syscall.GetProcAddress(syscall.Handle(handle), symbol)
	if err == 0 {
		return uintptr(ret), nil
	}
	return uintptr(ret), errors.New(syscall.Errstr(err))
}

func Close(handle uintptr) error {
	syscall.FreeLibrary(syscall.Handle(handle))
	return nil
}

func Call(p, a1, a2, a3 uintptr) uintptr {
	ret, _, _ := syscall.Syscall(p,
		uintptr(p),
		uintptr(a1),
		uintptr(a2),
		uintptr(a3))
	return uintptr(ret)
}
