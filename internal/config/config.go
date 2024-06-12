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

const PathToConfig = "configs/config.yaml"

func GetConfig() Config {
	f, err := os.Open(PathToConfig)
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
	return cfg
}
