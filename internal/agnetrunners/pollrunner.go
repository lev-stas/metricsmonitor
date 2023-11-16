package agnetrunners

import (
	"math/rand"
	"runtime"
)

func PollRunner(metricsList []string, metrics *map[string]float64, pollCount *PollCountMetric) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	for _, metric := range metricsList {
		switch metric {
		case "Alloc":
			(*metrics)[metric] = float64(m.Alloc)
		case "BuckHashSys":
			(*metrics)[metric] = float64(m.BuckHashSys)
		case "Frees":
			(*metrics)[metric] = float64(m.Frees)
		case "GCCPUFraction":
			(*metrics)[metric] = m.GCCPUFraction
		case "GCSys":
			(*metrics)[metric] = float64(m.GCSys)
		case "HeapAlloc":
			(*metrics)[metric] = float64(m.HeapAlloc)
		case "HeapIdle":
			(*metrics)[metric] = float64(m.HeapIdle)
		case "HeapInuse":
			(*metrics)[metric] = float64(m.HeapInuse)
		case "HeapObjects":
			(*metrics)[metric] = float64(m.HeapObjects)
		case "HeapReleased":
			(*metrics)[metric] = float64(m.HeapReleased)
		case "HeapSys":
			(*metrics)[metric] = float64(m.HeapSys)
		case "LastGC":
			(*metrics)[metric] = float64(m.LastGC)
		case "Lookups":
			(*metrics)[metric] = float64(m.Lookups)
		case "MCacheInuse":
			(*metrics)[metric] = float64(m.MCacheInuse)
		case "MCacheSys":
			(*metrics)[metric] = float64(m.MCacheSys)
		case "MSpanInuse":
			(*metrics)[metric] = float64(m.MSpanInuse)
		case "MSpanSys":
			(*metrics)[metric] = float64(m.MSpanSys)
		case "Mallocs":
			(*metrics)[metric] = float64(m.Mallocs)
		case "NextGC":
			(*metrics)[metric] = float64(m.NextGC)
		case "NumForcedGC":
			(*metrics)[metric] = float64(m.NumForcedGC)
		case "NumGC":
			(*metrics)[metric] = float64(m.NumGC)
		case "OtherSys":
			(*metrics)[metric] = float64(m.OtherSys)
		case "PauseTotalNs":
			(*metrics)[metric] = float64(m.PauseTotalNs)
		case "StackInuse":
			(*metrics)[metric] = float64(m.StackInuse)
		case "StackSys":
			(*metrics)[metric] = float64(m.StackSys)
		case "Sys":
			(*metrics)[metric] = float64(m.StackSys)
		case "TotalAlloc":
			(*metrics)[metric] = float64(m.StackSys)
		}
	}
	(*metrics)["RandomValue"] = rand.Float64()
	pollCount.IncrementPollCount()

}
