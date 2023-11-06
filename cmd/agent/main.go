package main

import (
	"github.com/lev-stas/metricsmonitor.git/internal/handlers"
	"time"
)

func main() {
	parseFlags()
	metrics := make(map[string]float64)
	PollCount := &handlers.PollCountMetric{}
	RandomValue := &handlers.RandomValueMetric{}

	pollInterval := time.Second * time.Duration(params.PollInterval)
	reportInterval := time.Second * time.Duration(params.ReportInterval)
	pollTicker := time.NewTicker(pollInterval)
	reportTicker := time.NewTicker(reportInterval)

	defer pollTicker.Stop()
	defer reportTicker.Stop()

	for {
		select {
		case <-pollTicker.C:
			handlers.PollMetricsRunner(&metrics, PollCount)

		case <-reportTicker.C:
			handlers.ReportRunner(params.Server, &metrics, PollCount, RandomValue)
		}
	}
}
