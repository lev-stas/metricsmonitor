package agnetrunners

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

func ReportRunner(server string, metrics *map[string]float64, pollCount *PollCountMetric) {
	client := resty.New()

	for metricName, metricValue := range *metrics {
		url := fmt.Sprintf("http://%s/update/gauge/%s/%f", server, metricName, metricValue)
		_, err := client.R().SetHeader("Content-Type", "text/plain").Post(url)
		if err != nil {
			fmt.Printf("ERROR: %s was not sent\n", metricName)
		}
		fmt.Printf("%s metric was sent successfully\n", metricName)
	}

	counterUrl := fmt.Sprintf("http://%s/update/counter/PollCount/%d", server, pollCount.PollCount)
	_, er := client.R().SetHeader("Content-Type", "text/plain").Post(counterUrl)
	if er != nil {
		fmt.Printf("ERROR:  PollCount metric was not sen\n")
	}
	fmt.Printf("PollCount metric was sent successfully\n")
	pollCount.ResetPollCount()

}