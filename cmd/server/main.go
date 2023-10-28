package main

import (
	"fmt"
	"github.com/lev-stas/metricsmonitor.git/internal/memstorage"
	"net/http"
	"strconv"
	"strings"
)

func HandleUpdate(w http.ResponseWriter, r *http.Request) {
	storage := memstorage.NewMemStorage()

	if r.Method != http.MethodPost {
		http.Error(w, "Method not Allows", http.StatusMethodNotAllowed)
		return
	}

	requestParts := strings.Split(r.URL.Path, "/")

	if len(requestParts) < 5 || requestParts[3] == "" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	metricsType := requestParts[2]
	metricsName := requestParts[3]
	metricsValueRaw := requestParts[4]

	if metricsType != "counter" && metricsType != "gauge" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if metricsType == "counter" {
		metricsValue, err := strconv.ParseInt(metricsValueRaw, 10, 64)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		storage.SetCounterMetric(metricsName, metricsValue)
	}

	if metricsType == "gauge" {
		metricsValue, err := strconv.ParseFloat(metricsValueRaw, 62)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		storage.SetGaugeMetric(metricsName, metricsValue)
	}

	fmt.Printf("Received metric update - Type: %s, Name: %s, Value: %s\n", metricsType, metricsName, metricsValueRaw)

	w.WriteHeader(http.StatusOK)

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/update/", HandleUpdate)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
