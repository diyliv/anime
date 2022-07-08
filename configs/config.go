package configs

import (
	"log"

	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   Server
	Postgres Postgres
}

type Server struct {
	Host              string
	Port              string
	Development       bool
	Timeout           time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	MaxConnectionIdle time.Duration
	MaxConnectionAge  time.Duration
}

type Postgres struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func ReadConfig() *Config {
	var cfg Config

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error while reading config: " + err.Error())
	}
	err := viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("Error while unmarshalling config file: " + err.Error())
	}
	return &cfg
}
