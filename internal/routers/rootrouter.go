package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/lev-stas/metricsmonitor.git/internal/handlers"
	"github.com/lev-stas/metricsmonitor.git/internal/metricsstorage"
	"net/http"
)

func RootRouter(storage *metricsstorage.MemStorage) http.Handler {
	r := chi.NewRouter()
	r.Get("/", handlers.RootHandler(storage))
	r.Post("/value/", handlers.ValueHandlerJSON(storage))
	r.Get("/value/{metricsType}/{metricsName}", handlers.ValueHandler(storage))
	r.Post("/update/", handlers.HandleUpdateJSON(storage))
	r.Post("/update/gauge/{metricsName}/{metricsValue}", handlers.HandleGaugeUpdate(storage))
	r.Post("/update/counter/{metricsName}/{metricsValue}", handlers.HandleCounterUpdate(storage))

	return r
}
