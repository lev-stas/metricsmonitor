package configs

import (
	"flag"
	"github.com/caarlos0/env/v6"
	"log"
)

type ServerConfigParams struct {
	Host string
}

type ServerEnvParams struct {
	Address string `env:"ADDRESS"`
}

var ServerParams ServerConfigParams
var ServerEnvs ServerEnvParams

func GetServerConfigs() {
	flag.StringVar(&ServerParams.Host, "a", ":8080", "Server address and port number")
	flag.Parse()

	err := env.Parse(&ServerEnvs)
	if err != nil {
		log.Fatalln(err)
	}
	if address := ServerEnvs.Address; address != "" {
		ServerParams.Host = address
	}
}
