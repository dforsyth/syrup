package syrup

import (
	"testing"
	"waffle"
)

const masterConfigPath = "testdata/master.cfg"
const workerConfigPath = "testdata/worker.cfg"

func TestMasterConfigLoad(t *testing.T) {
	m := new(waffle.Master)
	if err := LoadMasterConfig(m, masterConfigPath); err != nil {
		t.Errorf("Error: %v", err)
	}
	if m.Config.MinWorkers != 2 || m.Config.RegisterWait != 10 || m.Config.MinPartitionsPerWorker != 2 || m.Config.HeartbeatInterval != 10 || m.Config.MaxSteps != 30 || m.Config.JobId != "testJob" || m.Config.StartStep != 5 {
		t.Error("Error: config read")
	}
}

func TestWorkerConfigLoad(t *testing.T) {
	w := new(waffle.Worker)
	if err := LoadWorkerConfig(w, workerConfigPath); err != nil {
		t.Errorf("Error: %v", err)
	}
	if w.Config.MessageThreshold != 1000 || w.Config.VertexThreshold != 1000 || w.Config.MasterHost != "testhost" || w.Config.MasterPort != "50000" || w.Config.RegisterRetry != 10 {
		t.Error("Error: config read")
	}
}