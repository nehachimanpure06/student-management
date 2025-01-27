package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
	Log         `yaml:"log"`
	MysqlConfig `yaml:"mysql_config"`
	App         `yaml:"app"`
}

type HTTPServer struct {
	Port        string `yaml:"port" env:"HTTP_PORT" env-required:"true"`
	ReadTimeout int    `yaml:"read_timeout" env:"READ_TIMEOUT" env-required:"true" `
}

type Log struct {
	Level string `yaml:"log_level" env:"LOG_LEVEL" env-required:"true" `
}

type App struct {
	Name    string `env-required:"true" yaml:"name" env:"APP_NAME"`
	Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
}

type MysqlConfig struct {
	HostName string `env-required:"true" yaml:"db_host" env:"DB_HOST"`
	Port     int    `env-required:"true" yaml:"db_port" env:"DB_PORT"`
	Username string `env-required:"true" yaml:"db_username" env:"DB_USERNAME"`
	Password string `env-required:"true" yaml:"db_password" env:"DB_PASSWORD"`
	Database string `env-required:"true" yaml:"db_database" env:"DB_DATABASE"`
}

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yaml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
