package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App App `toml:"app"`
}

type App struct {
	Protocol string `toml:"protocol"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
}

func GetConfig(service string) *Config {
	cfg := new(Config)
	cleanenv.ReadConfig(fmt.Sprintf("./%s/env.toml", service), cfg)
	return cfg
}
