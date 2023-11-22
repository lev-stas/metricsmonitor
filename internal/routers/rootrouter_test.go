package routers

import (
	"github.com/go-resty/resty/v2"
	"github.com/lev-stas/metricsmonitor.git/internal/memstorage"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootRouter(t *testing.T) {
	storage := memstorage.NewMemStorage()
	storage.SetGaugeMetric("TestGauge", 3.14)
	storage.SetCounterMetric("TestCounter", 88)
	ts := httptest.NewServer(RootRouter(storage))
	client := resty.New()
	updateUrl := ts.URL + "/update"
	valueUrl := ts.URL + "/value"

	var testCases = []struct {
		testName     string
		method       string
		requestBody  string
		expectedBody string
		status       int
	}{
		{
			"Test valid gauge request",
			http.MethodPost,
			`{
				"id": "GaugeMetric",
				"type": "gauge",
				"value": 3.14
			}`,
			`{
							"id": "GaugeMetric",
							"type": "gauge",
							"value": 3.14
						}`,

			200,
		},
		{
			"Test valid counter request",
			http.MethodPost,
			`{
					"id": "CounterMetric",
					"type": "counter",
					"delta": 3
					}`,
			`{
							"id": "CounterMetric",
							"type": "counter",
							"delta": 3
								}`,
			200,
		},
		{
			"Test no metric name request",
			http.MethodPost,
			`{
					"type": "gauge",
					"value": 3.14
					}`,
			"",
			404,
		},
		{
			"Test wrong metric type request",
			http.MethodPost,
			`{
					"id": "WrongMetric",
					"type": "wrongtype",
					"value": 3.12
					}`,
			"",
			400,
		},
		{
			"Test wrong gauge type request",
			http.MethodPost,
			`{
						"id": "WrongMetric",
						"type": "gauge",
						"value": "abc"
					}`,
			"",
			400,
		},
		{
			"Test wrong counter type request",
			http.MethodPost,
			`{
					"id": "WrongMetric",
					"type": "counter",
					"delta": 3.14
					}`,
			"",
			400,
		},
	}
	for _, v := range testCases {
		t.Run(v.testName, func(t *testing.T) {
			request := resty.New().R()
			request.URL = updateUrl
			request.Method = v.method

			if len(v.requestBody) > 0 {
				request.SetHeader("Content-Type", "application/json")
				request.SetBody(v.requestBody)
			}
			resp, err := request.Send()
			assert.NoError(t, err)
			assert.Equal(t, v.status, resp.StatusCode())
			if v.expectedBody != "" {
				assert.JSONEq(t, v.expectedBody, string(resp.Body()))
			}
		})
	}

	t.Run("Wrong request method for update endpoint", func(t *testing.T) {
		resp, err := client.R().Get(updateUrl)
		assert.NoError(t, err)
		assert.Equal(t, 405, resp.StatusCode())
	})

	t.Run("Test main page", func(t *testing.T) {
		resp, err := client.R().Get(ts.URL)
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode())
		assert.Contains(t, string(resp.Body()), "<h1>Метрики</h1>")
		assert.Contains(t, string(resp.Body()), "TestGauge: 3.14")
		assert.Contains(t, string(resp.Body()), "TestCounter: 88")
	})
	t.Run("Get gauge metric", func(t *testing.T) {
		resp, err := client.R().SetBody(`{"id": "TestGauge", "type": "gauge"}`).Post(valueUrl)
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode())
		assert.JSONEq(t, `{"id": "TestGauge", "type": "gauge", "value": 3.14}`, string(resp.Body()))
	})
	t.Run("Get counter Metrics", func(t *testing.T) {
		resp, err := client.R().SetBody(`{"id": "TestCounter", "type": "counter"}`).Post(valueUrl)
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode())
		assert.JSONEq(t, `{"id": "TestCounter", "type": "counter", "delta": 88}`, string(resp.Body()))
	})
	t.Run("Get wrong metrics type", func(t *testing.T) {
		resp, err := client.R().SetBody(`{"id": "TestCounter", "type": "wrongType"}`).Post(valueUrl)
		assert.NoError(t, err)
		assert.Equal(t, 400, resp.StatusCode())
	})
	t.Run("Get missed metric", func(t *testing.T) {
		resp, err := client.R().SetBody(`{"id": "TestGauge", "type": "counter"}`).Post(valueUrl)
		assert.NoError(t, err)
		assert.Equal(t, 404, resp.StatusCode())
	})
	t.Run("Post method", func(t *testing.T) {
		resp, err := client.R().SetBody(`{"id": "TestCounter", "type": "counter"}`).Get(valueUrl)
		assert.NoError(t, err)
		assert.Equal(t, 405, resp.StatusCode())
	})

}
