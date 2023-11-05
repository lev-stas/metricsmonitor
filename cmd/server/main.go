package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/lev-stas/metricsmonitor.git/internal/memstorage"
	"github.com/lev-stas/metricsmonitor.git/internal/routers"
	"log"
	"net/http"
)

var storage *memstorage.MemStorage

func main() {
	storage = memstorage.NewMemStorage()
	//updateRouter := routers.UpdateRouter()
	//valueRouter := routers.ValueRouter()
	//
	r := chi.NewMux()
	r.Mount("/update", routers.UpdateRouter(storage))
	r.Mount("/value", routers.ValueRouter(storage))
	r.Mount("/", routers.RootRouter(storage))
	//r := routers.UpdateRouter(storage)
	log.Fatalln(http.ListenAndServe(":8080", r))

}
