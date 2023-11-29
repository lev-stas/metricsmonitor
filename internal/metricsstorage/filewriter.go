package metricsstorage

import (
	"bufio"
	"encoding/json"
	"github.com/lev-stas/metricsmonitor.git/internal/datamodels"
	"os"
)

type FileWriter struct {
	file   *os.File
	writer *bufio.Writer
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

func (w *FileWriter) Write(metrics *[]datamodels.Metric) error {
	data, err := json.MarshalIndent(&metrics, "", "    ")
	if err != nil {
		return err
	}
	data = append(data, '\n')

	if _, err := w.writer.Write(data); err != nil {
		return err
	}

	if err := w.writer.WriteByte('\n'); err != nil {
		return err
	}
	return w.writer.Flush()
}

func (w *FileWriter) Close() error {
	return w.file.Close()
}
