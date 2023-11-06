package main

import (
	"flag"
)

type configParams struct {
	host string
}

var params = configParams{}

func parseFlags() {

	flag.StringVar(&params.host, "a", ":8080", "server address and port number")

	flag.Parse()

}
