package config

import (
	"os"
	"strconv"

	"github.com/go-yaml/yaml"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}
	Database struct {
		Host   string `yaml:"host"`
		Port   int    `yaml:"port"`
		DBname string `yaml:"dbname"`
	}
}

const configPath = "configs/config.yaml"

// For Debugging
// const configDebugPath = "/../../configs/config.yaml"

func GetConfig() Config {
	f, err := os.Open(configPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		panic(err)
	}
	// Update config using Env
	readEnv(&cfg)

	return cfg
}

func readEnv(cfg *Config) {
	updateConfigFieldByName(cfg, func(cfg *Config, value string) {
		cfg.Server.Host = value
	}, "SERVER_HOST")
	updateConfigFieldByName(cfg, func(cfg *Config, value string) {
		res, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		cfg.Server.Port = res
	}, "SERVER_PORT")

	updateConfigFieldByName(cfg, func(cfg *Config, value string) {
		cfg.Database.Host = value
	}, "DATABASE_HOST")
	updateConfigFieldByName(cfg, func(cfg *Config, value string) {
		res, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		cfg.Database.Port = res
	}, "DATABASE_PORT")

	updateConfigFieldByName(cfg, func(cfg *Config, value string) {
		cfg.Database.DBname = value
	}, "DATABASE_NAME")
}

func updateConfigFieldByName(cfg *Config, fnSetter func(cfg *Config, value string), name string) {
	val, ok := os.LookupEnv(name)
	if ok {
		fnSetter(cfg, val)
	}
}
