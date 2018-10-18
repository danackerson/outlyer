package structures

// Metrics for system reporting
type Metrics struct {
	Sys struct {
		CPU struct {
			AvgPercent float64 `json:"avgcpu"`
			Load1      float64 `json:"load1"`
		}
		Mem struct {
			BytesActive    uint64 `json:"active"`
			BytesAvailable uint64 `json:"available"`
		} `json:"mem"`
		Disk struct {
			BytesFree uint64 `json:"free"`
			BytesUsed uint64 `json:"used"`
		} `json:"disk"`
		Net struct {
			BytesReceived    uint64 `json:"rx"`
			BytesTransmitted uint64 `json:"tx"`
		} `json:"net"`
	} `json:"sys"`
	// Add your custom item here! e.g.
	// Nginx struct {...} `json:"nginx"`
}

// MetricsRegistry to store the observed metrics
type MetricsRegistry struct {
	UnixTimeNano int64   `json:"unixtimenano"`
	Measurement  Metrics `json:"measurement"`
}
