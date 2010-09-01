include $(GOROOT)/src/Make.inc

TARG     = dl
ifeq ($(GOOS),windows)
GOFILES = dl_windows.go
else
CGOFILES = dl_unix.go
CGO_LDFLAGS = -ldl
endif

include $(GOROOT)/src/Make.pkg
