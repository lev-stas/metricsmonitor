package routers

import (
	"github.com/go-resty/resty/v2"
	"github.com/lev-stas/metricsmonitor.git/internal/memstorage"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestRootRouter(t *testing.T) {
	storage := memstorage.NewMemStorage()
	storage.SetGaugeMetric("TestGauge", 3.14)
	storage.SetCounterMetric("TestCounter", 88)
	ts := httptest.NewServer(RootRouter(storage))
	client := resty.New()

	var testCases = []struct {
		testName string
		url      string
		status   int
	}{
		{"Test valid gauge request", "/update/gauge/GaugeMetric/3.14", 200},
		{"Test valid counter request", "/update/counter/CounterMetric/3", 200},
		{"Test no metric name request", "/update/gauge//435", 404},
		{"Test no metric valid request", "/update/gauge/MetricName", 404},
		{"Test wrong metric type request", "/update/wrongType/WrongMetric/435", 400},
		{"Test wrong gauge type request", "/update/gauge/WrongMetric/abc", 400},
		{"Test wrong counter type request", "/update/counter/WrongMetric/3.14", 400},
	}
	for _, v := range testCases {
		t.Run(v.testName, func(t *testing.T) {
			resp, err := client.R().Post(ts.URL + v.url)
			assert.NoError(t, err)
			assert.Equal(t, v.status, resp.StatusCode())
		})
		t.Run("Wrong request method", func(t *testing.T) {
			resp, err := client.R().Get(ts.URL + "/update/counter/CounterMetric/3")
			assert.NoError(t, err)
			assert.Equal(t, 405, resp.StatusCode())
		})
	}

	t.Run("Test main page", func(t *testing.T) {
		resp, err := client.R().Get(ts.URL)
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode())
		assert.Contains(t, string(resp.Body()), "<h1>Метрики</h1>")
		assert.Contains(t, string(resp.Body()), "TestGauge: 3.14")
		assert.Contains(t, string(resp.Body()), "TestCounter: 88")
	})
	t.Run("Get gauge metric", func(t *testing.T) {
		resp, err := client.R().Get(ts.URL + "/value/gauge/TestGauge")
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode())
		assert.Equal(t, "3.14", string(resp.Body()))
	})
	t.Run("Get counter Metrics", func(t *testing.T) {
		resp, err := client.R().Get(ts.URL + "/value/counter/TestCounter")
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode())
		assert.Equal(t, "88", string(resp.Body()))
	})
	t.Run("Get wrong metrics type", func(t *testing.T) {
		resp, err := client.R().Get(ts.URL + "/value/wrongMetrics/TestCounter")
		assert.NoError(t, err)
		assert.Equal(t, 400, resp.StatusCode())
	})
	t.Run("Get missed metric", func(t *testing.T) {
		resp, err := client.R().Get(ts.URL + "/value/counter/TestGauge")
		assert.NoError(t, err)
		assert.Equal(t, 404, resp.StatusCode())
	})
	t.Run("Post method", func(t *testing.T) {
		resp, err := client.R().Post(ts.URL + "/value/counter/TestCounter")
		assert.NoError(t, err)
		assert.Equal(t, 405, resp.StatusCode())
	})

}
