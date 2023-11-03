package main

import (
	"fmt"
	"github.com/lev-stas/metricsmonitor.git/internal/handlers"
	"net/http"
	"strconv"
	"time"
)

func main() {
	metrics := make(map[string]float64)
	PollCount := &handlers.PollCountMetric{}
	RandomValue := &handlers.RandomValueMetric{}
	for {
		time.Sleep(time.Second * 2)
		metrics = handlers.PickMetrics(runtimeMetrics)
		PollCount.IncrementPollCount()
		time.Sleep(time.Second * 2)
		metrics = handlers.PickMetrics(runtimeMetrics)
		PollCount.IncrementPollCount()
		time.Sleep(time.Second * 2)
		metrics = handlers.PickMetrics(runtimeMetrics)
		PollCount.IncrementPollCount()
		time.Sleep(time.Second * 2)
		metrics = handlers.PickMetrics(runtimeMetrics)
		PollCount.IncrementPollCount()
		time.Sleep(time.Second * 2)
		metrics = handlers.PickMetrics(runtimeMetrics)
		PollCount.IncrementPollCount()
		for metricName, metricValue := range metrics {
			stringifiedValue := strconv.FormatFloat(metricValue, 'f', -1, 64)
			url := fmt.Sprintf("%s/update/gauge/%s/%s", URL, metricName, stringifiedValue)
			_, err := http.Post(url, "", nil)
			if err != nil {
			}
		}
		RandomValue.GenerateRandomValue()
		randomValueUrl := fmt.Sprintf("%s/update/gauge/RandomValue/%s", URL, RandomValue.RandomValue)
		counterUrl := fmt.Sprintf("%s/update/counter/pollcount/%s", URL, PollCount.PollCount)
		_, err := http.Post(randomValueUrl, "", nil)
		if err != nil {

		}
		_, er := http.Post(counterUrl, "", nil)
		if er != nil {

		}
		PollCount.ResetPollCount()
	}
}
