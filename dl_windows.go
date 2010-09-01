package dl

import "os"
import "syscall"

func Open(filename string) (uintptr, os.Error) {
  ret, err := syscall.LoadLibrary(filename)
  if err == 0 {
    return uintptr(ret), nil
  }
  return uintptr(ret), os.NewError(syscall.Errstr(err))
}

func Sym(handle uintptr, symbol string) (uintptr, os.Error) {
  ret, err := syscall.GetProcAddress(uint32(handle), symbol)
  if err == 0 {
    return uintptr(ret), nil
  }
  return uintptr(ret), os.NewError(syscall.Errstr(err))
}

func Close(handle uintptr) os.Error {
  syscall.FreeLibrary(uint32(handle))
  return nil
}

func Call(p, a1, a2, a3 uintptr) uintptr {
  ret, _, _ := syscall.Syscall(p, a1, a2, a3)
  return ret;
}
