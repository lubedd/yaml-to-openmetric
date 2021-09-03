package http

import (
	"net/http"
	"yaml-to-openmetric/internal/services"
)

type Handlers struct {
	metricService services.Metrics
}

func NewHandlers(services *services.Services) Handlers {
	return Handlers{
		metricService: services.Metrics,
	}
}

func (h Handlers) Init()  *http.ServeMux {

	mux := http.NewServeMux()

	mux.Handle("/metrics", h.Metrics())
	return mux
}