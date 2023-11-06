package main

import "flag"

type configParams struct {
	server         string
	reportInterval int
	pollInterval   int
}

var params = configParams{}

func parseFlags() {
	flag.StringVar(&params.server, "a", "localhost:8080", "server address and port")
	flag.IntVar(&params.reportInterval, "r", 10, "Report interval")
	flag.IntVar(&params.pollInterval, "p", 2, "Report interval")

	flag.Parse()
}
