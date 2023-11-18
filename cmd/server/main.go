package main

import (
	"github.com/lev-stas/metricsmonitor.git/internal/configs"
	"github.com/lev-stas/metricsmonitor.git/internal/logger"
	"github.com/lev-stas/metricsmonitor.git/internal/memstorage"
	"github.com/lev-stas/metricsmonitor.git/internal/routers"
	"log"
	"net/http"
)

var storage *memstorage.MemStorage

func main() {
	if err := logger.LogInit(configs.ServerParams.LogLevel); err != nil {
		log.Fatalln(err)
	}
	configs.GetServerConfigs()
	storage = memstorage.NewMemStorage()
	r := routers.RootRouter(storage)
	log.Fatalln(http.ListenAndServe(configs.ServerParams.Host, logger.RequestResponseLogger(r)))

}
