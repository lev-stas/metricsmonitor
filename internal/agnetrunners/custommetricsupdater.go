package agnetrunners

type PollCountMetric struct {
	PollCount int64
}

type PollCountMetricInterface interface {
	IncrementPollCount()
	ResetPollCount()
}

func (metric *PollCountMetric) IncrementPollCount() {
	metric.PollCount += 1
}

func (metric *PollCountMetric) ResetPollCount() {
	metric.PollCount = 0
}
