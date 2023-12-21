package handlers

import (
	"encoding/json"
	"github.com/lev-stas/metricsmonitor.git/internal/configs"
	"github.com/lev-stas/metricsmonitor.git/internal/datamodels"
	"github.com/lev-stas/metricsmonitor.git/internal/logger"
	"github.com/lev-stas/metricsmonitor.git/internal/metricsstorage"
	"go.uber.org/zap"
	"net/http"
)

var counterMetric string = "counter"
var gaugeMetric string = "gauge"

type UpdateStorageInterface interface {
	Set(metric string, value float64)
	Inc(metric string, value int64)
	GetCounterMetric(metric string) (int64, bool)
	GetAllCounterMetrics() map[string]int64
	GetAllGaugeMetrics() map[string]float64
}

func HandleUpdateJSON(storage *metricsstorage.MemStorage, fileWriter *metricsstorage.FileWriter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var metric datamodels.Metric
		if err := json.NewDecoder(r.Body).Decode(&metric); err != nil {
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
			storage.Inc(metric.ID, *metric.Delta)
			pollcounter, found := storage.GetCounterMetric(metric.ID)
			if !found {
				logger.Log.Error("Error during updating PollCount metric")
				http.Error(w, "Error during updating metric", http.StatusInternalServerError)
				return
			}
			*metric.Delta = pollcounter
		}

		if metric.MType == gaugeMetric {
			storage.Set(metric.ID, *metric.Value)
		}

		res, err := json.Marshal(metric)
		if err != nil {
			logger.Log.Error("Error during marshaling response", zap.Error(err))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(res)
		if err != nil {
			logger.Log.Errorw("Error during sending response", "error", err)
		}
		if configs.ServerParams.StorageInterval == 0 {
			err := metricsstorage.SaveMetricsToFile(fileWriter, storage)
			defer fileWriter.Close()
			if err != nil {
				//logger.Log.Errorw("Error during saving metrics to file")
			}
		}

	}

}
