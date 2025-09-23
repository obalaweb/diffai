package cli

import (
	"fmt"

	"github.com/diffai/diffai/internal/ai"
	"github.com/diffai/diffai/internal/config"
	"github.com/diffai/diffai/internal/git"
	"github.com/spf13/cobra"
)

// NewChangelogCommand creates the changelog command
func NewChangelogCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "changelog",
		Short: "Generate AI-powered changelog from recent commits",
		Long: `Generate a changelog from recent commits using AI.

This command analyzes recent commits and generates a structured changelog
grouped by features, fixes, documentation, and other categories.

Examples:
  git diffai changelog                    # Generate changelog from last 10 commits
  git diffai changelog --since v1.0.0     # Generate changelog since tag v1.0.0
  git diffai changelog --count 20          # Generate changelog from last 20 commits
  git diffai changelog --output CHANGELOG.md # Write to file`,
		RunE: runChangelog,
	}

	// Add flags
	cmd.Flags().StringP("since", "s", "", "Generate changelog since this tag or commit")
	cmd.Flags().IntP("count", "c", 10, "Number of commits to analyze")
	cmd.Flags().StringP("output", "o", "", "Output file (default: stdout)")
	cmd.Flags().StringP("format", "f", "markdown", "Output format (markdown, text, json)")
	cmd.Flags().BoolP("include-unreleased", "u", false, "Include unreleased changes")

	return cmd
}

func runChangelog(cmd *cobra.Command, args []string) error {
	// Load configuration
	configManager := config.NewManager()
	cfg, err := configManager.Load()
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	// Get command flags
	since, _ := cmd.Flags().GetString("since")
	count, _ := cmd.Flags().GetInt("count")
	output, _ := cmd.Flags().GetString("output")
	includeUnreleased, _ := cmd.Flags().GetBool("include-unreleased")

	// Initialize git repository
	gitRepo, err := git.NewRepository(".")
	if err != nil {
		return fmt.Errorf("failed to initialize git repository: %w", err)
	}

	// Get commits for changelog
	var commits []git.CommitInfo
	if since != "" {
		commits, err = gitRepo.GetCommitsSince(since)
		if err != nil {
			return fmt.Errorf("failed to get commits since %s: %w", since, err)
		}
	} else {
		commits, err = gitRepo.GetRecentCommits(count)
		if err != nil {
			return fmt.Errorf("failed to get recent commits: %w", err)
		}
	}

	if len(commits) == 0 {
		return fmt.Errorf("no commits found for changelog generation")
	}

	// Initialize AI client
	aiClient := ai.NewClient(cfg)

	// Generate changelog
	response, err := aiClient.GenerateChangelog(commits, includeUnreleased)
	if err != nil {
		return fmt.Errorf("failed to generate changelog: %w", err)
	}

	// Output changelog
	if output != "" {
		// Write to file
		err := writeToFile(output, response.Content)
		if err != nil {
			return fmt.Errorf("failed to write changelog to file: %w", err)
		}
		fmt.Printf("✅ Changelog written to %s\n", output)
	} else {
		// Print to stdout
		fmt.Println(response.Content)
	}

	return nil
}

func writeToFile(filename, content string) error {
	// TODO: Implement file writing
	return fmt.Errorf("file writing not yet implemented")
}
