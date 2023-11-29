package main

import (
	"github.com/lev-stas/metricsmonitor.git/internal/configs"
	"github.com/lev-stas/metricsmonitor.git/internal/gzipper"
	"github.com/lev-stas/metricsmonitor.git/internal/logger"
	"github.com/lev-stas/metricsmonitor.git/internal/routers"
	"log"
	"net/http"
)

var storage *storage.MemStorage

func main() {
	if err := logger.LogInit(configs.ServerParams.LogLevel); err != nil {
		log.Fatalln(err)
	}
	configs.GetServerConfigs()
	storage = storage.NewMemStorage()
	r := routers.RootRouter(storage)
	log.Fatalln(http.ListenAndServe(configs.ServerParams.Host, gzipper.GzipMiddleware(logger.RequestResponseLogger(r))))
}
