package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
)

type Config struct {
	Server ServerConfig `toml:"server"`

	Postgres PostgresConfig `toml:"postgres"`

	Redis RedisConfig `toml:"redis"`

	Log LogConfig `toml:"log"`
}

type ServerConfig struct {
	Port int `toml:"port"`
}

type PostgresConfig struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`

	Database string `toml:"database"`
}

type RedisConfig struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

type LogConfig struct {
	Level string `toml:"level"`
}

func Load() *Config {

	var cfg Config

	if _, err := toml.DecodeFile("app.toml", &cfg); err != nil {
		panic(err)
	}

	return &cfg
}

func LoadEnv() {

	err := godotenv.Load()

	if err != nil {
		log.Println(".env not found")
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
