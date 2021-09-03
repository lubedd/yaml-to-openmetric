package services

import (
	"sync"
	"testing"
	"yaml-to-openmetric/internal/config"
	"yaml-to-openmetric/internal/repository"
)
var wg sync.WaitGroup
func TestFileReading(t *testing.T) {
	config.GetConfig("../../configs/config.json")

	repos := repository.NewRepositories()
	services := NewServices(Deps{
		Repos: repos,
	})
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go read(services)
	}
	wg.Wait()
}

func read(services *Services) error {
	defer wg.Done()
	_, err := services.Metrics.GetCurrenciesMetrics()
	return err
}
