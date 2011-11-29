package syrup

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"path"
	"waffle"
)

var Host string
var Port string
var MinWorkers uint64
var MinPartitionsPerWorker uint64
var MessageThreshold int64
var VertexThreshold int64
var LoadPath string
var PersistPath string

func FlagsParse() {
	flag.StringVar(&Port, "port", "50000", "node port")
	flag.StringVar(&Host, "host", "127.0.0.1", "node address")
	flag.Uint64Var(&MinWorkers, "minWorkers", 2, "min workers")
	flag.Uint64Var(&MinPartitionsPerWorker, "ppw", 1, "min partitions per worker")
	flag.Int64Var(&MessageThreshold, "mthresh", 1000, "message threshold")
	flag.Int64Var(&VertexThreshold, "vthresh", 1000, "vertex threshold")
	flag.StringVar(&LoadPath, "loadPath", "data", "data load path")
	flag.StringVar(&PersistPath, "persistPath", "persist", "data persist path")

	flag.Parse()
}

func configPath(pathToConfig string) (string, error) {
	if path.IsAbs(pathToConfig) {
		return pathToConfig, nil
	}
	wd, err := os.Getwd()
	if err != nil {
		return "", nil
	}
	return path.Join(wd, pathToConfig), nil
}

func configLoader(config interface{}, pathToConfig string) (err error) {
	var path string
	if path, err = configPath(pathToConfig); err != nil {
		return
	}

	var configBytes []byte
	if configBytes, err = ioutil.ReadFile(path); err != nil {
		return
	}

	if err = json.Unmarshal(configBytes, config); err != nil {
		return
	}
	return
}

func LoadMasterConfig(m *waffle.Master, pathToConfig string) error {
	return configLoader(&m.Config, pathToConfig)
}

func LoadWorkerConfig(w *waffle.Worker, pathToConfig string) error {
	return configLoader(&w.Config, pathToConfig)
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
