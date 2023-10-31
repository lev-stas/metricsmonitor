package handlers

import (
	"reflect"
	"runtime"
)

func PickMetrics(metricsList []string) map[string]reflect.Value {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	memStatsType := reflect.TypeOf(m)

	metrics := make(map[string]reflect.Value)

	for _, metricName := range metricsList {
		if _, ok := memStatsType.FieldByName(metricName); ok {
			fieldValue := reflect.ValueOf(m).FieldByName(metricName)
			//value := fieldValue.Uint()
			metrics[metricName] = fieldValue
		}
	}

	return metrics
}
