package config

import "os"

const (
	EnvKeyDroneApiHost     = "DRONE_API_HOST"
	EnvKeyDroneApiGrpcPort = "DRONE_API_GRPC_PORT"

	EnvKeyMissionApiHost     = "MISSION_API_HOST"
	EnvKeyMissionApiGrpcPort = "MISSION_API_GRPC_PORT"
)

type DroneApiClientConfig struct {
	Host     string
	GrpcPort string
}

func NewDroneApiClientConfig() *DroneApiClientConfig {
	return &DroneApiClientConfig{
		Host:     os.Getenv(EnvKeyDroneApiHost),
		GrpcPort: os.Getenv(EnvKeyDroneApiGrpcPort),
	}
}

type MissionApiClientConfig struct {
	Host     string
	GrpcPort string
}

func NewMissionApiClientConfig() *MissionApiClientConfig {
	return &MissionApiClientConfig{
		Host:     os.Getenv(EnvKeyMissionApiHost),
		GrpcPort: os.Getenv(EnvKeyMissionApiGrpcPort),
	}
}
