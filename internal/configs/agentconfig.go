package configs

import (
	"flag"
	"github.com/caarlos0/env/v6"
	"log"
)

type AgentConfigParams struct {
	Server         string
	ReportInterval int
	PollInterval   int
}

type AgentEnvParams struct {
	Address        string `env:"ADDRESS"`
	PollInterval   int    `env:"POLL_INTERVAL"`
	ReportInterval int    `env:"REPORT_INTERVAL"`
}

var AgentParams AgentConfigParams
var AgentEnvs AgentEnvParams

func GetAgentConfigs() {
	flag.StringVar(&AgentParams.Server, "a", "localhost:8080", "server address and port")
	flag.IntVar(&AgentParams.ReportInterval, "r", 10, "Report Interval")
	flag.IntVar(&AgentParams.PollInterval, "p", 2, "Report interval")

	flag.Parse()

	err := env.Parse(&AgentEnvs)
	if err != nil {
		log.Fatal(err)
	}
	if address := AgentEnvs.Address; address != "" {
		AgentParams.Server = address
	}
	if pollInterval := AgentEnvs.PollInterval; pollInterval != 0 {
		AgentParams.PollInterval = pollInterval
	}
	if reportInterval := AgentEnvs.ReportInterval; reportInterval != 0 {
		AgentParams.ReportInterval = reportInterval
	}
}
