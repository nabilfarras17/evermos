package config

import (
	"github.com/kelseyhightower/envconfig"
	"time"
)

var conf Config

type Config struct {
	Port           string        `envconfig:"PORT" default:"9500"`
	ShutdownPeriod time.Duration `envconfig:"SHUTDOWN_PERIOD" default:"5s"`
}

func InitConfig() *Config {
	envconfig.Process("", &conf)
	return &conf
}
