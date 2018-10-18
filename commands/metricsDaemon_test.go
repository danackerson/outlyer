package commands

import (
	"testing"

	"github.com/danackerson/outlyer/structures"
)

func Test_getAvgCPU(t *testing.T) {
	t.Parallel()

	cpuPercentages := []float64{0, 0.5, 0.5, 1.5}
	systemSnap := new(structures.Metrics)

	getAvgCPU(cpuPercentages, systemSnap)
	if systemSnap.Sys.CPU.AvgPercent != 2.5/4/100 {
		t.Errorf("Incorrect CPU load: exp %f act %f",
			systemSnap.Sys.CPU, 2.5/4)
	}
}
