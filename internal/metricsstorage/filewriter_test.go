package metricsstorage

import (
	"encoding/json"
	"github.com/lev-stas/metricsmonitor.git/internal/datamodels"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestFileWriter(t *testing.T) {
	tempFile, err := os.CreateTemp("", "test_metrics")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	fileWriter, err := NewFileWriter(tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer fileWriter.Close()

	testMetric := datamodels.Metric{
		ID:    "test",
		MType: "gauge",
		Value: new(float64),
	}
	err = fileWriter.Write(testMetric)
	assert.Nil(t, err, "Error writing to file")

	data, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	var readMetric datamodels.Metric
	err = json.Unmarshal(data, &readMetric)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, testMetric, readMetric, "Written and read metrics do not match")
}

func TestSaveMetricsToFile(t *testing.T) {
	tempFile, err := os.CreateTemp("", "test_metrics")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	fileWriter, err := NewFileWriter(tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer fileWriter.Close()

	storage := &FakeStorage{
		GaugeMetrics: map[string]float64{"test_gauge": 42.0},
		CounterMetrics: map[string]int64{
			"test_counter": 10,
		},
	}

	err = SaveMetricsToFile(fileWriter, storage)
	assert.Nil(t, err, "Error saving metrics to file")

	data, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		var readMetric datamodels.Metric
		if err := json.Unmarshal([]byte(line), &readMetric); err != nil {
			t.Fatal(err)
		}
	}
}

type FakeStorage struct {
	GaugeMetrics   map[string]float64
	CounterMetrics map[string]int64
}

func (fs *FakeStorage) GetAllGaugeMetrics() map[string]float64 {
	return fs.GaugeMetrics
}

func (fs *FakeStorage) GetAllCounterMetrics() map[string]int64 {
	return fs.CounterMetrics
}
