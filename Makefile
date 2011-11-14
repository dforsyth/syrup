include $(GOROOT)/src/Make.inc
TARG= 	syrup
GOFILES= 	result/stdout_result.go setup/setup.go env/file.go vertex/vertex.go vertex/shell.go # rpc/zmq_rpc.go loader/loader.go

include $(GOROOT)/src/Make.pkg
