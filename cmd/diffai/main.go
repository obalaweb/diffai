package main

import (
	"os"

	"github.com/diffai/diffai/internal/cli"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "diffai",
		Short: "AI-powered Git assistant that understands your diffs",
		Long: `DiffAI helps developers make better commits, generate meaningful PR summaries, 
				and automate changelogs — all powered by AI that actually understands code changes.

				DiffAI integrates seamlessly with Git CLI and provides intelligent assistance for:
				- Generating conventional commit messages from staged diffs
				- Summarizing pull requests into developer-friendly overviews  
				- Auto-generating changelogs grouped by features & fixes
				- Extensible integrations for GitHub, GitLab, Bitbucket

				Examples:
				git diffai commit          # Generate commit message from staged changes
				git diffai pr 123          # Summarize PR #123
				git diffai changelog       # Generate changelog from recent commits
				git diffai config          # Manage configuration`,
		Version: "0.1.0",
	}

	// Add subcommands
	rootCmd.AddCommand(cli.NewCommitCommand())
	rootCmd.AddCommand(cli.NewPRCommand())
	rootCmd.AddCommand(cli.NewChangelogCommand())
	rootCmd.AddCommand(cli.NewConfigCommand())
	rootCmd.AddCommand(cli.NewVersionCommand())

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
