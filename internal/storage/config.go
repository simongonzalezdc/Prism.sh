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
	// TODO: Implement cross-platform config directory detection
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
	// TODO: Implement config loading
	// 1. Get config directory
	// 2. Check if config.toml exists
	// 3. Parse TOML file
	// 4. Return Config struct
	// 5. If file doesn't exist, return default config

	return &Config{
		Theme:            "amber-night",
		AutoSave:         true,
		AutoSaveInterval: 30,
	}, nil
}

// SaveConfig saves configuration to file
func SaveConfig(cfg *Config) error {
	// TODO: Implement config saving
	// 1. Get config directory
	// 2. Create directory if it doesn't exist (os.MkdirAll)
	// 3. Marshal config to TOML
	// 4. Write to config.toml using AtomicWrite
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
