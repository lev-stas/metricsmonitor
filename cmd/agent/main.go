package main

import (
	"github.com/lev-stas/metricsmonitor.git/internal/agentrunners"
	"github.com/lev-stas/metricsmonitor.git/internal/configs"
)

func main() {
	configs.GetAgentConfigs()
	metrics := make(map[string]float64)
	PollCount := &agentrunners.PollCountMetric{}

	agentrunners.MetricsSender(&metrics, PollCount)

}
