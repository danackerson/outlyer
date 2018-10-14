package structures

import "time"

// BaseMetrics for minimum reporting
type BaseMetrics struct {
	Sys struct {
		CPU float64 `json:"cpu"`
		Mem struct {
			Active    int64 `json:"active"`
			Available int64 `json:"available"`
		} `json:"mem"`
		Disk struct {
			BytesFree int64 `json:"free"`
			BytesUsed int64 `json:"used"`
		} `json:"disk"`
		Net struct {
			BytesReceived    float64 `json:"rx"`
			BytesTransmitted float64 `json:"tx"`
		} `json:"net"`
	} `json:"sys"`
}

// MetricsRegistry to store the observed metrics
type MetricsRegistry struct {
	Clock       time.Time
	Measurement BaseMetrics
}
