include $(GOROOT)/src/Make.inc
TARG= 	syrup
GOFILES= 	result/stdout_result.go setup/setup.go # rpc/zmq_rpc.go vertex/vertex.go loader/loader.go

include $(GOROOT)/src/Make.pkg
