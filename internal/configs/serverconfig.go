package configs

import (
	"flag"
	"github.com/caarlos0/env/v6"
	"log"
)

type ServerConfigParams struct {
	Host     string
	LogLevel string
}

type ServerEnvParams struct {
	Address  string `env:"ADDRESS"`
	LogLevel string `env:"LOG_LEVEL"`
}

var ServerParams ServerConfigParams
var ServerEnvs ServerEnvParams

func GetServerConfigs() {
	flag.StringVar(&ServerParams.Host, "a", ":8080", "Server address and port number")
	flag.StringVar(&ServerParams.LogLevel, "l", "info", "log level")
	flag.Parse()

	err := env.Parse(&ServerEnvs)
	if err != nil {
		log.Fatalln(err)
	}
	if address := ServerEnvs.Address; address != "" {
		ServerParams.Host = address
	}
	if logLevel := ServerEnvs.LogLevel; logLevel != "" {
		ServerParams.LogLevel = logLevel
	}
}
