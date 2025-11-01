package models

import "fmt"

// Problem represents a coding problem with its metadata
type Problem struct {
	ID     int      `yaml:"id"`
	Title  string   `yaml:"title"`
	Tags   []string `yaml:"tags"`
	Source string   `yaml:"source"`
	Folder string   `yaml:"-"` // Computed, not in YAML
}

// GetFolder returns the formatted folder name for the problem
func (p *Problem) GetFolder() string {
	if p.Folder != "" {
		return p.Folder
	}
	return fmt.Sprintf("%04d-%s", p.ID, slugify(p.Title))
}

// slugify converts a title to a URL-friendly slug
func slugify(s string) string {
	var result []rune
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			result = append(result, toLower(r))
		} else if r == ' ' || r == '-' {
			if len(result) > 0 && result[len(result)-1] != '-' {
				result = append(result, '-')
			}
		}
	}
	// Trim trailing dash
	if len(result) > 0 && result[len(result)-1] == '-' {
		result = result[:len(result)-1]
	}
	return string(result)
}

func toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + ('a' - 'A')
	}
	return r
}
