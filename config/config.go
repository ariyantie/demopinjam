package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Database struct {
	Host       string `yaml:"localhost"`
	Port       int    `yaml:"3306"`
	Name       string `yaml:"kredit_system"`
	Username   string `yaml:"admin"`
	Password   string `yaml:"Password123#@!"`
	ActivePool bool   `yaml:"true"`
	MaxPool    int    `yaml:"30"`
	MinPool    int    `yaml:"5"`
}

type Config struct {
	DB Database `yaml:"db"`
}

func ReadConfig() Config {
	data, err := os.ReadFile("config/app.yaml")
	if err != nil {
		panic(err)
	}

	// create a person struct and deserialize the data into that struct
	var configuration Config

	if err := yaml.Unmarshal(data, &configuration); err != nil {
		panic(err)
	}
	return configuration
}
