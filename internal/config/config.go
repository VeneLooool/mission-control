package config

import (
	"context"
	"os"
)

const (
	EnvKeyHttpPort = "MISSION_CONTROL_HTTP_PORT"
	EnvKeyGrpcPort = "MISSION_CONTROL_GRPC_PORT"
)

type Config struct {
	HttpPort string
	GrpcPort string

	KafkaConfig            *KafkaConfig
	DroneApiClientConfig   *DroneApiClientConfig
	MissionApiClientConfig *MissionApiClientConfig
}

func (c *Config) GetKafkaConfig() *KafkaConfig {
	return c.KafkaConfig
}

func (c *Config) GetDroneApiClientConfig() *DroneApiClientConfig {
	return c.DroneApiClientConfig
}

func (c *Config) GetMissionApiClientConfig() *MissionApiClientConfig {
	return c.MissionApiClientConfig
}

func New(ctx context.Context) (*Config, error) {
	return &Config{
		HttpPort: os.Getenv(EnvKeyHttpPort),
		GrpcPort: os.Getenv(EnvKeyGrpcPort),

		KafkaConfig:            NewKafkaConfig(),
		DroneApiClientConfig:   NewDroneApiClientConfig(),
		MissionApiClientConfig: NewMissionApiClientConfig(),
	}, nil
}
