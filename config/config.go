package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GRPCServerAddress string
	APIerverAddress   string
	SensorType        string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file:", err)
	}

	return &Config{
		GRPCServerAddress: getEnv("GRPC_SERVER", "0.0.0.0:50051"),
		APIerverAddress:   getEnv("API_SERVER", "0.0.0.0:3000"),
		SensorType:        getEnv("SENSOR_TYPE", "DEFAULT"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
