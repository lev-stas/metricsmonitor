package memstorage

type MemStorage struct {
	gaugeMetrics   map[string]float64
	counterMetrics map[string]int64
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		gaugeMetrics:   make(map[string]float64),
		counterMetrics: make(map[string]int64),
	}
}

func (storage *MemStorage) SetGaugeMetric(metric string, value float64) {
	storage.gaugeMetrics[metric] = value
}

func (storage *MemStorage) SetCounterMetric(metric string, value int64) {
	storage.counterMetrics[metric] += value
}

func (storage *MemStorage) GetGaugeMetric(metric string) (float64, bool) {
	value, found := storage.gaugeMetrics[metric]
	return value, found
}

func (storage *MemStorage) GetCounterMetric(metric string) (int64, bool) {
	value, found := storage.counterMetrics[metric]
	return value, found
}

func (storage *MemStorage) GetAllGaugeMetrics() map[string]float64 {
	return storage.gaugeMetrics
}

func (storage *MemStorage) GetAllCounterMetrics() map[string]int64 {
	return storage.counterMetrics
}
