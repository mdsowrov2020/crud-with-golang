package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configurations Config

type Config struct {
	Version      string
	ServiceName  string
	HTTPPort     int
	JwtSecretKey string
}

func loadConfig() {
	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is required")
		os.Exit(1)
	}
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("Service name is required")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println(" Http port is required")
		os.Exit(1)
	}

	convertedHttpport, err := strconv.Atoi(httpPort)
	if err != nil {
		fmt.Println("failed to convert")
		return
	}
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		fmt.Println("Jwt secret is required")
		os.Exit(1)
	}
	configurations = Config{
		Version:      version,
		ServiceName:  serviceName,
		HTTPPort:     convertedHttpport,
		JwtSecretKey: jwtSecretKey,
	}
}

func GetConfig() Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load the env file: ", err)
		os.Exit(1)
	}
	loadConfig()
	return configurations
}
