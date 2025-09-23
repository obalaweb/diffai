package types

import "time"

// Diff represents a code change
type Diff struct {
	Path      string `json:"path"`
	OldContent string `json:"old_content,omitempty"`
	NewContent string `json:"new_content,omitempty"`
	ChangeType string `json:"change_type"` // added, modified, deleted, renamed
	Hunk      string `json:"hunk"`         // unified diff format
}

// CommitInfo represents commit metadata
type CommitInfo struct {
	Hash        string    `json:"hash"`
	Author      string    `json:"author"`
	Email       string    `json:"email"`
	Message     string    `json:"message"`
	Date        time.Time `json:"date"`
	Files       []string  `json:"files"`
	Insertions  int       `json:"insertions"`
	Deletions   int       `json:"deletions"`
}

// PRInfo represents pull request metadata
type PRInfo struct {
	Number      int       `json:"number"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Author      string    `json:"author"`
	BaseBranch  string    `json:"base_branch"`
	HeadBranch  string    `json:"head_branch"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Commits     []CommitInfo `json:"commits"`
	Files       []string  `json:"files"`
}

// AIRequest represents a request to the AI service
type AIRequest struct {
	Type        string      `json:"type"`        // commit, pr, changelog
	Diff        []Diff      `json:"diff"`
	Context     interface{} `json:"context"`     // CommitInfo, PRInfo, etc.
	Options     AIOptions   `json:"options"`
}

// AIOptions represents AI generation options
type AIOptions struct {
	Style       string `json:"style"`        // conventional, detailed, minimal
	Language    string `json:"language"`     // en, es, fr, etc.
	MaxLength   int    `json:"max_length"`
	IncludeRisk bool   `json:"include_risk"` // for PR summaries
}

// AIResponse represents a response from the AI service
type AIResponse struct {
	Success   bool        `json:"success"`
	Content   string      `json:"content"`
	Metadata  interface{} `json:"metadata,omitempty"`
	Error     string      `json:"error,omitempty"`
	Usage     *Usage      `json:"usage,omitempty"`
}

// Usage represents AI service usage statistics
type Usage struct {
	Tokens     int     `json:"tokens"`
	Cost       float64 `json:"cost"`
	Model      string  `json:"model"`
	Provider   string  `json:"provider"`
	Duration   int64   `json:"duration_ms"`
}

// Config represents application configuration
type Config struct {
	AI      AIConfig      `yaml:"ai"`
	Git     GitConfig     `yaml:"git"`
	Output  OutputConfig  `yaml:"output"`
	Service ServiceConfig `yaml:"service"`
}

// AIConfig represents AI service configuration
type AIConfig struct {
	Provider string            `yaml:"provider"` // openai, anthropic, local
	Model    string            `yaml:"model"`
	APIKey   string            `yaml:"api_key"`
	BaseURL  string            `yaml:"base_url"`
	Options  map[string]string `yaml:"options"`
}

// GitConfig represents Git-specific configuration
type GitConfig struct {
	ConventionalCommits bool   `yaml:"conventional_commits"`
	MaxCommitLength     int    `yaml:"max_commit_length"`
	AutoCommit          bool   `yaml:"auto_commit"`
	DefaultBranch       string `yaml:"default_branch"`
}

// OutputConfig represents output formatting configuration
type OutputConfig struct {
	Format  string `yaml:"format"`  // text, json, markdown
	Verbose bool   `yaml:"verbose"`
	Color   bool   `yaml:"color"`
}

// ServiceConfig represents AI service configuration
type ServiceConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Timeout  int    `yaml:"timeout_seconds"`
	Retries  int    `yaml:"retries"`
}
