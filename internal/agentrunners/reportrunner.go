package agentrunners

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/lev-stas/metricsmonitor.git/internal/datamodels"
	"github.com/lev-stas/metricsmonitor.git/internal/logger"
)

func ReportRunner(server string, metrics *map[string]float64, pollCount *PollCountMetric) {
	gaugeUrl := fmt.Sprintf("http://%s/update/", server)
	counterUrl := fmt.Sprintf("http://%s/update/", server)
	client := resty.New()

	for metricName, metricValue := range *metrics {
		metric := datamodels.Metric{
			ID:    metricName,
			MType: "gauge",
			Delta: nil,
			Value: &metricValue,
		}
		body, err := json.Marshal(metric)
		if err != nil {
			logger.Log.Errorw("Error during marshaling metric object", "error", err)
		}

		compressedBody, err := GzipCompress(body)

		if err != nil {
			logger.Log.Errorw("Error during compressing gauge request", "error", err)
		}
		_, er := client.R().
			SetHeaderMultiValues(map[string][]string{
				"Content-Type":     []string{"application/json"},
				"Accept-Encoding":  []string{"gzip"},
				"Content-Encoding": []string{"gzip"},
			}).
			SetBody(compressedBody).
			Post(gaugeUrl)
		if er != nil {
			logger.Log.Errorw("Error during sending metric to server", "error", er)
		}
		logger.Log.Debugw("Successfully sent metric", "metric", metric.ID)
	}
	counterMetric := datamodels.Metric{
		ID:    "PollCount",
		MType: "counter",
		Delta: &pollCount.PollCount,
		Value: nil,
	}
	body, err := json.Marshal(counterMetric)
	if err != nil {
		logger.Log.Errorw("Error during marshaling PollCountMetric", "error", err)
	}
	compressedBody, er := GzipCompress(body)
	if er != nil {
		logger.Log.Errorw("Error during compressing counter request", "error", er)
	}
	_, err = client.R().
		SetHeaderMultiValues(map[string][]string{
			"Content-Type":     []string{"application/json"},
			"Accept-Encoding":  []string{"gzip"},
			"Content-Encoding": []string{"gzip"},
		}).
		SetBody(compressedBody).
		Post(counterUrl)
	if err != nil {
		logger.Log.Errorw("Error during sending PollCount metric to server", "error", err)
	}
	logger.Log.Debug("Successfully sent PollCount Metric to server", "metric", counterMetric.ID)
	pollCount.ResetPollCount()

}
