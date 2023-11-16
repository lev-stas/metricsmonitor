package handlers

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type StorageValueInterface interface {
	GetGaugeMetric(metric string) (float64, bool)
	GetCounterMetric(metric string) (int64, bool)
}

func ValueHandler(storage StorageValueInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		metricsType := chi.URLParam(r, "metricsType")
		metricsName := chi.URLParam(r, "metricsName")
		var response []byte

		if r.Method != http.MethodGet {
			http.Error(w, "Wrong request method", http.StatusMethodNotAllowed)
			return
		}

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

		_, err := w.Write(response)
		if err != nil {
			http.Error(w, "Can't response", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
