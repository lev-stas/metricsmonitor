package storage

type MemStorage struct {
	GaugeMetrics   map[string]float64
	CounterMetrics map[string]int64
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		GaugeMetrics:   make(map[string]float64),
		CounterMetrics: make(map[string]int64),
	}
}

func (storage *MemStorage) SetGaugeMetric(metric string, value float64) {
	storage.GaugeMetrics[metric] = value
}

func (storage *MemStorage) SetCounterMetric(metric string, value int64) {
	storage.CounterMetrics[metric] += value
}

func (storage *MemStorage) GetGaugeMetric(metric string) (float64, bool) {
	value, found := storage.GaugeMetrics[metric]
	return value, found
}

func (storage *MemStorage) GetCounterMetric(metric string) (int64, bool) {
	value, found := storage.CounterMetrics[metric]
	return value, found
}

func (storage *MemStorage) GetAllGaugeMetrics() map[string]float64 {
	return storage.GaugeMetrics
}

func (storage *MemStorage) GetAllCounterMetrics() map[string]int64 {
	return storage.CounterMetrics
}
