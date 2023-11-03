package handlers

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleUpdate(t *testing.T) {
	t.Run("Test valid gauge metric request", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "/update/gauge/GaugeMetric/3.14", nil)
		w := httptest.NewRecorder()
		HandleUpdate(w, request)
		res := w.Result()
		assert.Equal(t, 200, res.StatusCode)
	})
	t.Run("Test valid counter metric request", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "/update/counter/CounterMetric/3", nil)
		w := httptest.NewRecorder()
		HandleUpdate(w, request)
		res := w.Result()
		assert.Equal(t, 200, res.StatusCode)
	})
	t.Run("Test wrong type request", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/update/gauge/SomeMetric/45fhjr", nil)
		w := httptest.NewRecorder()
		HandleUpdate(w, request)
		res := w.Result()
		assert.Equal(t, 405, res.StatusCode)
	})
	t.Run("Test no metric name request", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "/update/gauge//435", nil)
		w := httptest.NewRecorder()
		HandleUpdate(w, request)
		res := w.Result()
		assert.Equal(t, 404, res.StatusCode)
	})
	t.Run("Test no metric value request", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "/update/gauge/MetricName", nil)
		w := httptest.NewRecorder()
		HandleUpdate(w, request)
		res := w.Result()
		assert.Equal(t, 404, res.StatusCode)
	})
	t.Run("Test wrong type metric request", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "/update/wrongType/WrongMetric/435", nil)
		w := httptest.NewRecorder()
		HandleUpdate(w, request)
		res := w.Result()
		assert.Equal(t, 400, res.StatusCode)
	})
	t.Run("Test wrong gauge metric value request", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "/update/gauge/WrongMetric/abc", nil)
		w := httptest.NewRecorder()
		HandleUpdate(w, request)
		res := w.Result()
		assert.Equal(t, 400, res.StatusCode)
	})
	t.Run("Test wrong counter metric value request", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "/update/counter/WrongMetric/3.14", nil)
		w := httptest.NewRecorder()
		HandleUpdate(w, request)
		res := w.Result()
		assert.Equal(t, 400, res.StatusCode)
	})
}
