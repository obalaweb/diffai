package cli

import (
	"fmt"
	"os"

	"github.com/diffai/diffai/internal/config"
	"github.com/spf13/cobra"
)

// NewConfigCommand creates the config command
func NewConfigCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manage DiffAI configuration",
		Long: `Manage DiffAI configuration settings.

This command allows you to view, set, and initialize configuration options
for DiffAI including AI provider settings, Git preferences, and output formatting.

Examples:
  git diffai config                    # Show current configuration
  git diffai config init               # Initialize default configuration
  git diffai config set ai.provider openai  # Set AI provider
  git diffai config get ai.model       # Get AI model setting`,
		RunE: runConfig,
	}

	// Add subcommands
	cmd.AddCommand(NewConfigInitCommand())
	cmd.AddCommand(NewConfigSetCommand())
	cmd.AddCommand(NewConfigGetCommand())
	cmd.AddCommand(NewConfigShowCommand())

	return cmd
}

func runConfig(cmd *cobra.Command, args []string) error {
	// Show current configuration by default
	return runConfigShow(cmd, args)
}

// NewConfigInitCommand creates the config init subcommand
func NewConfigInitCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize default configuration",
		Long:  "Create a default configuration file in the user's home directory.",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := config.InitConfig()
			if err != nil {
				return fmt.Errorf("failed to initialize configuration: %w", err)
			}
			
			configPath := config.GetConfigPath()
			fmt.Printf("✅ Configuration initialized at: %s\n", configPath)
			return nil
		},
	}
}

// NewConfigSetCommand creates the config set subcommand
func NewConfigSetCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "set <key> <value>",
		Short: "Set a configuration value",
		Long:  "Set a configuration value. Use dot notation for nested keys (e.g., ai.provider).",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			key := args[0]
			value := args[1]
			
			// TODO: Implement config setting
			fmt.Printf("Setting %s = %s (not yet implemented)\n", key, value)
			return nil
		},
	}
}

// NewConfigGetCommand creates the config get subcommand
func NewConfigGetCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "get <key>",
		Short: "Get a configuration value",
		Long:  "Get a configuration value. Use dot notation for nested keys (e.g., ai.provider).",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			key := args[0]
			
			// TODO: Implement config getting
			fmt.Printf("Getting %s (not yet implemented)\n", key)
			return nil
		},
	}
}

// NewConfigShowCommand creates the config show subcommand
func NewConfigShowCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "show",
		Short: "Show current configuration",
		Long:  "Display the current configuration settings.",
		RunE:  runConfigShow,
	}
}

func runConfigShow(cmd *cobra.Command, args []string) error {
	configManager := config.NewManager()
	cfg, err := configManager.Load()
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	// Display configuration
	fmt.Println("Current DiffAI Configuration:")
	fmt.Println("=============================")
	fmt.Printf("AI Provider: %s\n", cfg.AI.Provider)
	fmt.Printf("AI Model: %s\n", cfg.AI.Model)
	fmt.Printf("AI Base URL: %s\n", cfg.AI.BaseURL)
	fmt.Printf("Conventional Commits: %t\n", cfg.Git.ConventionalCommits)
	fmt.Printf("Max Commit Length: %d\n", cfg.Git.MaxCommitLength)
	fmt.Printf("Auto Commit: %t\n", cfg.Git.AutoCommit)
	fmt.Printf("Output Format: %s\n", cfg.Output.Format)
	fmt.Printf("Verbose Output: %t\n", cfg.Output.Verbose)
	fmt.Printf("Service Host: %s\n", cfg.Service.Host)
	fmt.Printf("Service Port: %d\n", cfg.Service.Port)

	// Show config file path
	configPath := config.GetConfigPath()
	if _, err := os.Stat(configPath); err == nil {
		fmt.Printf("\nConfiguration file: %s\n", configPath)
	} else {
		fmt.Println("\nNo configuration file found. Run 'git diffai config init' to create one.")
	}

	return nil
}
