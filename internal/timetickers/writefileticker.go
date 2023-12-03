package timetickers

import (
	"github.com/lev-stas/metricsmonitor.git/internal/configs"
	"github.com/lev-stas/metricsmonitor.git/internal/logger"
	"github.com/lev-stas/metricsmonitor.git/internal/metricsstorage"
	"time"
)

func WriteMetricsTicker(storage *metricsstorage.MemStorage) {
	writeInterval := time.Second * time.Duration(configs.ServerParams.StorageInterval)
	writeTicker := time.NewTicker(writeInterval)
	fileWriter, err := metricsstorage.NewFileWriter(configs.ServerParams.StorageFile)
	if err != nil {
		logger.Log.Errorw("Error")
		return
	}

	go func() {
		for {
			select {
			case <-writeTicker.C:
				err := metricsstorage.SaveMetricsToFile(fileWriter, storage)
				if err != nil {
					logger.Log.Errorw("Error during saving metrics to file")
				}
				writeTicker.Reset(writeInterval)
			}
		}
	}()
}
