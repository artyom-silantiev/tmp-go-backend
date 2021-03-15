package config

import (
	"os"
	"strconv"
	"strings"
)

type Config struct {
	EnvironmentType string // developer | production

	Port int

	RootPublic           string
	RootFrontendSpa      string
	RootFrontendAdminSpa string
}

var config *Config

func GetConfig() *Config {
	if config == nil {
		config = createConfig()
	}
	return config
}

func createConfig() *Config {
	config := &Config{
		EnvironmentType: getEnv("ENV", "developer"),

		Port: getEnvAsInt("PORT", 3000),

		RootPublic:           getEnv("ROOT_PUBLIC", ""),
		RootFrontendSpa:      getEnv("ROOT_FRONTEND_SPA", ""),
		RootFrontendAdminSpa: getEnv("ROOT_FRONTEND_ADMIN_SPA", ""),
	}
	return config
}

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
