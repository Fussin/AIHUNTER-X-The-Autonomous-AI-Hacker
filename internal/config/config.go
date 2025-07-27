package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Config holds the application's configuration.
type Config struct {
	Log   LogConfig `yaml:"log"`
	Queue Queue     `yaml:"queue"`
}

// LogConfig holds the logging configuration.
type LogConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}

type Queue struct {
	Kafka KafkaConfig `yaml:"kafka"`
	Redis RedisConfig `yaml:"redis"`
}

type KafkaConfig struct {
	Brokers []string `yaml:"brokers"`
}

type RedisConfig struct {
	Address string `yaml:"address"`
}

// Load loads the configuration from the given path.
func Load(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
