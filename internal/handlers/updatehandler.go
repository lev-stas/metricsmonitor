package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/lev-stas/metricsmonitor.git/internal/datamodels"
	"github.com/lev-stas/metricsmonitor.git/internal/logger"
	"go.uber.org/zap"
	"net/http"
)

var counterMetric string = "counter"
var gaugeMetric string = "gauge"

type UpdateStorageInterface interface {
	SetGaugeMetric(metric string, value float64)
	SetCounterMetric(metric string, value int64)
}

func HandleUpdate(storage UpdateStorageInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not Allows", http.StatusMethodNotAllowed)
			return
		}

		var metric datamodels.Metric
		if err := json.NewDecoder(r.Body).Decode(&metric); err != nil {
			logger.Log.Error("Error during decoding metric object", zap.Error(err))
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		if metric.ID == "" || (metric.Value == nil && metric.Delta == nil) {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		if metric.MType != counterMetric && metric.MType != gaugeMetric {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		if metric.MType == counterMetric {
			storage.SetCounterMetric(metric.ID, *metric.Delta)
		}

		if metric.MType == gaugeMetric {
			storage.SetGaugeMetric(metric.ID, *metric.Value)
		}

		logger.Log.Debug("Received metric: ", zap.Any("metric", metric))
		fmt.Println("Successfully sent metric")

		res, err := json.Marshal(metric)
		if err != nil {
			logger.Log.Error("Error during marshaling response", zap.Error(err))
			return
		}

		w.WriteHeader(http.StatusOK)
		_, err = w.Write(res)
		if err != nil {
			logger.Log.Error("Error during sending response", zap.Error(err))
		}
	}

}
