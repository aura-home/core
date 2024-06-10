package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Core CoreConfig `yaml:"core"`
}

type CoreConfig struct {
	Logger LoggerConfig `yaml:"logger"`
}
type LoggerConfig struct {
	Enable bool   `yaml:"enable"  env-default:"enable"`
	Level  string `yaml:"level"  env-default:"debug"`
	Type   string `yaml:"type" env-default:"raw"`
}

func Load(path string) (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
