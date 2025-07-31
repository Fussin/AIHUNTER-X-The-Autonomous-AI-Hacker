package config

import (
	"strings"

	"github.com/spf13/viper"
)

// Config holds the application's configuration.
type Config struct {
	// Log is the logging configuration.
	Log LogConfig `mapstructure:"log" validate:"required"`
	// Queue is the queue configuration.
	Queue Queue `mapstructure:"queue" validate:"required"`
}

// LogConfig holds the logging configuration.
type LogConfig struct {
	// Level is the logging level.
	Level string `mapstructure:"level" validate:"required,oneof=debug info warn error fatal panic" default:"info"`
	// Format is the logging format.
	Format string `mapstructure:"format" validate:"required,oneof=console json" default:"console"`
	// File is the path to the log file.
	File string `mapstructure:"file"`
	// Rotate is the log rotation configuration.
	Rotate RotateConfig `mapstructure:"rotate"`
	// SessionID is the session ID.
	SessionID string `mapstructure:"session_id"`
}

// RotateConfig holds the log rotation configuration.
type RotateConfig struct {
	// MaxSize is the maximum size in megabytes of the log file before it gets
	// rotated.
	MaxSize int `mapstructure:"maxsize"`
	// MaxBackups is the maximum number of old log files to retain.
	MaxBackups int `mapstructure:"maxbackups"`
	// MaxAge is the maximum number of days to retain old log files based on the
	// timestamp encoded in their filename.
	MaxAge int `mapstructure:"maxage"`
	// Compress is whether to compress the rotated log files.
	Compress bool `mapstructure:"compress"`
}

// Queue holds the queue configuration.
type Queue struct {
	// Kafka is the Kafka configuration.
	Kafka KafkaConfig `mapstructure:"kafka" validate:"required"`
	// Redis is the Redis configuration.
	Redis RedisConfig `mapstructure:"redis" validate:"required"`
}

// KafkaConfig holds the Kafka configuration.
type KafkaConfig struct {
	// Brokers is a list of Kafka brokers.
	Brokers []string `mapstructure:"brokers" validate:"required"`
}

// RedisConfig holds the Redis configuration.
type RedisConfig struct {
	// Address is the Redis address.
	Address string `mapstructure:"address" validate:"required,hostname_port"`
}

// Load loads the configuration from the given path.
// It uses viper to load the configuration from a file, environment variables, and command-line flags.
func Load(path string) (*Config, error) {
	viper.SetDefault("log.level", "info")
	viper.SetDefault("log.format", "console")

	viper.SetConfigFile(path)
	viper.SetConfigType("yaml") // or json, toml, etc.
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
