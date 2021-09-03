package repository

import (
	"yaml-to-openmetric/internal/domain"
)

type Currencies interface {
	Read(string) (*domain.Currencies, error)
}

type Repositories struct {
	Currencies Currencies
}

func NewRepositories() *Repositories {
	return &Repositories{
		Currencies: NewCurrenciesRepo(),
	}
}
