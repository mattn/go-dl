include $(GOROOT)/src/Make.inc

TARG     = dl
CGOFILES = dl.go
CGO_LDFLAGS = -ldl

include $(GOROOT)/src/Make.pkg
