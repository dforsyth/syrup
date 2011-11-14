package syrup

import (
	"flag"
	"waffle"
)

var Host string
var Port string
var MinWorkers uint64
var MinPartitionsPerWorker uint64
var MessageThreshold int64
var VertexThreshold int64

func FlagsParse() {
	flag.StringVar(&Port, "port", "50000", "node port")
	flag.StringVar(&Host, "host", "127.0.0.1", "node address")
	flag.Uint64Var(&MinWorkers, "minWorkers", 2, "min workers")
	flag.Uint64Var(&MinPartitionsPerWorker, "ppw", 1, "min partitions per worker")
	flag.Int64Var(&MessageThreshold, "mthresh", 1000, "message threshold")
	flag.Int64Var(&VertexThreshold, "vthresh", 1000, "vertex threshold")

	flag.Parse()
}

func CreateWorker(s waffle.WorkerRpcServer, c waffle.WorkerRpcClient) *waffle.Worker {
	w := waffle.NewWorker(Host, Port)
	w.Config.MessageThreshold = MessageThreshold
	w.Config.VertexThreshold = VertexThreshold
	w.SetRpcServer(s)
	w.SetRpcClient(c)
	return w
}

func CreateMaster(s waffle.MasterRpcServer, c waffle.MasterRpcClient) *waffle.Master {
	m := waffle.NewMaster(Host, Port)
	m.Config.MinWorkers = MinWorkers
	m.Config.MinPartitionsPerWorker = MinPartitionsPerWorker
	m.SetRpcServer(s)
	m.SetRpcClient(c)
	return m
}
