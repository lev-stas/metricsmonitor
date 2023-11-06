package main

import (
	"flag"
	"github.com/caarlos0/env/v6"
	"log"
)

type ConfigParams struct {
	Server         string
	ReportInterval int
	PollInterval   int
}

type EnvParams struct {
	Address        string `env:"ADDRESS"`
	PollInterval   int    `env:"POLL_INTERVAL"`
	ReportInterval int    `env:"REPORT_INTERVAL"`
}

var params = ConfigParams{}
var envs EnvParams

func parseFlags() {
	flag.StringVar(&params.Server, "a", "localhost:8080", "server address and port")
	flag.IntVar(&params.ReportInterval, "r", 10, "Report interval")
	flag.IntVar(&params.PollInterval, "p", 2, "Report interval")

	flag.Parse()

	err := env.Parse(&envs)
	if err != nil {
		log.Fatal(err)
	}
	if address := envs.Address; address != "" {
		params.Server = address
	}
	if pollInterval := envs.PollInterval; pollInterval != 0 {
		params.PollInterval = pollInterval
	}
	if reportInterval := envs.ReportInterval; reportInterval != 0 {
		params.ReportInterval = reportInterval
	}

}
