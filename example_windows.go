package main

import "dl"
import "unsafe"

func clen(n []byte) int {
  for i := 0; i < len(n); i++ {
    if n[i] == 0 {
      return i
    }
  }
  return len(n)
}

func main() {
  handle, _ := dl.Open("msvcrt.dll")
  proc, _ := dl.Sym(handle, "getenv")
  ret := dl.Call(
    proc,
    uintptr(unsafe.Pointer(&([]byte)("USERPROFILE")[0])),
    0,
    0)
  bytes := ((*[255]byte)(unsafe.Pointer(ret)))
  println(string(bytes[0:clen(bytes)]))
}
