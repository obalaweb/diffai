package cli

import (
	"fmt"

	"github.com/diffai/diffai/internal/ai"
	"github.com/diffai/diffai/internal/config"
	"github.com/diffai/diffai/internal/git"
	"github.com/spf13/cobra"
)

// NewCommitCommand creates the commit command
func NewCommitCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "commit",
		Short: "Generate AI-powered commit message from staged changes",
		Long: `Generate a conventional commit message from your staged changes using AI.

This command analyzes your staged diff and generates a meaningful commit message
that follows conventional commit standards.

Examples:
  git diffai commit                    # Generate commit message
  git diffai commit --auto             # Auto-commit with generated message
  git diffai commit --style detailed   # Use detailed commit style`,
		RunE: runCommit,
	}

	// Add flags
	cmd.Flags().BoolP("auto", "a", false, "Automatically commit with generated message")
	cmd.Flags().StringP("style", "s", "conventional", "Commit message style (conventional, detailed, minimal)")
	cmd.Flags().BoolP("dry-run", "d", false, "Show generated message without committing")
	cmd.Flags().StringP("message", "m", "", "Override AI-generated message")

	return cmd
}

func runCommit(cmd *cobra.Command, args []string) error {
	// Load configuration
	configManager := config.NewManager()
	cfg, err := configManager.Load()
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	// Get command flags
	autoCommit, _ := cmd.Flags().GetBool("auto")
	dryRun, _ := cmd.Flags().GetBool("dry-run")
	style, _ := cmd.Flags().GetString("style")
	overrideMessage, _ := cmd.Flags().GetString("message")

	// Check if there are staged changes
	gitRepo, err := git.NewRepository(".")
	if err != nil {
		return fmt.Errorf("failed to initialize git repository: %w", err)
	}

	hasStaged, err := gitRepo.HasStagedChanges()
	if err != nil {
		return fmt.Errorf("failed to check staged changes: %w", err)
	}

	if !hasStaged {
		return fmt.Errorf("no staged changes found. Please stage your changes with 'git add' first")
	}

	// Get staged diff
	diff, err := gitRepo.GetStagedDiff()
	if err != nil {
		return fmt.Errorf("failed to get staged diff: %w", err)
	}

	// Generate commit message
	var commitMessage string
	if overrideMessage != "" {
		commitMessage = overrideMessage
	} else {
		// Initialize AI client
		aiClient := ai.NewClient(cfg)
		
		// Generate message using AI
		response, err := aiClient.GenerateCommitMessage(diff, style)
		if err != nil {
			return fmt.Errorf("failed to generate commit message: %w", err)
		}

		commitMessage = response.Content
	}

	// Display generated message
	fmt.Printf("Generated commit message:\n\n%s\n\n", commitMessage)

	if dryRun {
		fmt.Println("Dry run mode - no commit created")
		return nil
	}

	// Confirm with user unless auto-commit is enabled
	if !autoCommit {
		fmt.Print("Commit with this message? [y/N]: ")
		var response string
		fmt.Scanln(&response)
		
		if response != "y" && response != "Y" && response != "yes" {
			fmt.Println("Commit cancelled")
			return nil
		}
	}

	// Create commit
	commitHash, err := gitRepo.CreateCommit(commitMessage)
	if err != nil {
		return fmt.Errorf("failed to create commit: %w", err)
	}

	fmt.Printf("✅ Commit created successfully: %s\n", commitHash)
	return nil
}
