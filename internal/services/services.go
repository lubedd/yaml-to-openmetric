package services

import (
	"github.com/prometheus/client_golang/prometheus"
	"yaml-to-openmetric/internal/repository"
)

type Metrics interface {
	GetCurrenciesMetrics() (*prometheus.Registry, error)
}

type Services struct {
	Metrics        Metrics
}

type Deps struct {
	Repos    *repository.Repositories
}

func NewServices(deps Deps) *Services {
	return &Services{
		Metrics: NewCurrenciesService(deps.Repos.Currencies),
	}
}
