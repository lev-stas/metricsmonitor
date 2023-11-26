package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/lev-stas/metricsmonitor.git/internal/datamodels"
	"github.com/lev-stas/metricsmonitor.git/internal/logger"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

var counterMetric string = "counter"
var gaugeMetric string = "gauge"

type UpdateStorageInterface interface {
	SetGaugeMetric(metric string, value float64)
	SetCounterMetric(metric string, value int64)
	GetCounterMetric(metric string) (int64, bool)
}

func HandleUpdateJSON(storage UpdateStorageInterface) http.HandlerFunc {
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
			pollcounter, found := storage.GetCounterMetric(metric.ID)
			if !found {
				logger.Log.Error("Error during updating PollCount metric")
				http.Error(w, "Error during updating metric", http.StatusInternalServerError)
				return
			}
			*metric.Delta = pollcounter
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

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(res)
		if err != nil {
			logger.Log.Error("Error during sending response", zap.Error(err))
		}

	}

}

func HandleUpdate(storage UpdateStorageInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not Allows", http.StatusMethodNotAllowed)
			return
		}

		metricsType := chi.URLParam(r, "metricsType")
		metricsName := chi.URLParam(r, "metricsName")
		metricsValueRaw := chi.URLParam(r, "metricsValue")

		if metricsName == "" || metricsValueRaw == "" {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		if metricsType != counterMetric && metricsType != gaugeMetric {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		if metricsType == counterMetric {
			metricsValue, err := strconv.ParseInt(metricsValueRaw, 10, 64)
			if err != nil {
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}
			storage.SetCounterMetric(metricsName, metricsValue)
		}

		if metricsType == gaugeMetric {
			metricsValue, err := strconv.ParseFloat(metricsValueRaw, 62)
			if err != nil {
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}
			storage.SetGaugeMetric(metricsName, metricsValue)
		}

		fmt.Printf("Received metric update - Type: %s, Name: %s, Value: %s\n", metricsType, metricsName, metricsValueRaw)

		w.WriteHeader(http.StatusOK)
	}

}
