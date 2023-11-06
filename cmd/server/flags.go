package main

import (
	"flag"
	"fmt"
	"github.com/caarlos0/env/v6"
	"log"
)

type ConfigParams struct {
	Host string
}

type EnvParams struct {
	Address string `env:"ADDRESS"`
}

var params ConfigParams
var envs EnvParams

func parseFlags() {

	flag.StringVar(&params.Host, "a", ":8080", "server address and port number")

	flag.Parse()

	err := env.Parse(&envs)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("ADDRESS ENV: %s\n", envs.Address)

	if address := envs.Address; address != "" {
		params.Host = address
	}

}
