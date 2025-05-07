package config

import "os"

const (
	EnvKeyKafkaHost = "KAFKA_INTERNAL_HOST"
	EnvKeyKafkaPort = "KAFKA_INTERNAL_PORT"
)

type KafkaConfig struct {
	KafkaHost string
	KafkaPort string
}

func NewKafkaConfig() *KafkaConfig {
	return &KafkaConfig{
		KafkaPort: os.Getenv(EnvKeyKafkaPort),
		KafkaHost: os.Getenv(EnvKeyKafkaHost),
	}
}
