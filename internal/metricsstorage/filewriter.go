package metricsstorage

import (
	"encoding/json"
	"github.com/lev-stas/metricsmonitor.git/internal/configs"
	"github.com/lev-stas/metricsmonitor.git/internal/datamodels"
	"github.com/lev-stas/metricsmonitor.git/internal/logger"
	"os"
	"sync"
)

type FileWriter struct {
	file     *os.File
	filename string
	mu       sync.Mutex
}

//type FileWriterInterface interface {
//	Write(metrics []datamodels.Metric) error
//	Close() error
//}

func NewFileWriter(config *configs.ServerConfigParams) (*FileWriter, error) {
	if config.SyncSave() {
		file, err := os.OpenFile(config.StorageFile, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			return nil, err
		}
		return &FileWriter{
			file: file,
		}, nil
	}
	return &FileWriter{
		filename: config.StorageFile,
	}, nil
}

func (w *FileWriter) Write(metrics []datamodels.Metric) error {
	w.mu.Lock()
	defer w.mu.Unlock()
	data := []byte{}

	for _, metric := range metrics {
		item, err := json.Marshal(&metric)
		if err != nil {
			logger.Log.Errorw("Error during marshalling data before writing")
			return err
		}
		data = append(data, item...)
		data = append(data, '\n')
	}

	if w.file != nil {
		if _, err := w.file.Write(data); err != nil {
			//logger.Log.Errorw("Error during writing to file")
			return err
		}
		return nil
	}

	if err := os.WriteFile(w.filename, data, 0666); err != nil {
		logger.Log.Errorw("Error during writing to file")
		return err
	}
	return nil
}

func (w *FileWriter) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	return w.file.Close()
}

func SaveMetricsToFile(fileWriter *FileWriter, storage *MemStorage) error {
	gaugeMetrics := storage.GetAllGaugeMetrics()
	counterMetrics := storage.GetAllCounterMetrics()
	metrics := []datamodels.Metric{}

	for id, value := range gaugeMetrics {
		newValue := value
		metric := datamodels.Metric{
			ID:    id,
			MType: "gauge",
			Value: &newValue,
		}
		metrics = append(metrics, metric)
	}

	for id, delta := range counterMetrics {
		newDelta := delta
		metric := datamodels.Metric{
			ID:    id,
			MType: "counter",
			Delta: &newDelta,
		}

		metrics = append(metrics, metric)
	}

	err := fileWriter.Write(metrics)
	if err != nil {
		return err
	}
	return nil
}
