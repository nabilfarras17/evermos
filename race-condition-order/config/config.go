package config

import (
	"github.com/kelseyhightower/envconfig"
	"time"
)

var conf Config

type Config struct {
	Port           string        `envconfig:"PORT" default:"9500"`
	ShutdownPeriod time.Duration `envconfig:"SHUTDOWN_PERIOD" default:"5s"`

	// NSQ broker
	ConsumerName    string `envconfig:"CONSUMER_NAME" default:"order"`
	ConsumerEnabled bool   `envconfig:"CONSUMER_ENABLE" default:"true"`
	ConsumerHost    string `envconfig:"CONSUMER_HOST" default:"localhost:4160"`
	ConsumerTopic   string `envconfig:"CONSUMER_TOPIC" default:"order-save"`
	ProducerHost    string `envconfig:"PRODUCER_HOST" default:"localhost:4160"`
}

func InitConfig() *Config {
	envconfig.Process("", &conf)
	return &conf
}
