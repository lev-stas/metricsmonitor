package handlers

import (
	"github.com/lev-stas/metricsmonitor.git/internal/memstorage"
	"html/template"
	"net/http"
)

func RootHandler(storage *memstorage.MemStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		gaugeMetrics := storage.GetAllGaugeMetrics()
		counterMetrics := storage.GetAllCounterMetrics()

		pageTemplate := `
        <!DOCTYPE html>
        <html>
            <head>
                <title>Metrics</title>
            </head>
            <body>
                <h1>Метрики</h1>
                <ul>
                    {{range $name, $value := .GaugeMetrics}}
    					<li>{{ $name }}: {{ index $.GaugeMetrics $name }}</li>
					{{end}}
                
                    {{range $name, $value := .CounterMetrics}}
    					<li>{{ $name }}: {{ index $.CounterMetrics $name }}</li>
					{{end}}
                </ul>
            </body>
        </html>
        `

		metrics := struct {
			GaugeMetrics   map[string]float64
			CounterMetrics map[string]int64
		}{
			GaugeMetrics:   gaugeMetrics,
			CounterMetrics: counterMetrics,
		}
		t, err := template.New("html").Parse(pageTemplate)
		if err != nil {
			http.Error(w, "Can't parse page", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := t.Execute(w, metrics); err != nil {
			http.Error(w, "Can't generate page", http.StatusInternalServerError)
			return
		}
	}
}
