package metricsstorage

import (
	"bufio"
	"encoding/json"
	"github.com/lev-stas/metricsmonitor.git/internal/datamodels"
	"os"
)

type FileReader struct {
	file    *os.File
	scanner *bufio.Scanner
}

func NewMetricsReader(filename string) (*FileReader, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	return &FileReader{
		file:    file,
		scanner: bufio.NewScanner(file),
	}, nil
}

func (r *FileReader) ReadMetric() (*datamodels.Metric, error) {
	if !r.scanner.Scan() {
		return nil, r.scanner.Err()
	}
	data := r.scanner.Bytes()

	metric := datamodels.Metric{}

	err := json.Unmarshal(data, &metric)
	if err != nil {
		return nil, err
	}
	return &metric, nil
}

func (r *FileReader) Close() error {
	return r.file.Close()
}
