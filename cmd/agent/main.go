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

	pollInterval := time.Second * time.Duration(params.pollInterval)
	reportInterval := time.Second * time.Duration(params.reportInterval)
	pollTicker := time.NewTicker(pollInterval)
	reportTicker := time.NewTicker(reportInterval)

	defer pollTicker.Stop()
	defer reportTicker.Stop()

	for {
		select {
		case <-pollTicker.C:
			handlers.PollMetricsRunner(&metrics, PollCount)

		case <-reportTicker.C:
			handlers.ReportRunner(params.server, &metrics, PollCount, RandomValue)
		}
	}
}
