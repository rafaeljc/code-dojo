package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/rafaeljc/code-dojo/tools/dojo-cli/internal/models"
)

// GenerateTagIndexes creates or updates tag index files based on problems
func GenerateTagIndexes(repoRoot string, problems []*models.Problem) error {
	// Group problems by tag
	tagMap := make(map[string][]*models.Problem)
	for _, problem := range problems {
		for _, tag := range problem.Tags {
			tagMap[tag] = append(tagMap[tag], problem)
		}
	}

	// Create tags directory if it doesn't exist
	tagsDir := filepath.Join(repoRoot, "tags")
	if err := os.MkdirAll(tagsDir, 0755); err != nil {
		return fmt.Errorf("failed to create tags directory: %w", err)
	}

	// Load template from file
	templatePath := filepath.Join(repoRoot, "tools", "dojo-cli", "templates", "tag_index.tmpl")
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	for tagName, tagProblems := range tagMap {
		// Sort problems by ID
		sort.Slice(tagProblems, func(i, j int) bool {
			return tagProblems[i].ID < tagProblems[j].ID
		})

		// Create tag file
		tagFile := filepath.Join(tagsDir, strings.ToLower(tagName)+".md")
		f, err := os.Create(tagFile)
		if err != nil {
			return fmt.Errorf("failed to create tag file %s: %w", tagFile, err)
		}

		data := struct {
			TagName  string
			Problems []*models.Problem
		}{
			TagName:  formatTagName(tagName),
			Problems: tagProblems,
		}

		if err := tmpl.Execute(f, data); err != nil {
			f.Close()
			return fmt.Errorf("failed to write tag file %s: %w", tagFile, err)
		}
		f.Close()
	}

	return nil
}

// GetAllTags returns a sorted list of all unique tags from problems
func GetAllTags(problems []*models.Problem) []string {
	tagSet := make(map[string]struct{})
	for _, problem := range problems {
		for _, tag := range problem.Tags {
			tagSet[tag] = struct{}{}
		}
	}

	tags := make([]string, 0, len(tagSet))
	for tag := range tagSet {
		tags = append(tags, tag)
	}
	sort.Strings(tags)
	return tags
}

// formatTagName converts a tag slug to a display name
func formatTagName(tag string) string {
	words := strings.Split(tag, "-")
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(word[:1]) + word[1:]
		}
	}
	return strings.Join(words, " ")
}
