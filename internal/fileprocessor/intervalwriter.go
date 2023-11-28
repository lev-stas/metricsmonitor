package handlers

import (
	"github.com/lev-stas/metricsmonitor.git/internal/configs"
	"github.com/lev-stas/metricsmonitor.git/internal/memstorage"
	"time"
)

func MetricsWriter(storage memstorage.MemStorage, file string) {
	storageInterval := time.Second * time.Duration(configs.ServerParams.StorageInterval)
	writerTicker := time.NewTicker(storageInterval)

	defer writerTicker.Stop()

	for {
		select {
		case <-writerTicker.C:
			
		}
	}
}
