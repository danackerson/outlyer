package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/danackerson/outlyer/commands"
	"github.com/danackerson/outlyer/structures"
)

func Test_manyMetrics(t *testing.T) {
	t.Parallel()

	req, err := http.NewRequest("GET", "http://localhost"+
		getHTTPPort()+"/metrics", nil)
	if err != nil {
		t.Fatal(err)
	}

	go startMetricsDaemon()
	sleeping := time.Duration(10)
	time.Sleep(sleeping * time.Second) // let the system collect some metrics

	res := httptest.NewRecorder()
	getAllMetrics(res, req)

	var allMetrics []structures.MetricsRegistry
	json.NewDecoder(res.Body).Decode(&allMetrics)

	if len(allMetrics) < 5 {
		t.Errorf("Only %d metrics stored after %d secs",
			len(allMetrics), sleeping)
	}
}

func Test_oneMetric(t *testing.T) {
	t.Parallel()

	req, err := http.NewRequest("GET", "http://localhost"+
		getHTTPPort()+"/metrics", nil)
	if err != nil {
		t.Fatal(err)
	}

	commands.StoreMetricMeasurement()

	res := httptest.NewRecorder()
	getAllMetrics(res, req)

	var allMetrics []structures.MetricsRegistry
	json.NewDecoder(res.Body).Decode(&allMetrics)

	nowNano := time.Now().UnixNano()

	firstMetric := allMetrics[0]
	if firstMetric.UnixTimeNano >= nowNano {
		t.Errorf("expected recorded time %d before now (%d)",
			firstMetric.UnixTimeNano, nowNano)
	}
	if firstMetric.Measurement.Sys.Mem.BytesActive == 0 {
		t.Errorf("expected non-zero Memory report")
	}
	if firstMetric.Measurement.Sys.Disk.BytesUsed == 0 {
		t.Errorf("expected non-zero Disk report")
	}
	if firstMetric.Measurement.Sys.Net.BytesTransmitted == 0 {
		t.Errorf("expected non-zero Net report")
	}
}
