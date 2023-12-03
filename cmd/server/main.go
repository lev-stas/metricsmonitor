package main

import (
	"github.com/lev-stas/metricsmonitor.git/internal/configs"
	"github.com/lev-stas/metricsmonitor.git/internal/gzipper"
	"github.com/lev-stas/metricsmonitor.git/internal/logger"
	"github.com/lev-stas/metricsmonitor.git/internal/metricsstorage"
	"github.com/lev-stas/metricsmonitor.git/internal/routers"
	"github.com/lev-stas/metricsmonitor.git/internal/timetickers"
	"log"
	"net/http"
)

var storage *metricsstorage.MemStorage

func main() {
	if err := logger.LogInit(configs.ServerParams.LogLevel); err != nil {
		log.Fatalln(err)
	}
	configs.GetServerConfigs()
	storage = metricsstorage.NewMemStorage()
	r := routers.RootRouter(storage)
	if configs.ServerParams.Restore {
		fileReader, err := metricsstorage.NewMetricsReader(configs.ServerParams.StorageFile)
		if err != nil {
			log.Fatalf("Error during creating Filereader: %v", err)
		}
		defer fileReader.Close()

		for {
			metric, er := fileReader.ReadMetric()
			if er != nil {
				break

			}
			if metric != nil {
				switch metric.MType {
				case "gauge":
					storage.Set(metric.ID, *metric.Value)
				case "counter":
					storage.Inc(metric.ID, *metric.Delta)
				}
			} else {
				break
			}
		}
	}
	if configs.ServerParams.StorageInterval > 0 {
		timetickers.WriteMetricsTicker(storage)
	}
	log.Fatalln(http.ListenAndServe(configs.ServerParams.Host, gzipper.GzipMiddleware(logger.RequestResponseLogger(r))))
}
