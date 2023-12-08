package agentrunners

import (
	"github.com/go-resty/resty/v2"
	"github.com/lev-stas/metricsmonitor.git/internal/configs"
	"time"
)

func MetricsSender(metrics *map[string]float64, pollCountMetric *PollCountMetric) {

	metricsList := RuntimeMetrics
	pollInterval := time.Second * time.Duration(configs.AgentParams.PollInterval)
	reportInterval := time.Second * time.Duration(configs.AgentParams.ReportInterval)
	pollTicker := time.NewTicker(pollInterval)
	reportTicker := time.NewTicker(reportInterval)
	client := resty.New().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept-Encoding", "gzip").
		SetHeader("Content-Encoding", "gzip")
	//client.R().SetHeaderMultiValues(map[string][]string{
	//	"Content-Type":     []string{"application/json"},
	//	"Accept-Encoding":  []string{"gzip"},
	//	"Content-Encoding": []string{"gzip"},
	//})

	defer pollTicker.Stop()
	defer reportTicker.Stop()

	for {
		select {
		case <-pollTicker.C:
			PollRunner(metricsList, metrics, pollCountMetric)

		case <-reportTicker.C:
			ReportRunner(configs.AgentParams.Server, metrics, pollCountMetric, client)
		}
	}
}
