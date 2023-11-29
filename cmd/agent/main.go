package main

import (
	"github.com/lev-stas/metricsmonitor.git/internal/agentrunners"
	"github.com/lev-stas/metricsmonitor.git/internal/configs"
	"github.com/lev-stas/metricsmonitor.git/internal/logger"
	"log"
)

func main() {
	if err := logger.LogInit(configs.ServerParams.LogLevel); err != nil {
		log.Fatalln(err)
	}
	configs.GetAgentConfigs()
	metrics := make(map[string]float64)
	PollCount := &agentrunners.PollCountMetric{}

	agentrunners.MetricsSender(&metrics, PollCount)

}
