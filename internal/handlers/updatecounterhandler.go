package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/lev-stas/metricsmonitor.git/internal/configs"
	"github.com/lev-stas/metricsmonitor.git/internal/logger"
	"github.com/lev-stas/metricsmonitor.git/internal/metricsstorage"
	"net/http"
	"strconv"
)

func HandleCounterUpdate(storage *metricsstorage.MemStorage, fileWriter metricsstorage.FileWriterInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not Allows", http.StatusMethodNotAllowed)
			return
		}

		metricsName := chi.URLParam(r, "metricsName")
		metricsValueRaw := chi.URLParam(r, "metricsValue")

		if metricsName == "" || metricsValueRaw == "" {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		metricsValue, err := strconv.ParseInt(metricsValueRaw, 10, 64)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		storage.Inc(metricsName, metricsValue)

		w.WriteHeader(http.StatusOK)

		if configs.ServerParams.StorageInterval == 0 {
			err := metricsstorage.SaveMetricsToFile(fileWriter, storage)
			defer fileWriter.Close()
			if err != nil {
				logger.Log.Errorw("Error during saving metrics to file")
			}
		}

	}
}
