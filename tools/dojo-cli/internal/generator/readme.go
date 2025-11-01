package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// UpdateRootREADME updates the tags section in the root README.md
func UpdateRootREADME(repoRoot string, tags []string) error {
	readmePath := filepath.Join(repoRoot, "README.md")

	// Read the current README
	content, err := os.ReadFile(readmePath)
	if err != nil {
		return fmt.Errorf("failed to read README.md: %w", err)
	}

	// Find the tags section
	lines := strings.Split(string(content), "\n")
	tagSectionStart := -1
	tagSectionEnd := -1

	for i, line := range lines {
		if strings.Contains(line, "## 🏷️ Tags") {
			tagSectionStart = i
		} else if tagSectionStart != -1 && strings.HasPrefix(line, "## ") {
			tagSectionEnd = i
			break
		}
	}

	if tagSectionStart == -1 {
		return fmt.Errorf("could not find '## 🏷️ Tags' section in README.md")
	}

	// If no next section found, go to end of file
	if tagSectionEnd == -1 {
		tagSectionEnd = len(lines)
	}

	// Build new tags section
	var newTagsSection []string
	newTagsSection = append(newTagsSection, "## 🏷️ Tags")
	newTagsSection = append(newTagsSection, "")
	newTagsSection = append(newTagsSection, "Problems are organized by topic using tag indexes in `/tags/`.")

	for _, tag := range tags {
		displayName := formatTagName(tag)
		newTagsSection = append(newTagsSection, fmt.Sprintf("- [%s](tags/%s.md)", displayName, tag))
	}
	newTagsSection = append(newTagsSection, "")

	// Reconstruct the file
	var newLines []string
	newLines = append(newLines, lines[:tagSectionStart]...)
	newLines = append(newLines, newTagsSection...)
	newLines = append(newLines, lines[tagSectionEnd:]...)

	// Write back to file
	newContent := strings.Join(newLines, "\n")
	if err := os.WriteFile(readmePath, []byte(newContent), 0644); err != nil {
		return fmt.Errorf("failed to write README.md: %w", err)
	}

	return nil
}
