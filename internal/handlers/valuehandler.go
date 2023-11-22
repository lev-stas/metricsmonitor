package handlers

import (
	"encoding/json"
	"github.com/lev-stas/metricsmonitor.git/internal/datamodels"
	"github.com/lev-stas/metricsmonitor.git/internal/logger"
	"go.uber.org/zap"
	"net/http"
)

type StorageValueInterface interface {
	GetGaugeMetric(metric string) (float64, bool)
	GetCounterMetric(metric string) (int64, bool)
}

func ValueHandler(storage StorageValueInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var metric datamodels.Metric
		var res datamodels.Metric

		if err := json.NewDecoder(r.Body).Decode(&metric); err != nil {
			logger.Log.Error("Error during decoding request body", zap.Error(err))
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		if r.Method != http.MethodPost {
			http.Error(w, "Wrong request method", http.StatusMethodNotAllowed)
			return
		}

		if metric.MType != "gauge" && metric.MType != "counter" {
			http.Error(w, "Wrong metric type", http.StatusBadRequest)
			return
		}

		if metric.MType == "gauge" {
			value, found := storage.GetGaugeMetric(metric.ID)
			if !found {
				http.Error(w, "Metric not found", http.StatusNotFound)
				return
			} else {
				res = datamodels.Metric{
					ID:    metric.ID,
					MType: "gauge",
					Value: &value,
				}
			}
		} else if metric.MType == "counter" {
			value, found := storage.GetCounterMetric(metric.ID)
			if !found {
				http.Error(w, "Metric not found", http.StatusNotFound)
				return
			} else {
				res = datamodels.Metric{
					ID:    metric.ID,
					MType: "gauge",
					Delta: &value,
				}
			}
		}

		body, err := json.Marshal(res)
		if err != nil {
			logger.Log.Error("Error during marshaling response", zap.Error(err))
			return
		}

		w.WriteHeader(http.StatusOK)

		_, err = w.Write(body)
		if err != nil {
			http.Error(w, "Can't response", http.StatusInternalServerError)
			return
		}

	}
}
