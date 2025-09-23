package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/diffai/diffai/pkg/types"
	"github.com/spf13/viper"
)

const (
	DefaultConfigDir  = ".config/diffai"
	DefaultConfigFile = "config.yaml"
)

// Manager handles configuration loading and management
type Manager struct {
	config *types.Config
	viper  *viper.Viper
}

// NewManager creates a new configuration manager
func NewManager() *Manager {
	v := viper.New()
	
	// Set default values
	setDefaults(v)
	
	// Configure viper
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	
	// Add config paths
	homeDir, _ := os.UserHomeDir()
	configDir := filepath.Join(homeDir, DefaultConfigDir)
	v.AddConfigPath(configDir)
	v.AddConfigPath(".")
	v.AddConfigPath("./configs")
	
	// Enable environment variable support
	v.AutomaticEnv()
	v.SetEnvPrefix("DIFFAI")
	
	return &Manager{
		viper: v,
	}
}

// Load loads configuration from file and environment
func (m *Manager) Load() (*types.Config, error) {
	// Try to read config file
	if err := m.viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
		// Config file not found, use defaults
	}
	
	// Unmarshal into config struct
	var config types.Config
	if err := m.viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}
	
	// Validate configuration
	if err := m.validate(&config); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}
	
	m.config = &config
	return &config, nil
}

// Save saves configuration to file
func (m *Manager) Save(config *types.Config) error {
	homeDir, _ := os.UserHomeDir()
	configDir := filepath.Join(homeDir, DefaultConfigDir)
	
	// Create config directory if it doesn't exist
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("error creating config directory: %w", err)
	}
	
	// Set values in viper
	m.viper.Set("ai", config.AI)
	m.viper.Set("git", config.Git)
	m.viper.Set("output", config.Output)
	m.viper.Set("service", config.Service)
	
	// Write config file
	configFile := filepath.Join(configDir, DefaultConfigFile)
	return m.viper.WriteConfigAs(configFile)
}

// Get returns the current configuration
func (m *Manager) Get() *types.Config {
	return m.config
}

// setDefaults sets default configuration values
func setDefaults(v *viper.Viper) {
	// AI defaults
	v.SetDefault("ai.provider", "openai")
	v.SetDefault("ai.model", "gpt-4")
	v.SetDefault("ai.base_url", "https://api.openai.com/v1")
	
	// Git defaults
	v.SetDefault("git.conventional_commits", true)
	v.SetDefault("git.max_commit_length", 50)
	v.SetDefault("git.auto_commit", false)
	v.SetDefault("git.default_branch", "main")
	
	// Output defaults
	v.SetDefault("output.format", "text")
	v.SetDefault("output.verbose", false)
	v.SetDefault("output.color", true)
	
	// Service defaults
	v.SetDefault("service.host", "localhost")
	v.SetDefault("service.port", 8080)
	v.SetDefault("service.timeout_seconds", 30)
	v.SetDefault("service.retries", 3)
}

// validate validates the configuration
func (m *Manager) validate(config *types.Config) error {
	// Validate AI provider
	validProviders := map[string]bool{
		"openai":    true,
		"anthropic": true,
		"local":     true,
	}
	
	if !validProviders[config.AI.Provider] {
		return fmt.Errorf("invalid AI provider: %s", config.AI.Provider)
	}
	
	// Validate output format
	validFormats := map[string]bool{
		"text":     true,
		"json":     true,
		"markdown": true,
	}
	
	if !validFormats[config.Output.Format] {
		return fmt.Errorf("invalid output format: %s", config.Output.Format)
	}
	
	// Validate service port
	if config.Service.Port < 1 || config.Service.Port > 65535 {
		return fmt.Errorf("invalid service port: %d", config.Service.Port)
	}
	
	return nil
}

// GetConfigPath returns the path to the configuration file
func GetConfigPath() string {
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, DefaultConfigDir, DefaultConfigFile)
}

// InitConfig initializes a default configuration file
func InitConfig() error {
	manager := NewManager()
	config, err := manager.Load()
	if err != nil {
		return err
	}
	
	return manager.Save(config)
}
