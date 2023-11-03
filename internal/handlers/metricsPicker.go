package handlers

import (
	"reflect"
	"runtime"
)

func PickMetrics(metricsList []string) map[string]float64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	memStatsType := reflect.TypeOf(m)

	metrics := make(map[string]float64)

	for _, metricName := range metricsList {
		if _, ok := memStatsType.FieldByName(metricName); ok {
			fieldValue := reflect.ValueOf(m).FieldByName(metricName)
			switch true {
			case fieldValue.CanFloat():
				metrics[metricName] = fieldValue.Float()
			case fieldValue.CanUint():
				metrics[metricName] = float64(fieldValue.Uint())
			default:
				{
				}
			}
		}
	}

	return metrics
}
