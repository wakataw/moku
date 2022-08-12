package config

import (
	"github.com/spf13/viper"
	"time"
)

type (
	Config struct {
		Env string
		Mysql
		HttpConfig `mapstructure:"Http"`
		Cache
		Redis
		Auth
		Limiter
	}

	Mysql struct {
		Host            string
		Protocol        string
		Port            string
		DbName          string
		Username        string
		Password        string
		MaxIdleConn     int
		MaxOPenConn     int
		ConnMaxLifeTime time.Duration
	}

	HttpConfig struct {
		Port           string
		MaxHeaderBytes int
	}

	Cache struct {
		TTL time.Duration
	}

	Redis struct {
		Host     string
		Protocol string
		Port     string
		DbName   string
	}

	Auth struct {
		AccessTokenTTL  time.Duration
		RefreshTokenTTL time.Duration
	}

	Limiter struct {
		RequestPerSecond int
		Burst            int
		TTL              time.Duration
	}
)

func NewConfig(configDir string) (*Config, error) {
	viper.AddConfigPath(configDir)
	viper.SetConfigName("config")
	var config Config

	if err := viper.ReadInConfig(); err != nil {
		return &config, err
	}

	err := viper.Unmarshal(&config)

	return &config, err
}
