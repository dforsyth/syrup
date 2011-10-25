package syrup

import (
	"flag"
	"waffle"
)

var Host string
var Port string
var MinWorkers uint64
var PartitionsPerWorker uint64
var MessageThreshold int64
var VertexThreshold int64

func FlagsParse() {
	flag.StringVar(&Port, "port", "50000", "node port")
	flag.StringVar(&Host, "host", "127.0.0.1", "node address")
	flag.Uint64Var(&MinWorkers, "minWorkers", 2, "min workers")
	flag.Uint64Var(&PartitionsPerWorker, "ppw", 1, "partitions per worker")
	flag.Int64Var(&MessageThreshold, "mthresh", 50, "message threshold")
	flag.Int64Var(&VertexThreshold, "vthresh", 50, "vertex threshold")

	flag.Parse()
}

func CreateWorker() *waffle.Worker {
	w := waffle.NewWorker(Host, Port)
	w.Config.MessageThreshold = MessageThreshold
	w.Config.VertexThreshold = VertexThreshold
	return w
}

func CreateMaster() *waffle.Master {
	m := waffle.NewMaster(Host, Port)
	m.Config.MinWorkers = MinWorkers
	m.Config.PartitionsPerWorker = PartitionsPerWorker
	return m
}
