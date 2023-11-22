package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/lev-stas/metricsmonitor.git/internal/handlers"
	"github.com/lev-stas/metricsmonitor.git/internal/memstorage"
	"net/http"
)

func RootRouter(storage *memstorage.MemStorage) http.Handler {
	r := chi.NewRouter()
	r.Get("/", handlers.RootHandler(storage))
	r.Post("/value", handlers.ValueHandler(storage))
	r.Post("/update", handlers.HandleUpdate(storage))

	return r
}
