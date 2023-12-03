package metricsstorage

import (
	"bufio"
	"encoding/json"
	"github.com/lev-stas/metricsmonitor.git/internal/datamodels"
	"github.com/lev-stas/metricsmonitor.git/internal/logger"
	"os"
	"sync"
)

type FileWriter struct {
	file   *os.File
	writer *bufio.Writer
	mu     sync.Mutex
}

type FileWriterInterface interface {
	Write(metric datamodels.Metric) error
	Close() error
}

type StorageInterface interface {
	GetAllGaugeMetrics() map[string]float64
	GetAllCounterMetrics() map[string]int64
}

func NewFileWriter(filename string) (*FileWriter, error) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	return &FileWriter{
		file:   file,
		writer: bufio.NewWriter(file),
	}, nil
}

func (w *FileWriter) Write(metric datamodels.Metric) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	data, err := json.Marshal(&metric)
	if err != nil {
		logger.Log.Errorw("Error during marshalling data before writing")
		return err
	}

	data = append(data, '\n')

	if _, err = w.writer.Write(data); err != nil {
		logger.Log.Errorw("Error during writing to file")
		return err
	}
	return w.writer.Flush()
}

func (w *FileWriter) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	return w.file.Close()
}

func SaveMetricsToFile(fileWriter FileWriterInterface, storage StorageInterface) error {
	gaugeMetrics := storage.GetAllGaugeMetrics()
	counterMetrics := storage.GetAllCounterMetrics()

	defer fileWriter.Close()

	for id, value := range gaugeMetrics {
		metric := datamodels.Metric{
			ID:    id,
			MType: "gauge",
			Value: &value,
		}
		err := fileWriter.Write(metric)
		if err != nil {
			return err
		}
	}

	for id, delta := range counterMetrics {
		metric := datamodels.Metric{
			ID:    id,
			MType: "counter",
			Delta: &delta,
		}
		err := fileWriter.Write(metric)
		if err != nil {
			return err
		}
	}
	fileWriter.Close()
	return nil
}
