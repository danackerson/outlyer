package structures

import (
	"encoding/json"
	"strings"
	"testing"
)

func Test_metricStructure(t *testing.T) {
	t.Parallel()

	const jsonMeasurement = `{"sys":{"cpu":0.02985074626882049,
		"mem":{"active":1045004288,"available":10264137728},
		"disk":{"free":414975938560,"used":2934190080},
		"net":{"tx":996775,"rx":3762064}}
	}`

	var target Metrics
	json.NewDecoder(strings.NewReader(jsonMeasurement)).Decode(&target)

	expectedCPU := 0.02985074626882049
	if target.Sys.CPU != expectedCPU {
		t.Errorf("CPU act %f != exp %f", target.Sys.CPU, expectedCPU)
	}
	expectedMemActive := uint64(1045004288)
	if target.Sys.Mem.BytesActive != expectedMemActive {
		t.Errorf("MemActive act %d != exp %d", target.Sys.Mem.BytesActive,
			expectedMemActive)
	}
	expectedMemAvailable := uint64(10264137728)
	if target.Sys.Mem.BytesAvailable != expectedMemAvailable {
		t.Errorf("MemAvail act %d != exp %d", target.Sys.Mem.BytesActive,
			expectedMemAvailable)
	}
	expectedDiskFree := uint64(414975938560)
	if target.Sys.Disk.BytesFree != expectedDiskFree {
		t.Errorf("DiskFree act %d != exp %d", target.Sys.Disk.BytesFree,
			expectedDiskFree)
	}
	expectedDiskUsed := uint64(2934190080)
	if target.Sys.Disk.BytesUsed != expectedDiskUsed {
		t.Errorf("DiskUsed act %d != exp %d", target.Sys.Disk.BytesUsed,
			expectedDiskUsed)
	}
	expectedNetSend := uint64(996775)
	if target.Sys.Net.BytesTransmitted != expectedNetSend {
		t.Errorf("NetSend act %d != exp %d", target.Sys.Net.BytesTransmitted,
			expectedNetSend)
	}
	expectedNetRcvd := uint64(3762064)
	if target.Sys.Net.BytesReceived != expectedNetRcvd {
		t.Errorf("NetRcvd act %d != exp %d", target.Sys.Net.BytesReceived,
			expectedNetRcvd)
	}
}
