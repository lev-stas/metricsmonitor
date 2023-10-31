package main

import (
	"fmt"
	"github.com/lev-stas/metricsmonitor.git/internal/handlers"
)

func main() {

	var metrics = handlers.PickMetrics(runtimeMetrics)

	for metricName, metricValue := range metrics {
		fmt.Println(metricName, metricValue)
	}
}
