package agentrunners

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/lev-stas/metricsmonitor.git/internal/datamodels"
	"github.com/lev-stas/metricsmonitor.git/internal/logger"
	"go.uber.org/zap"
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
			logger.Log.Error("Error during marshaling metric object", zap.Error(err))
		}

		compressedBody, err := GzipCompress(body)

		//fmt.Printf("Request body after compressing: %s", string(compressedBody))
		if err != nil {
			logger.Log.Error("Error during compressing gauge request")
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
			logger.Log.Error("Error during sending metric to server", zap.Error(err))
		}
		logger.Log.Debug("Successfully sent metric", zap.ByteString("metric", body))
	}
	counterMetric := datamodels.Metric{
		ID:    "PollCount",
		MType: "counter",
		Delta: &pollCount.PollCount,
		Value: nil,
	}
	body, err := json.Marshal(counterMetric)
	if err != nil {
		logger.Log.Error("Error during marshaling PollCountMetric", zap.Error(err))
	}
	compressedBody, er := GzipCompress(body)
	if er != nil {
		logger.Log.Error("Error during compressing counter request", zap.Error(er))
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
		logger.Log.Error("Error during sending PollCount metric to server", zap.Error(err))
	}
	logger.Log.Debug("Successfully sent PollCount Metric to server", zap.ByteString("metric", body))
	pollCount.ResetPollCount()

}
