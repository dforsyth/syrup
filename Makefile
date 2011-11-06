include $(GOROOT)/src/Make.inc
TARG= 	syrup
GOFILES= 	result.go setup.go rpc/zmq_rpc.go vertex.go loader.go

include $(GOROOT)/src/Make.pkg