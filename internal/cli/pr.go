package cli

import (
	"fmt"
	"strconv"

	"github.com/diffai/diffai/internal/ai"
	"github.com/diffai/diffai/internal/config"
	"github.com/diffai/diffai/internal/git"
	"github.com/spf13/cobra"
)

// NewPRCommand creates the PR command
func NewPRCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pr <number>",
		Short: "Generate AI-powered PR summary",
		Long: `Generate a comprehensive summary of a pull request using AI.

This command analyzes the PR diff and generates a developer-friendly summary
that highlights what changed, why it changed, and the potential impact.

Examples:
  git diffai pr 123                    # Summarize PR #123
  git diffai pr 123 --include-risk     # Include risk assessment
  git diffai pr 123 --format markdown  # Output in markdown format`,
		Args: cobra.ExactArgs(1),
		RunE: runPR,
	}

	// Add flags
	cmd.Flags().BoolP("include-risk", "r", false, "Include risk assessment in summary")
	cmd.Flags().StringP("format", "f", "text", "Output format (text, markdown, json)")
	cmd.Flags().BoolP("verbose", "v", false, "Verbose output with additional details")

	return cmd
}

func runPR(cmd *cobra.Command, args []string) error {
	// Parse PR number
	prNumber, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid PR number: %s", args[0])
	}

	// Load configuration
	configManager := config.NewManager()
	cfg, err := configManager.Load()
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	// Get command flags
	includeRisk, _ := cmd.Flags().GetBool("include-risk")
	format, _ := cmd.Flags().GetString("format")
	verbose, _ := cmd.Flags().GetBool("verbose")

	// Initialize git repository
	gitRepo, err := git.NewRepository(".")
	if err != nil {
		return fmt.Errorf("failed to initialize git repository: %w", err)
	}

	// Get PR information
	prInfo, err := gitRepo.GetPRInfo(prNumber)
	if err != nil {
		return fmt.Errorf("failed to get PR information: %w", err)
	}

	// Get PR diff
	diff, err := gitRepo.GetPRDiff(prNumber)
	if err != nil {
		return fmt.Errorf("failed to get PR diff: %w", err)
	}

	// Initialize AI client
	aiClient := ai.NewClient(cfg)

	// Generate PR summary
	response, err := aiClient.GeneratePRSummary(prInfo, diff, includeRisk)
	if err != nil {
		return fmt.Errorf("failed to generate PR summary: %w", err)
	}

	// Output summary based on format
	switch format {
	case "json":
		// TODO: Implement JSON output
		fmt.Println("JSON output not yet implemented")
	case "markdown":
		// TODO: Implement markdown output
		fmt.Println("Markdown output not yet implemented")
	default:
		// Text output
		fmt.Printf("PR #%d Summary\n", prNumber)
		fmt.Println("=" + fmt.Sprintf("%*s", len(fmt.Sprintf("PR #%d Summary", prNumber)), "="))
		fmt.Println()
		fmt.Println(response.Content)
		
		if verbose && response.Usage != nil {
			fmt.Println()
			fmt.Printf("AI Usage: %d tokens, %.4f cost, %dms\n", 
				response.Usage.Tokens, 
				response.Usage.Cost, 
				response.Usage.Duration)
		}
	}

	return nil
}
