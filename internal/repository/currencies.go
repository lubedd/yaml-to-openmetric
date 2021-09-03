package repository

import (
	"github.com/gofrs/flock"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"sync"
	"yaml-to-openmetric/internal/domain"
)

type CurrenciesRepo struct {
	mutex *sync.Mutex
}

func NewCurrenciesRepo() *CurrenciesRepo {
	return &CurrenciesRepo{&sync.Mutex{}}
}

func (c CurrenciesRepo) Read(pathToYAMLFile string) (*domain.Currencies, error) {
	yamlFile, err := c.readFile(pathToYAMLFile)
	if err != nil {
		return nil, err
	}

	var cur *domain.Currencies
	if err = yaml.Unmarshal(yamlFile, &cur); err != nil {
		return nil, err
	}

	return cur, nil
}

func (c CurrenciesRepo) readFile(pathToYAMLFile string) ([]byte, error) {
	fileLock := flock.New(pathToYAMLFile)

	locked, err := fileLock.TryRLock()
	if err != nil {
		return nil, err
	}

	c.mutex.Lock()
	yamlFile, err := ioutil.ReadFile(pathToYAMLFile)
	c.mutex.Unlock()

	if locked {
		fileLock.Unlock()
	}
	if err != nil {
		return nil, err
	}

	return yamlFile, nil
}
