package routers

import (
	"github.com/go-resty/resty/v2"
	"github.com/lev-stas/metricsmonitor.git/internal/memstorage"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestValueRouter(t *testing.T) {
	storage := memstorage.NewMemStorage()
	storage.SetGaugeMetric("TestGauge", 3.14)
	storage.SetCounterMetric("TestCounter", 42)
	ts := httptest.NewServer(ValueRouter(storage))
	client := resty.New()

	t.Run("Get gauge metric", func(t *testing.T) {
		resp, err := client.R().Get(ts.URL + "/gauge/TestGauge")
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode())
		assert.Equal(t, "3.14", string(resp.Body()))
	})
	t.Run("Get counter Metrics", func(t *testing.T) {
		resp, err := client.R().Get(ts.URL + "/counter/TestCounter")
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode())
		assert.Equal(t, "42", string(resp.Body()))
	})
	t.Run("Get wrong metrics type", func(t *testing.T) {
		resp, err := client.R().Get(ts.URL + "/wrongMetrics/TestCounter")
		assert.NoError(t, err)
		assert.Equal(t, 400, resp.StatusCode())
	})
	t.Run("Get missed metric", func(t *testing.T) {
		resp, err := client.R().Get(ts.URL + "/counter/TestGauge")
		assert.NoError(t, err)
		assert.Equal(t, 404, resp.StatusCode())
	})
	t.Run("Post method", func(t *testing.T) {
		resp, err := client.R().Post(ts.URL + "/counter/TestCounter")
		assert.NoError(t, err)
		assert.Equal(t, 405, resp.StatusCode())
	})
}
