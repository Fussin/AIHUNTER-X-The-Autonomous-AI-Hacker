package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/user/aihunter-x/aihunter-core-cli/configs/loader"
)

func TestValidate(t *testing.T) {
	// Test with a valid config
	validCfg := &config.Config{
		Log: config.LogConfig{
			Level:  "info",
			Format: "console",
		},
		Queue: config.Queue{
			Kafka: config.KafkaConfig{
				Brokers: []string{"localhost:9092"},
			},
			Redis: config.RedisConfig{
				Address: "localhost:6379",
			},
		},
	}
	err := Validate(validCfg)
	assert.NoError(t, err)

	// Test with an invalid config
	invalidCfg := &config.Config{
		Log: config.LogConfig{
			Level:  "invalid",
			Format: "console",
		},
		Queue: config.Queue{
			Kafka: config.KafkaConfig{
				Brokers: []string{"localhost:9092"},
			},
			Redis: config.RedisConfig{
				Address: "localhost:6379",
			},
		},
	}
	err = Validate(invalidCfg)
	assert.Error(t, err)
}
