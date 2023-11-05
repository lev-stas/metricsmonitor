package routers

import (
	"github.com/go-resty/resty/v2"
	"github.com/lev-stas/metricsmonitor.git/internal/memstorage"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestUpdateRouter(t *testing.T) {
	storage := memstorage.NewMemStorage()
	ts := httptest.NewServer(UpdateRouter(storage))
	client := resty.New()

	var testCases = []struct {
		testName string
		url      string
		status   int
	}{
		{"Test valid gauge request", "/gauge/GaugeMetric/3.14", 200},
		{"Test valid counter request", "/counter/CounterMetric/3", 200},
		{"Test no metric name request", "/gauge//435", 404},
		{"Test no metric valid request", "/gauge/MetricName", 404},
		{"Test wrong metric type request", "/wrongType/WrongMetric/435", 400},
		{"Test wrong gauge type request", "/gauge/WrongMetric/abc", 400},
		{"Test wrong counter type request", "/counter/WrongMetric/3.14", 400},
	}
	for _, v := range testCases {
		t.Run(v.testName, func(t *testing.T) {
			resp, err := client.R().Post(ts.URL + v.url)
			assert.NoError(t, err)
			assert.Equal(t, v.status, resp.StatusCode())
		})
		t.Run("Wrong request method", func(t *testing.T) {
			resp, err := client.R().Get(ts.URL + "/counter/CounterMetric/3")
			assert.NoError(t, err)
			assert.Equal(t, 405, resp.StatusCode())
		})
	}
}
