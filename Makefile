include $(GOROOT)/src/Make.inc

TARG     = github.com/mattn/go-dl/dl
ifeq ($(GOOS),windows)
CGOFILES = dl_windows.go
EXAMPLE= example_windows
else
CGOFILES = dl_unix.go
CGO_LDFLAGS = -ldl
EXAMPLE= example_unix
endif

include $(GOROOT)/src/Make.pkg

example: install
	$(GC) $(EXAMPLE).go
	$(LD) -o $(EXAMPLE) $(EXAMPLE).8

