package handlers

import (
	"github.com/lev-stas/metricsmonitor.git/internal/configs"
	"time"
)

func MetricsSender(metrics *map[string]float64, pollCountMetric *PollCountMetric, randomValueMetric *RandomValueMetric) {

	metricsList := RuntimeMetrics
	pollInterval := time.Second * time.Duration(configs.AgentParams.PollInterval)
	reportInterval := time.Second * time.Duration(configs.AgentParams.ReportInterval)
	pollTicker := time.NewTicker(pollInterval)
	reportTicker := time.NewTicker(reportInterval)

	defer pollTicker.Stop()
	defer reportTicker.Stop()

	for {
		select {
		case <-pollTicker.C:
			PollRunner(metricsList, metrics, pollCountMetric)

		case <-reportTicker.C:
			ReportRunner(configs.AgentParams.Server, metrics, pollCountMetric, randomValueMetric)
		}
	}
}
