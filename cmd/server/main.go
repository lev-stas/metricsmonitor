package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/lev-stas/metricsmonitor.git/internal/configs"
	"github.com/lev-stas/metricsmonitor.git/internal/memstorage"
	"github.com/lev-stas/metricsmonitor.git/internal/routers"
	"log"
	"net/http"
)

var storage *memstorage.MemStorage

func main() {
	configs.GetServerConfigs()
	storage = memstorage.NewMemStorage()
	r := chi.NewMux()
	r.Mount("/update", routers.UpdateRouter(storage))
	r.Mount("/value", routers.ValueRouter(storage))
	r.Mount("/", routers.RootRouter(storage))
	log.Fatalln(http.ListenAndServe(configs.ServerParams.Host, r))

}
