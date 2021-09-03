package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Configuration struct {
	App struct {
		ServerAddr string `json:"server_addr"`
	} `json:"app"`
	Metrics struct {
		Currencies string `json:"currencies"`
	} `json:"metrics"`
}

var Config *Configuration

func GetConfig(path string) *Configuration {
	cfg := new(Configuration)
	err := cfg.read(path)
	if err != nil {
		log.Println("Error parsing config. Load default settings")
		cfg.loadDefaultConfig()
	}

	Config = cfg

	return cfg
}

func (cfg *Configuration) read(path string) error {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &cfg)
	if err != nil {
		return err
	}
	return nil
}

func (cfg *Configuration) loadDefaultConfig() {
	cfg.App.ServerAddr = "127.0.0.1:8080"
	cfg.Metrics.Currencies="./currencies.yaml"
}