package dl

/*
#include <windows.h.h>

unsigned long dlcall(unsigned long p, unsigned long a1, unsigned long a2, unsigned long a3) {
  typedef unsigned long (*type_call)(unsigned long, unsigned long, unsigned long);
  type_call pcall = (type_call)p;
  return pcall(a1,a2,a3);
}
*/
import "C"
import "os"
import "unsafe"

/*
const (
  RTLD_LAZY = C.RTLD_LAZY
  RTLD_NOW = C.RTLD_NOW
  RTLD_GLOBAL = C.RTLD_GLOBAL
)
*/

func Open(filename string) (uintptr, os.Error) {
  ret, err := syscall.LoadLibrary(filename)
  if err == 0 {
    return uintptr(ret), nil
  }
  return uintptr(ret), os.NewError(syscall.Errstr(err))
}

func Sym(handle uintptr, symbol string) (uintptr, os.Error) {
  ret, err := syscall.GetProcAddress(handle, symbol)
  if err == 0 {
    return uintptr(ret), nil
  }
  return uintptr(ret), os.NewError(syscall.Errstr(err))
}

func Close(handle uintptr) os.Error {
  syscall.FreeLibrary(handle)
  return nil
}

func Call(p, a1, a2, a3 uintptr) uintptr {
  ret := C.dlcall(
    C.ulong(p),
    C.ulong(a1),
    C.ulong(a2),
    C.ulong(a3))
  return uintptr(ret)
}
