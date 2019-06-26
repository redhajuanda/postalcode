package config

import (
	"os"
	"strconv"
	"strings"
)

type DBConfig struct {
	Username string
	Password string
	Host     string
	Database string
}

type Config struct {
	DB   DBConfig
	Port int
}

func New() *Config {
	return &Config{
		DB: DBConfig{
			Username: getEnv("DB_USERNAME", "root"),
			Password: getEnv("DB_PASSWORD", "rezareza123"),
			Host:     getEnv("DB_HOST", "localhost"),
			Database: getEnv("DB_NAME", "kodepos"),
		},
		Port: getEnvAsInt("PORT", 3000),
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// Simple helper function to read an environment variable into integer or return a default value
func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

// Helper to read an environment variable into a bool or return default value
func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

// Helper to read an environment variable into a string slice or return default value
func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := getEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}
