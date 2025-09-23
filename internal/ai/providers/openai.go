package providers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/diffai/diffai/pkg/types"
)

// OpenAIProvider implements AI functionality using OpenAI API
type OpenAIProvider struct {
	apiKey  string
	model   string
	baseURL string
	client  *http.Client
}

// NewOpenAIProvider creates a new OpenAI provider
func NewOpenAIProvider(apiKey, model, baseURL string) *OpenAIProvider {
	return &OpenAIProvider{
		apiKey:  apiKey,
		model:   model,
		baseURL: baseURL,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// GenerateCommitMessage generates a commit message using OpenAI
func (p *OpenAIProvider) GenerateCommitMessage(diffs []types.Diff, options types.AIOptions) (*types.AIResponse, error) {
	// Create prompt for commit message generation
	prompt := p.buildCommitPrompt(diffs, options)

	// Call OpenAI API
	response, err := p.callOpenAI(prompt, options)
	if err != nil {
		return nil, fmt.Errorf("failed to call OpenAI: %w", err)
	}

	return &types.AIResponse{
		Success: true,
		Content: response.Choices[0].Message.Content,
		Usage: &types.Usage{
			Tokens:   response.Usage.TotalTokens,
			Model:    p.model,
			Provider: "openai",
		},
	}, nil
}

// GeneratePRSummary generates a PR summary using OpenAI
func (p *OpenAIProvider) GeneratePRSummary(prInfo interface{}, diffs []types.Diff, options types.AIOptions) (*types.AIResponse, error) {
	// Create prompt for PR summary generation
	prompt := p.buildPRPrompt(prInfo, diffs, options)

	// Call OpenAI API
	response, err := p.callOpenAI(prompt, options)
	if err != nil {
		return nil, fmt.Errorf("failed to call OpenAI: %w", err)
	}

	return &types.AIResponse{
		Success: true,
		Content: response.Choices[0].Message.Content,
		Usage: &types.Usage{
			Tokens:   response.Usage.TotalTokens,
			Model:    p.model,
			Provider: "openai",
		},
	}, nil
}

// GenerateChangelog generates a changelog using OpenAI
func (p *OpenAIProvider) GenerateChangelog(commits interface{}, options types.AIOptions) (*types.AIResponse, error) {
	// Create prompt for changelog generation
	prompt := p.buildChangelogPrompt(commits, options)

	// Call OpenAI API
	response, err := p.callOpenAI(prompt, options)
	if err != nil {
		return nil, fmt.Errorf("failed to call OpenAI: %w", err)
	}

	return &types.AIResponse{
		Success: true,
		Content: response.Choices[0].Message.Content,
		Usage: &types.Usage{
			Tokens:   response.Usage.TotalTokens,
			Model:    p.model,
			Provider: "openai",
		},
	}, nil
}

// callOpenAI makes a request to the OpenAI API
func (p *OpenAIProvider) callOpenAI(prompt string, options types.AIOptions) (*OpenAIResponse, error) {
	request := OpenAIRequest{
		Model: p.model,
		Messages: []OpenAIMessage{
			{
				Role:    "system",
				Content: "You are an expert software developer and Git specialist. Generate clear, concise, and professional commit messages, PR summaries, and changelogs.",
			},
			{
				Role:    "user",
				Content: prompt,
			},
		},
		MaxTokens:   500,
		Temperature: 0.7,
	}

	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", p.baseURL+"/chat/completions", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+p.apiKey)

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("OpenAI API error: status %d", resp.StatusCode)
	}

	var response OpenAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &response, nil
}

// buildCommitPrompt builds a prompt for commit message generation
func (p *OpenAIProvider) buildCommitPrompt(diffs []types.Diff, options types.AIOptions) string {
	// TODO: Implement prompt building
	return "Generate a commit message for the following changes:\n\n[DIFFS_PLACEHOLDER]"
}

// buildPRPrompt builds a prompt for PR summary generation
func (p *OpenAIProvider) buildPRPrompt(prInfo interface{}, diffs []types.Diff, options types.AIOptions) string {
	// TODO: Implement prompt building
	return "Generate a PR summary for the following changes:\n\n[PR_INFO_PLACEHOLDER]\n[DIFFS_PLACEHOLDER]"
}

// buildChangelogPrompt builds a prompt for changelog generation
func (p *OpenAIProvider) buildChangelogPrompt(commits interface{}, options types.AIOptions) string {
	// TODO: Implement prompt building
	return "Generate a changelog for the following commits:\n\n[COMMITS_PLACEHOLDER]"
}

// OpenAI API types
type OpenAIRequest struct {
	Model       string          `json:"model"`
	Messages    []OpenAIMessage `json:"messages"`
	MaxTokens   int             `json:"max_tokens"`
	Temperature float64         `json:"temperature"`
}

type OpenAIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIResponse struct {
	Choices []OpenAIChoice `json:"choices"`
	Usage   OpenAIUsage    `json:"usage"`
}

type OpenAIChoice struct {
	Message OpenAIMessage `json:"message"`
}

type OpenAIUsage struct {
	TotalTokens int `json:"total_tokens"`
}
