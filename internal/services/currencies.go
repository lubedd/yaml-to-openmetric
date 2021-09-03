package services

import (
	"github.com/prometheus/client_golang/prometheus"
	"yaml-to-openmetric/internal/config"
	"yaml-to-openmetric/internal/repository"
)

type currenciesService struct {
	currenciesRepo repository.Currencies
}

func NewCurrenciesService(currenciesRepo repository.Currencies) *currenciesService {
	return &currenciesService{
		currenciesRepo,
	}
}

func (c currenciesService) GetCurrenciesMetrics() (*prometheus.Registry, error) {
	r := prometheus.NewRegistry()

	cur, err := c.currenciesRepo.Read(config.Config.Metrics.Currencies)
	if err != nil {
		return r, err
	}

	for _, curr := range cur.Currencies {
		currItem := prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: curr.Name,
			})
		currItem.Set(float64(curr.Value))
		r.MustRegister(currItem)
	}

	return r, nil
}
