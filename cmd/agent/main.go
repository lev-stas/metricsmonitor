package main

import (
	"github.com/lev-stas/metricsmonitor.git/internal/configs"
	"github.com/lev-stas/metricsmonitor.git/internal/handlers"
	"time"
)

func main() {
	configs.GetAgentConfigs()
	metrics := make(map[string]float64)
	PollCount := &handlers.PollCountMetric{}
	RandomValue := &handlers.RandomValueMetric{}

	pollInterval := time.Second * time.Duration(configs.AgentParams.PollInterval)
	reportInterval := time.Second * time.Duration(configs.AgentParams.ReportInterval)
	pollTicker := time.NewTicker(pollInterval)
	reportTicker := time.NewTicker(reportInterval)

	defer pollTicker.Stop()
	defer reportTicker.Stop()

	for {
		select {
		case <-pollTicker.C:
			handlers.PollMetricsRunner(&metrics, PollCount)

		case <-reportTicker.C:
			handlers.ReportRunner(configs.AgentParams.Server, &metrics, PollCount, RandomValue)
		}
	}
}
