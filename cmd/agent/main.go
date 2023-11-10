package main

import (
	"github.com/lev-stas/metricsmonitor.git/internal/configs"
	"github.com/lev-stas/metricsmonitor.git/internal/handlers"
)

func main() {
	configs.GetAgentConfigs()
	metrics := make(map[string]float64)
	PollCount := &handlers.PollCountMetric{}
	RandomValue := &handlers.RandomValueMetric{}

	handlers.MetricsSender(&metrics, PollCount, RandomValue)

}
