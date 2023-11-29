package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/lev-stas/metricsmonitor.git/internal/datamodels"
	"github.com/lev-stas/metricsmonitor.git/internal/logger"
	"net/http"
	"strconv"
)

type StorageValueInterface interface {
	GetGaugeMetric(metric string) (float64, bool)
	GetCounterMetric(metric string) (int64, bool)
}

func ValueHandlerJSON(storage StorageValueInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var metric datamodels.Metric
		var res datamodels.Metric

		if err := json.NewDecoder(r.Body).Decode(&metric); err != nil {
			logger.Log.Errorw("Error during decoding request body", "error", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
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
					MType: "counter",
					Delta: &value,
				}
			}
		}

		body, err := json.Marshal(res)
		if err != nil {
			logger.Log.Errorw("Error during marshaling response", "error", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(body)
		if err != nil {
			http.Error(w, "Can't response", http.StatusInternalServerError)
			return
		}
	}
}

func ValueHandler(storage StorageValueInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		metricsType := chi.URLParam(r, "metricsType")
		metricsName := chi.URLParam(r, "metricsName")
		var response []byte

		if metricsType != "gauge" && metricsType != "counter" {
			http.Error(w, "Wrong metric type", http.StatusBadRequest)
			return
		}

		if metricsType == "gauge" {
			value, found := storage.GetGaugeMetric(metricsName)
			if !found {
				http.Error(w, "Metric not found", http.StatusNotFound)
				return
			} else {
				response = []byte(strconv.FormatFloat(value, 'f', -2, 64))
			}
		} else if metricsType == "counter" {
			value, found := storage.GetCounterMetric(metricsName)
			if !found {
				http.Error(w, "Metric not found", http.StatusNotFound)
				return
			} else {
				response = []byte(strconv.FormatInt(value, 10))
			}
		}

		w.WriteHeader(http.StatusOK)
		_, err := w.Write(response)
		if err != nil {
			http.Error(w, "Can't response", http.StatusInternalServerError)
			return
		}

	}
}
