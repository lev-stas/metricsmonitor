package handlers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/lev-stas/metricsmonitor.git/internal/memstorage"
	"net/http"
	"strconv"
)

func HandleUpdate(storage *memstorage.MemStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not Allows", http.StatusMethodNotAllowed)
			return
		}

		//requestParts := strings.Split(r.URL.Path, "/")
		//

		metricsType := chi.URLParam(r, "metricsType")
		metricsName := chi.URLParam(r, "metricsName")
		metricsValueRaw := chi.URLParam(r, "metricsValue")
		//fmt.Printf("metricsType: %s\n", metricsType)
		//fmt.Printf("metricsName: %s\n", metricsName)
		//fmt.Printf("metricsValue: %s\n", metricsValueRaw)

		if metricsName == "" || metricsValueRaw == "" {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		//if metricsValueRaw == "" {
		//	http.Error(w, "Bad Request", http.StatusBadRequest)
		//}

		if metricsType != "counter" && metricsType != "gauge" {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		if metricsType == "counter" {
			metricsValue, err := strconv.ParseInt(metricsValueRaw, 10, 64)
			if err != nil {
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}
			storage.SetCounterMetric(metricsName, metricsValue)
		}

		if metricsType == "gauge" {
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
