package app

import (
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Environment           string `env:"ENVIRONMENT,required"`
	IsDebuggingEnabled    bool   `env:"ENABLE_DEBUGGING"`
	CabDbConnectionString string `env:"CAB_DB_CONNECTION_STRING"`
	RedisURI              string `env:"REDIS_URI"`
	RedisPassword         string `env:"REDIS_PASSWORD"`
	RedisDb               int    `env:"REDIS_DB"`
	ApiPort               string `env:"API_PORT" envDefault:":8080"`
	GrpcPort              string `env:"GRPC_PORT" envDefault:":8081"`
	GrpcHost              string `env:"GRPC_HOST" envDefault:"localhost"`
}

func LoadConfig() *Config {

	initialEnv := os.Getenv("ENVIRONMENT")
	if len(strings.TrimSpace(initialEnv)) == 0 {
		log.Fatal("'ENVIRONMENT' variable not set, exiting.")
	}
	cfg := Config{}

	if _, err := toml.DecodeFile("configs/"+initialEnv+".toml", &cfg); err != nil {
		log.Fatalf("Could not load %s config with error: %s", err.Error())
	}

	err := env.Parse(&cfg)

	if err != nil {
		log.Fatalf("Failed to load env variables. %+v\n", err)
	}

	return &cfg
}
