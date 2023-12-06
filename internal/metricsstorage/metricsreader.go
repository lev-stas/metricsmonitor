package metricsstorage

import (
	"bufio"
	"bytes"
	"encoding/json"
	"github.com/lev-stas/metricsmonitor.git/internal/datamodels"
	"os"
)

type FileReader struct {
	filename string
	scanner  *bufio.Scanner
}

func NewMetricsReader(filename string) (*FileReader, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	file.Close()

	return &FileReader{
		filename: filename,
		scanner:  bufio.NewScanner(file),
	}, nil
}

func (r *FileReader) ReadFile() ([]datamodels.Metric, error) {
	data, err := os.ReadFile(r.filename)
	if err != nil {
		return nil, err
	}

	metrics := []datamodels.Metric{}
	lines := bytes.Split(data, []byte("\n"))
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		metric := datamodels.Metric{}
		err = json.Unmarshal(line, &metric)
		if err != nil {
			return nil, err
		}
		metrics = append(metrics, metric)
	}

	return metrics, nil

}
