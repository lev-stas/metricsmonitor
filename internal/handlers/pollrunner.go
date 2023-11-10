package handlers

func PollMetricsRunner(metrics *map[string]float64, pollCount *PollCountMetric) {
	PickMetrics(RuntimeMetrics, metrics)
	pollCount.IncrementPollCount()
}
