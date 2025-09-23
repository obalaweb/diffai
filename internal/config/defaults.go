package config

import "github.com/diffai/diffai/pkg/types"

// GetDefaultConfig returns a default configuration
func GetDefaultConfig() *types.Config {
	return &types.Config{
		AI: types.AIConfig{
			Provider: "openai",
			Model:    "gpt-4",
			BaseURL:  "https://api.openai.com/v1",
			Options: map[string]string{
				"temperature": "0.7",
				"max_tokens":  "500",
			},
		},
		Git: types.GitConfig{
			ConventionalCommits: true,
			MaxCommitLength:     50,
			AutoCommit:          false,
			DefaultBranch:       "main",
		},
		Output: types.OutputConfig{
			Format:  "text",
			Verbose: false,
			Color:   true,
		},
		Service: types.ServiceConfig{
			Host:     "localhost",
			Port:     8080,
			Timeout:  30,
			Retries:  3,
		},
	}
}
