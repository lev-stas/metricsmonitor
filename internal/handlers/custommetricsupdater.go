package handlers

import "math/rand"

type PollCountMetric struct {
	PollCount int64
}

type RandomValueMetric struct {
	RandomValue float64
}

type PollCountMetricInterface interface {
	IncrementPollCount()
	ResetPollCount()
}

type RandomValueMetricInterface interface {
	GenerateRandomValue()
}

func (metric *PollCountMetric) IncrementPollCount() {
	metric.PollCount += 1
}

func (metric *PollCountMetric) ResetPollCount() {
	metric.PollCount = 0
}

func (metric *RandomValueMetric) GenerateRandomValue() {
	metric.RandomValue = rand.Float64()
}
