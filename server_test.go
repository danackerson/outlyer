package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danackerson/outlyer/commands"
	"github.com/danackerson/outlyer/structures"
)

/*"sys.cpu": "4.15",
"sys.mem.active": "8029741056",
"sys.mem.available": "6602049536",
"sys.disk.free": "13605457920",
"sys.disk.used": "76558884181",
"sys.net.rx": "217003.31",
"sys.net.tx": "11390.29"*/

func Test_baseMetrics(t *testing.T) {
	t.Parallel()

	// redirect case
	req, err := http.NewRequest("GET", "http://localhost"+getHTTPPort()+"/metrics", nil)
	if err != nil {
		t.Fatal(err)
	}

	commands.StoreMetrics()

	res := httptest.NewRecorder()
	getMetrics(res, req)

	var target structures.BaseMetrics
	json.NewDecoder(res.Body).Decode(&target)
	if target.Sys.CPU == 0 {
		log.Printf("CPU: %v", target.Sys.CPU)
		t.Fatalf("expected non-zero CPU load report; %f", target.Sys.CPU)
	}
	// TODO: test for valid JSON
	// TODO: test matching values in test
}
