package config

import (
	"os"

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
	arr := [][]interface{}{
		{&cfg.Server.Host, "SERVER_HOST"},
		{&cfg.Server.Port, "SERVER_PORT"},
		{&cfg.Database.Host, "DATABASE_HOST"},
		{&cfg.Database.Port, "DATABASE_PORT"},
		{&cfg.Database.DBname, "DATABASE_NAME"},
	}
	updateConfigFieldByNameV2(arr)
}

func updateConfigFieldByNameV2(arr [][]interface{}) {
	for _, tup := range arr {
		val, ok := os.LookupEnv(tup[1].(string))
		if ok {
			tup[0] = val
		}
	}
}
