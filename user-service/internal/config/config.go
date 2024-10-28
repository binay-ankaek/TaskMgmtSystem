package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DBHost     string
	DBUser     string
	DBName     string
	DBPassword string
	Port       string
	DBSSLMode  string
	GRPCPort   string
}

func LoadConfig() (*Config, error) {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
	return &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		Port:       os.Getenv("APP_PORT"),
		DBSSLMode:  os.Getenv("DB_SSLMODE"),
		GRPCPort:   os.Getenv("GRPC_PORT"),
	}, nil
}
