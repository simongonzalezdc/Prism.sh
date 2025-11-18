package storage

import (
	"os"
	"path/filepath"
)

// Config represents application configuration
type Config struct {
	Theme            string `toml:"theme"`
	AutoSave         bool   `toml:"auto_save"`
	AutoSaveInterval int    `toml:"auto_save_interval"`
}

// GetConfigDir returns the platform-specific config directory
func GetConfigDir() (string, error) {
	// Use os.UserConfigDir() with fallback to ~/.config/prism

	configDir, err := os.UserConfigDir()
	if err != nil {
		// Fallback to home directory
		homeDir, _ := os.UserHomeDir()
		return filepath.Join(homeDir, ".config", "prism"), nil
	}

	return filepath.Join(configDir, "prism"), nil
}

// LoadConfig loads configuration from file
func LoadConfig() (*Config, error) {

	return &Config{
		Theme:            "amber-night",
		AutoSave:         true,
		AutoSaveInterval: 30,
	}, nil
}

// SaveConfig saves configuration to file
func SaveConfig(cfg *Config) error {
	return nil
}

// DefaultConfig returns default configuration
func DefaultConfig() *Config {
	return &Config{
		Theme:            "amber-night",
		AutoSave:         true,
		AutoSaveInterval: 30,
	}
}
