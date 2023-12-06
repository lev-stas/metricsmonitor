package preloadworkers

import (
	"github.com/lev-stas/metricsmonitor.git/internal/configs"
	"github.com/lev-stas/metricsmonitor.git/internal/logger"
	"github.com/lev-stas/metricsmonitor.git/internal/metricsstorage"
	"time"
)

func WriteMetricsTicker(storage *metricsstorage.MemStorage, fileWriter metricsstorage.FileWriterInterface) {
	writeInterval := time.Second * time.Duration(configs.ServerParams.StorageInterval)
	writeTicker := time.NewTicker(writeInterval)

	go func() {
		for {
			select {
			case <-writeTicker.C:
				err := metricsstorage.SaveMetricsToFile(fileWriter, storage)
				defer fileWriter.Close()
				if err != nil {
					logger.Log.Errorw("Error during saving metrics to file")
				}
				writeTicker.Reset(writeInterval)
			}
		}
	}()
}
