package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/rafaeljc/code-dojo/tools/dojo-cli/internal/models"
	"gopkg.in/yaml.v3"
)

var frontmatterRegex = regexp.MustCompile(`^---\s*\n([\s\S]*?)\n---\s*\n`)

// ParseProblem reads a problem README and extracts metadata from YAML frontmatter
func ParseProblem(readmePath string) (*models.Problem, error) {
	content, err := os.ReadFile(readmePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	matches := frontmatterRegex.FindSubmatch(content)
	if matches == nil {
		return nil, fmt.Errorf("no frontmatter found in %s", readmePath)
	}

	var problem models.Problem
	if err := yaml.Unmarshal(matches[1], &problem); err != nil {
		return nil, fmt.Errorf("failed to parse frontmatter: %w", err)
	}

	// Extract folder name from path
	dir := filepath.Dir(readmePath)
	problem.Folder = filepath.Base(dir)

	return &problem, nil
}

// ParseAllProblems scans the problems directory and parses all problem READMEs
func ParseAllProblems(repoRoot string) ([]*models.Problem, error) {
	problemsDir := filepath.Join(repoRoot, "problems")

	entries, err := os.ReadDir(problemsDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read problems directory: %w", err)
	}

	var problems []*models.Problem
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		readmePath := filepath.Join(problemsDir, entry.Name(), "README.md")
		if _, err := os.Stat(readmePath); os.IsNotExist(err) {
			continue
		}

		problem, err := ParseProblem(readmePath)
		if err != nil {
			// Log warning but continue with other problems
			fmt.Fprintf(os.Stderr, "Warning: %v\n", err)
			continue
		}

		problems = append(problems, problem)
	}

	return problems, nil
}

// ValidateProblem checks if a problem has all required fields
func ValidateProblem(problem *models.Problem) []string {
	var errors []string

	if problem.ID == 0 {
		errors = append(errors, "missing 'id' field")
	}
	if problem.Title == "" {
		errors = append(errors, "missing 'title' field")
	}
	if len(problem.Tags) == 0 {
		errors = append(errors, "missing 'tags' field")
	}
	if problem.Source == "" {
		errors = append(errors, "missing 'source' field")
	}

	return errors
}
