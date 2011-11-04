package dl

/*
#include <dlfcn.h>
#include <stdlib.h>

static inline void** make_list(int count) {
	return (void**)malloc(sizeof(void*) * count);
}

static inline void destroy_list(void** list) {
	free(list);
}

static inline void* get_item(void** list, int n) {
	return list[n];
}

static inline void set_item(void** list, int n, void* item) {
	list[n] = item;
}

static unsigned long dlcall(unsigned long p, unsigned long a1, unsigned long a2, unsigned long a3) {
  typedef unsigned long (*type_call)(unsigned long, unsigned long, unsigned long);
  type_call pcall = (type_call)p;
  return pcall(a1,a2,a3);
}
*/
import "C"
import "errors"
import "unsafe"

/*
const (
  RTLD_LAZY = C.RTLD_LAZY
  RTLD_NOW = C.RTLD_NOW
  RTLD_GLOBAL = C.RTLD_GLOBAL
)
*/

func Open(filename string /*, flag int*/ ) (uintptr, error) {
	ptr := C.CString(filename)
	defer C.free(unsafe.Pointer(ptr))
	ret := C.dlopen(ptr, /*C.int(flag)*/ C.RTLD_LAZY)
	if ret != nil {
		return uintptr(ret), nil
	}
	return uintptr(ret), errors.New(C.GoString(C.dlerror()))
}

func Sym(handle uintptr, symbol string) (uintptr, error) {
	ptr := C.CString(symbol)
	defer C.free(unsafe.Pointer(ptr))
	ret := C.dlsym(unsafe.Pointer(handle), ptr)
	if ret != nil {
		return uintptr(ret), nil
	}
	return uintptr(ret), errors.New(C.GoString(C.dlerror()))
}

func Close(handle uintptr) error {
	ret := C.dlclose(unsafe.Pointer(handle))
	if ret != 0 {
		return errors.New(C.GoString(C.dlerror()))
	}
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
