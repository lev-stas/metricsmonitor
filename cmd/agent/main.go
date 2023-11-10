package main

import (
	"github.com/lev-stas/metricsmonitor.git/internal/agnetrunners"
	"github.com/lev-stas/metricsmonitor.git/internal/configs"
)

func main() {
	configs.GetAgentConfigs()
	metrics := make(map[string]float64)
	PollCount := &agnetrunners.PollCountMetric{}

	agnetrunners.MetricsSender(&metrics, PollCount)

}
