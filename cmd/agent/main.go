package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/lev-stas/metricsmonitor.git/internal/handlers"
	"time"
)

func main() {
	client := resty.New()
	metrics := make(map[string]float64)
	PollCount := &handlers.PollCountMetric{}
	RandomValue := &handlers.RandomValueMetric{}
	for {
		time.Sleep(time.Second * 2)
		metrics = handlers.PickMetrics(RuntimeMetrics)
		PollCount.IncrementPollCount()
		time.Sleep(time.Second * 2)
		metrics = handlers.PickMetrics(RuntimeMetrics)
		PollCount.IncrementPollCount()
		time.Sleep(time.Second * 2)
		metrics = handlers.PickMetrics(RuntimeMetrics)
		PollCount.IncrementPollCount()
		time.Sleep(time.Second * 2)
		metrics = handlers.PickMetrics(RuntimeMetrics)
		PollCount.IncrementPollCount()
		time.Sleep(time.Second * 2)
		metrics = handlers.PickMetrics(RuntimeMetrics)
		PollCount.IncrementPollCount()
		for metricName, metricValue := range metrics {
			url := fmt.Sprintf("%s/update/gauge/%s/%f", URL, metricName, metricValue)
			_, err := client.R().
				SetHeader("Content-Type", "text/plain").
				Post(url)
			//_, err := http.Post(url, "text/plain", nil)
			if err != nil {
				fmt.Printf("ERROR: %s was not sent\n", metricName)
			}
			fmt.Printf("%s metric was sent successfully\n", metricName)
		}
		RandomValue.GenerateRandomValue()
		randomValueUrl := fmt.Sprintf("%s/update/gauge/RandomValue/%f", URL, RandomValue.RandomValue)
		counterUrl := fmt.Sprintf("%s/update/counter/PollCount/%d", URL, PollCount.PollCount)
		_, err := client.R().SetHeader("Content-Type", "text/plain").Post(randomValueUrl)
		//_, err := http.Post(randomValueUrl, "text/plain", nil)
		if err != nil {
			fmt.Printf("ERROR: RandomValue  was not sent\n")
		}
		fmt.Printf("RandomValue metric was sent successfully\n")
		_, er := client.R().SetHeader("Content-Type", "text/plain").Post(counterUrl)
		//_, er := http.Post(counterUrl, "text/plain", nil)
		if er != nil {
			fmt.Printf("ERROR: PollCount was not sent\n")
		}
		fmt.Printf("PollCount metric was sent successfully\n")
		PollCount.ResetPollCount()
	}
}
