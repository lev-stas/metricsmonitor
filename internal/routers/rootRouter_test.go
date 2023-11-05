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

	t.Run("Test main page", func(t *testing.T) {
		resp, err := client.R().Get(ts.URL)
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode())
		assert.Contains(t, string(resp.Body()), "<h1>Метрики</h1>")
		assert.Contains(t, string(resp.Body()), "TestGauge: 3.14")
		assert.Contains(t, string(resp.Body()), "TestCounter: 88")
	})

}
