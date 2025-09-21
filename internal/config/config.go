// Copyright 2025 (C) https://github.com/3ziye team@3ziye.com

package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config Define server configuration structure
type Config struct {
	ServerPort string `json:"ServerPort"`
	RootDir    string `json:"RootDir"`
	Username   string `json:"Username"`
	Password   string `json:"Password"`
	LogLevel   string `json:"LogLevel"`
	LogFile    string `json:"LogFile"`
}

// LoadConfig Load configuration from file and environment variables
func LoadConfig(configPath string) (*Config, error) {
	// Read configuration file
	configData, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("Failed to read config file: %v", err)
	}

	// Parse JSON configuration
	var config Config
	if err := json.Unmarshal(configData, &config); err != nil {
		return nil, fmt.Errorf("Failed to parse config file: %v", err)
	}

	// Override with environment variables if they exist
	if port := os.Getenv("SERVER_PORT"); port != "" {
		config.ServerPort = port
	}

	if rootDir := os.Getenv("ROOT_DIR"); rootDir != "" {
		config.RootDir = rootDir
	}

	if username := os.Getenv("USERNAME"); username != "" {
		config.Username = username
	}

	if password := os.Getenv("PASSWORD"); password != "" {
		config.Password = password
	}

	if logLevel := os.Getenv("LOG_LEVEL"); logLevel != "" {
		config.LogLevel = logLevel
	}

	if logFile := os.Getenv("LOG_FILE"); logFile != "" {
		config.LogFile = logFile
	}

	return &config, nil
}

// GetDefaultConfig Return default configuration
func GetDefaultConfig() *Config {
	return &Config{
		ServerPort: "8080",
		RootDir:    "./uploads",
		Username:   "admin",
		Password:   "admin123",
		LogLevel:   "INFO",
		LogFile:    "",
	}
}