package handlers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/lev-stas/metricsmonitor.git/internal/memstorage"
	"net/http"
	"strconv"
)

var counterMetric string = "counter"
var gaugeMetric string = "gauge"

func HandleUpdate(storage *memstorage.MemStorage) http.HandlerFunc {
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
