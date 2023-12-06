package main

import (
	"github.com/lev-stas/metricsmonitor.git/internal/configs"
	"github.com/lev-stas/metricsmonitor.git/internal/gzipper"
	"github.com/lev-stas/metricsmonitor.git/internal/logger"
	"github.com/lev-stas/metricsmonitor.git/internal/metricsstorage"
	"github.com/lev-stas/metricsmonitor.git/internal/preloadworkers"
	"github.com/lev-stas/metricsmonitor.git/internal/routers"
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
	writer, err := metricsstorage.NewFileWriter(&configs.ServerParams)
	if err != nil {
		logger.Log.Errorw("Error during creating New File Writer")
	}
	if configs.ServerParams.InitLoad() {
		if err := preloadworkers.RestoreMetricsFromFile(storage); err != nil {
			logger.Log.Fatal("Error during reading metrics from file")
		}
	}

	if !configs.ServerParams.SyncSave() {
		preloadworkers.WriteMetricsTicker(storage, writer)
	}

	log.Fatalln(http.ListenAndServe(configs.ServerParams.Host, gzipper.GzipMiddleware(logger.RequestResponseLogger(r))))
}
