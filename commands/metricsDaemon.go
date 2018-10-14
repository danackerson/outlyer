package commands

import (
	"log"
	"math/rand"
	"time"

	"github.com/danackerson/outlyer/structures"
)

var registryStore = []structures.MetricsRegistry{}

// StoreMetrics to store in the registry
func StoreMetrics() {
	systemSnap := new(structures.BaseMetrics)
	systemSnap.Sys.CPU = rand.Float64()

	registryStore = append(registryStore,
		structures.MetricsRegistry{Clock: time.Now(), Measurement: *systemSnap})

	log.Printf("%v", registryStore[len(registryStore)-1])
}

// GetLastMetricSample returns last taken store metrics object
func GetLastMetricSample() structures.BaseMetrics {
	return registryStore[len(registryStore)-1].Measurement
}
