package http

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)


func (h Handlers) Metrics() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		registry, err := h.metricService.GetCurrenciesMetrics()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
			return
		}

		prom := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
		prom.ServeHTTP(w, r)
	}
}
