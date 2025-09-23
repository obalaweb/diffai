package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/diffai/diffai/internal/git"
	"github.com/diffai/diffai/pkg/types"
)

// Client represents an AI service client
type Client struct {
	config     *types.Config
	httpClient *http.Client
	baseURL    string
}

// NewClient creates a new AI client
func NewClient(config *types.Config) *Client {
	return &Client{
		config: config,
		httpClient: &http.Client{
			Timeout: time.Duration(config.Service.Timeout) * time.Second,
		},
		baseURL: fmt.Sprintf("http://%s:%d", config.Service.Host, config.Service.Port),
	}
}

// GenerateCommitMessage generates a commit message from staged diff
func (c *Client) GenerateCommitMessage(diffs []git.Diff, style string) (*types.AIResponse, error) {
	request := types.AIRequest{
		Type: "commit",
		Diff: convertDiffs(diffs),
		Options: types.AIOptions{
			Style:     style,
			Language:  "en",
			MaxLength: c.config.Git.MaxCommitLength,
		},
	}

	return c.makeRequest("/api/v1/commit", request)
}

// GeneratePRSummary generates a PR summary
func (c *Client) GeneratePRSummary(prInfo *git.PRInfo, diffs []git.Diff, includeRisk bool) (*types.AIResponse, error) {
	request := types.AIRequest{
		Type:    "pr",
		Diff:    convertDiffs(diffs),
		Context: prInfo,
		Options: types.AIOptions{
			Style:       "detailed",
			Language:    "en",
			IncludeRisk: includeRisk,
		},
	}

	return c.makeRequest("/api/v1/pr", request)
}

// GenerateChangelog generates a changelog from commits
func (c *Client) GenerateChangelog(commits []git.CommitInfo, includeUnreleased bool) (*types.AIResponse, error) {
	request := types.AIRequest{
		Type:    "changelog",
		Context: commits,
		Options: types.AIOptions{
			Style:  "markdown",
			Language: "en",
		},
	}

	return c.makeRequest("/api/v1/changelog", request)
}

// makeRequest makes an HTTP request to the AI service
func (c *Client) makeRequest(endpoint string, request types.AIRequest) (*types.AIResponse, error) {
	// Marshal request
	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create HTTP request
	url := c.baseURL + endpoint
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "DiffAI-CLI/0.1.0")

	// Add API key if configured
	if c.config.AI.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.config.AI.APIKey)
	}

	// Make request with retries
	var response *types.AIResponse
	for i := 0; i <= c.config.Service.Retries; i++ {
		resp, err := c.httpClient.Do(req)
		if err != nil {
			if i == c.config.Service.Retries {
				return nil, fmt.Errorf("failed to make request after %d retries: %w", c.config.Service.Retries, err)
			}
			time.Sleep(time.Duration(i+1) * time.Second)
			continue
		}
		defer resp.Body.Close()

		// Parse response
		response = &types.AIResponse{}
		if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("AI service error: %s", response.Error)
		}

		break
	}

	return response, nil
}

// convertDiffs converts git.Diff to types.Diff
func convertDiffs(gitDiffs []git.Diff) []types.Diff {
	var diffs []types.Diff
	for _, gitDiff := range gitDiffs {
		diffs = append(diffs, types.Diff{
			Path:       gitDiff.Path,
			OldContent: gitDiff.OldContent,
			NewContent: gitDiff.NewContent,
			ChangeType: gitDiff.ChangeType,
			Hunk:       gitDiff.Hunk,
		})
	}
	return diffs
}
