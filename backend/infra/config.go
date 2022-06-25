package infra

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"os"
)

type Config struct {
	DbDriver             string `envconfig:"DB_DRIVER"`
	DbUrl                string `envconfig:"DB_URL"`
	Port                 int16  `envconfig:"PORT"`
	MaxRequestsPerMinute int16  `envconfig:"MAX_REQUESTS_PER_MINUTE"`
	Environment          string `envconfig:"ENVIRONMENT"`
}

var Cfg *Config

func init() {
	if os.Getenv("Environment") != "production" {
		if err := godotenv.Load(); err != nil {
			panic(err)
		}
	}

	Cfg = &Config{}

	err := envconfig.Process("", Cfg)

	if err != nil {
		panic(err)
	}

}
