package git

import (
	"fmt"
	"strings"
)

// Diff represents a code change
type Diff struct {
	Path       string `json:"path"`
	OldContent string `json:"old_content,omitempty"`
	NewContent string `json:"new_content,omitempty"`
	ChangeType string `json:"change_type"` // added, modified, deleted, renamed
	Hunk       string `json:"hunk"`        // unified diff format
}

// CommitInfo represents commit metadata
type CommitInfo struct {
	Hash       string   `json:"hash"`
	Author     string   `json:"author"`
	Email      string   `json:"email"`
	Message    string   `json:"message"`
	Date       string   `json:"date"`
	Files      []string `json:"files"`
	Insertions int      `json:"insertions"`
	Deletions  int      `json:"deletions"`
}

// PRInfo represents pull request metadata
type PRInfo struct {
	Number      int          `json:"number"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Author      string       `json:"author"`
	BaseBranch  string       `json:"base_branch"`
	HeadBranch  string       `json:"head_branch"`
	CreatedAt   string       `json:"created_at"`
	UpdatedAt   string       `json:"updated_at"`
	Commits     []CommitInfo `json:"commits"`
	Files       []string     `json:"files"`
}

// ParseUnifiedDiff parses a unified diff string
func ParseUnifiedDiff(diffStr string) ([]Diff, error) {
	// TODO: Implement unified diff parsing
	// This is a placeholder implementation
	return []Diff{}, fmt.Errorf("unified diff parsing not yet implemented")
}

// FormatDiff formats a diff for display
func FormatDiff(diff Diff) string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("File: %s\n", diff.Path))
	builder.WriteString(fmt.Sprintf("Type: %s\n", diff.ChangeType))
	
	if diff.Hunk != "" {
		builder.WriteString("Hunk:\n")
		builder.WriteString(diff.Hunk)
	}

	return builder.String()
}

// GetChangeType determines the change type from git status
func GetChangeType(status string) string {
	switch status {
	case "A":
		return "added"
	case "M":
		return "modified"
	case "D":
		return "deleted"
	case "R":
		return "renamed"
	case "C":
		return "copied"
	default:
		return "unknown"
	}
}

// SummarizeDiffs provides a summary of multiple diffs
func SummarizeDiffs(diffs []Diff) string {
	if len(diffs) == 0 {
		return "No changes"
	}

	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("Summary: %d files changed\n", len(diffs)))

	// Count by change type
	counts := make(map[string]int)
	for _, diff := range diffs {
		counts[diff.ChangeType]++
	}

	for changeType, count := range counts {
		builder.WriteString(fmt.Sprintf("  %s: %d files\n", changeType, count))
	}

	return builder.String()
}
