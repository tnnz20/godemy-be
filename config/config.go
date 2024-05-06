package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App      AppConfig      `yaml:"app"`
	Database DatabaseConfig `yaml:"database"`
}

type AppConfig struct {
	Name       string           `yaml:"name"`
	Port       string           `yaml:"port"`
	Encryption EncryptionConfig `yaml:"encryption"`
}

type EncryptionConfig struct {
	JWTSecret string `yaml:"jwt_secret"`
}

type DatabaseConfig struct {
	Postgres PostgresConfig `yaml:"postgres"`
}

type PostgresConfig struct {
	Host           string                       `yaml:"host"`
	Port           string                       `yaml:"port"`
	User           string                       `yaml:"user"`
	Password       string                       `yaml:"password"`
	Name           string                       `yaml:"name"`
	SSLMode        string                       `yaml:"SSLMode"`
	ConnectionPool PostgresConnectionPoolConfig `yaml:"connection_pool"`
}

type PostgresConnectionPoolConfig struct {
	MaxIdleConnection     uint8 `yaml:"max_idle_connection"`
	MaxOpenConnection     uint8 `yaml:"max_open_connection"`
	MaxLifetimeConnection uint8 `yaml:"max_lifetime_connection"`
	MaxIdleTimeConnection uint8 `yaml:"max_idle_time_connection"`
}

var Cfg Config

func Load(filename string) (err error) {

	configByte, err := os.ReadFile(filename)
	if err != nil {
		return
	}

	return yaml.Unmarshal(configByte, &Cfg)
}
