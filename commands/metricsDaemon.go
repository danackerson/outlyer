package commands

import (
	"log"
	"sync"
	"time"

	"github.com/shirou/gopsutil/load"

	"github.com/danackerson/outlyer/structures"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

var registryStore = []structures.MetricsRegistry{}
var registryMutex = &sync.Mutex{}

// StoreMetricMeasurement populates the registry with requested data points
func StoreMetricMeasurement() {
	systemSnap := new(structures.Metrics)

	cpuPercentages, err := cpu.Percent(0, false)
	if err != nil || len(cpuPercentages) == 0 {
		log.Printf("CPU measurement failed: %s\n", err.Error())
	}
	getAvgCPU(cpuPercentages, systemSnap)

	avgStat, err := load.Avg()
	if err != nil {
		log.Printf("failed to measure CPU.Load1: %s", err.Error())
	}
	getCPULoad(*avgStat, systemSnap)

	virtualMemory, err := mem.VirtualMemory()
	if err != nil || virtualMemory.Active == 0 {
		log.Printf("VM measurement failed: %s\n", err.Error())
	}
	getMemory(virtualMemory, systemSnap)

	disk, err := disk.Usage("/")
	if err != nil || disk.Total == 0 {
		log.Printf("Disk measurement failed: %s\n", err.Error())
	}
	getDisk(disk, systemSnap)

	net, err := net.IOCounters(false)
	if err != nil || net[0].BytesSent == 0 {
		log.Printf("Network measurement failed: %s\n", err.Error())
	}
	getNetwork(net[0], systemSnap)
	// Populate your custom metric here e.g.
	// getNginxStats(systemSnap)

	nextMeasurement := structures.MetricsRegistry{
		UnixTimeNano: time.Now().UnixNano(), Measurement: *systemSnap}

	// protect against dirty writes/reads
	registryMutex.Lock()
	registryStore = append(registryStore, nextMeasurement)
	registryMutex.Unlock()

	//log.Printf("Reg: %v\n", registryStore[len(registryStore)-1])
}

func getCPULoad(cpuAvgStats load.AvgStat, measurement *structures.Metrics) {
	measurement.Sys.CPU.Load1 = cpuAvgStats.Load1
}

func getAvgCPU(cpuPercentages []float64, measurement *structures.Metrics) {
	totalCPU := 0.0
	for _, cpuPercent := range cpuPercentages {
		totalCPU += cpuPercent / 100.0
	}

	measurement.Sys.CPU.AvgPercent = totalCPU / float64(len(cpuPercentages))
}

func getMemory(virtualMemory *mem.VirtualMemoryStat,
	measurement *structures.Metrics) {
	measurement.Sys.Mem.BytesActive = virtualMemory.Used
	measurement.Sys.Mem.BytesAvailable = virtualMemory.Available
}

func getDisk(disk *disk.UsageStat, measurement *structures.Metrics) {
	measurement.Sys.Disk.BytesFree = disk.Free
	measurement.Sys.Disk.BytesUsed = disk.Used
}

func getNetwork(net net.IOCountersStat, measurement *structures.Metrics) {
	measurement.Sys.Net.BytesReceived = net.BytesRecv
	measurement.Sys.Net.BytesTransmitted = net.BytesSent
}

// GetAllMetrics returns the json list of timestamps & measurements
func GetAllMetrics() []structures.MetricsRegistry {
	return registryStore
}
