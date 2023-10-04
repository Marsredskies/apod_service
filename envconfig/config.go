package envconfig

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DB   Database
	APOD Apod
}

func GetConfig() (Config, error) {
	var cfg Config

	if err := envconfig.Process("apod", &cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}

func MustGetConfig() Config {
	cnf, err := GetConfig()
	if err != nil {
		panic(fmt.Errorf("failed to parse env: %v", err))
	}
	return cnf
}
