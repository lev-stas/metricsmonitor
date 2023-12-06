package preloadworkers

import (
	"github.com/lev-stas/metricsmonitor.git/internal/configs"
	"github.com/lev-stas/metricsmonitor.git/internal/logger"
	"github.com/lev-stas/metricsmonitor.git/internal/metricsstorage"
)

type AddMetrics interface {
	Set(metric string, value float64)
	Inc(metric string, value int64)
}

func RestoreMetricsFromFile(storage *metricsstorage.MemStorage) error {
	fileReader, err := metricsstorage.NewMetricsReader(configs.ServerParams.StorageFile)
	if err != nil {
		logger.Log.Errorw("Error during creating Filereader", "error", err)
	}
	metrics, err := fileReader.ReadFile()
	if err != nil {
		logger.Log.Errorw("Error during reading metrics from file")
		return err
	}
	for _, metric := range metrics {
		switch metric.MType {
		case "gauge":
			storage.Set(metric.ID, *metric.Value)
		case "counter":
			storage.Inc(metric.ID, *metric.Delta)
		}
	}
	return nil
	//for {
	//	metric, err := fileReader.ReadFile()
	//	if err == io.EOF {
	//		break
	//	}
	//	if err != nil {
	//		logger.Log.Errorw("Error during reading metric", "error", err)
	//		return err
	//	}
	//	if metric != nil {
	//		switch metric.MType {
	//		case "gauge":
	//			storage.Set(metric.ID, *metric.Value)
	//		case "counter":
	//			storage.Inc(metric.ID, *metric.Delta)
	//		}
	//	} else {
	//		logger.Log.Errorw("Error during setting metrics")
	//		return err
	//	}
	//}
}
