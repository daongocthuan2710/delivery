package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		GRPC `yaml:"grpc"`
		Log  `yaml:"logger"`
		PG   `yaml:"postgres"`
		GHN
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
		ENV     string `env:"ENV"`
		Debug   bool   `env:"DEBUG"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// GRPC -.
	GRPC struct {
		Port     string `env-required:"true" yaml:"port" env:"GRPC_PORT"`
		OrderAdd string `yaml:"order_add" env:"GRPC_ORDER_ADD"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	// PG -.
	PG struct {
		PoolMax int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
		URL     string `env-required:"true"                 env:"PG_URL"`
	}

	GHN struct {
		Token  string `env:"GHN_TOKEN"`
		ShopID string `env:"GHN_SHOP_ID"`
	}
)

// Init returns app config.
func Init() (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
