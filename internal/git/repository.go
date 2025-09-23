package git

import (
	"fmt"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// Repository represents a Git repository
type Repository struct {
	repo *git.Repository
	path string
}

// NewRepository creates a new repository instance
func NewRepository(path string) (*Repository, error) {
	// Resolve absolute path
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve path: %w", err)
	}

	// Open repository
	repo, err := git.PlainOpen(absPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open repository: %w", err)
	}

	return &Repository{
		repo: repo,
		path: absPath,
	}, nil
}

// HasStagedChanges checks if there are staged changes
func (r *Repository) HasStagedChanges() (bool, error) {
	worktree, err := r.repo.Worktree()
	if err != nil {
		return false, fmt.Errorf("failed to get worktree: %w", err)
	}

	status, err := worktree.Status()
	if err != nil {
		return false, fmt.Errorf("failed to get status: %w", err)
	}

	// Check if there are any staged changes
	for _, fileStatus := range status {
		if fileStatus.Staging != git.Unmodified {
			return true, nil
		}
	}

	return false, nil
}

// GetStagedDiff returns the diff of staged changes
func (r *Repository) GetStagedDiff() ([]Diff, error) {
	worktree, err := r.repo.Worktree()
	if err != nil {
		return nil, fmt.Errorf("failed to get worktree: %w", err)
	}

	status, err := worktree.Status()
	if err != nil {
		return nil, fmt.Errorf("failed to get status: %w", err)
	}

	var diffs []Diff

	// Get HEAD commit for comparison
	head, err := r.repo.Head()
	if err != nil {
		return nil, fmt.Errorf("failed to get HEAD: %w", err)
	}

	headCommit, err := r.repo.CommitObject(head.Hash())
	if err != nil {
		return nil, fmt.Errorf("failed to get HEAD commit: %w", err)
	}

	// Process each staged file
	for filePath, fileStatus := range status {
		if fileStatus.Staging == git.Unmodified {
			continue
		}

		diff, err := r.getFileDiff(filePath, fileStatus.Staging, headCommit)
		if err != nil {
			return nil, fmt.Errorf("failed to get diff for %s: %w", filePath, err)
		}

		diffs = append(diffs, *diff)
	}

	return diffs, nil
}

// CreateCommit creates a commit with the given message
func (r *Repository) CreateCommit(message string) (string, error) {
	worktree, err := r.repo.Worktree()
	if err != nil {
		return "", fmt.Errorf("failed to get worktree: %w", err)
	}

	// Create commit
	commitHash, err := worktree.Commit(message, &git.CommitOptions{})
	if err != nil {
		return "", fmt.Errorf("failed to create commit: %w", err)
	}

	return commitHash.String(), nil
}

// GetRecentCommits returns recent commits
func (r *Repository) GetRecentCommits(count int) ([]CommitInfo, error) {
	head, err := r.repo.Head()
	if err != nil {
		return nil, fmt.Errorf("failed to get HEAD: %w", err)
	}

	commitIter, err := r.repo.Log(&git.LogOptions{
		From: head.Hash(),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get log: %w", err)
	}
	defer commitIter.Close()

	var commits []CommitInfo
	for i := 0; i < count; i++ {
		commit, err := commitIter.Next()
		if err != nil {
			break
		}

		commitInfo := r.convertCommit(commit)
		commits = append(commits, commitInfo)
	}

	return commits, nil
}

// GetCommitsSince returns commits since a specific tag or commit
func (r *Repository) GetCommitsSince(since string) ([]CommitInfo, error) {
	// TODO: Implement commits since functionality
	return nil, fmt.Errorf("commits since functionality not yet implemented")
}

// GetPRInfo returns information about a pull request
func (r *Repository) GetPRInfo(prNumber int) (*PRInfo, error) {
	// TODO: Implement PR info retrieval
	return nil, fmt.Errorf("PR info retrieval not yet implemented")
}

// GetPRDiff returns the diff for a pull request
func (r *Repository) GetPRDiff(prNumber int) ([]Diff, error) {
	// TODO: Implement PR diff retrieval
	return nil, fmt.Errorf("PR diff retrieval not yet implemented")
}

// getFileDiff gets the diff for a specific file
func (r *Repository) getFileDiff(filePath string, fileStatus git.StatusCode, headCommit *object.Commit) (*Diff, error) {
	// TODO: Implement file diff extraction
	return &Diff{
		Path:       filePath,
		ChangeType: "modified", // Placeholder
		Hunk:       "",         // Placeholder
	}, nil
}

// convertCommit converts a git commit to our CommitInfo struct
func (r *Repository) convertCommit(commit *object.Commit) CommitInfo {
	// Get file statistics
	stats, _ := commit.Stats()

	var files []string
	var insertions, deletions int

	for _, stat := range stats {
		files = append(files, stat.Name)
		insertions += stat.Addition
		deletions += stat.Deletion
	}

	return CommitInfo{
		Hash:       commit.Hash.String(),
		Author:     commit.Author.Name,
		Email:      commit.Author.Email,
		Message:    commit.Message,
		Date:       commit.Author.When.Format("2006-01-02 15:04:05"),
		Files:      files,
		Insertions: insertions,
		Deletions:  deletions,
	}
}
