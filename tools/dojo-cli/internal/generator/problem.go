package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/rafaeljc/code-dojo/tools/dojo-cli/internal/models"
)

// GenerateProblem creates a new problem folder with README.md and solution file
func GenerateProblem(repoRoot string, problem *models.Problem, extension string) error {
	// Create problem directory
	problemDir := filepath.Join(repoRoot, "problems", problem.GetFolder())
	if err := os.MkdirAll(problemDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Check if README already exists
	readmePath := filepath.Join(problemDir, "README.md")
	if _, err := os.Stat(readmePath); err == nil {
		return fmt.Errorf("problem already exists: %s", problemDir)
	}

	// Load template from file with custom functions
	templatePath := filepath.Join(repoRoot, "tools", "dojo-cli", "templates", "problem_readme.tmpl")
	tmpl, err := template.New("problem_readme.tmpl").Funcs(template.FuncMap{
		"lower": strings.ToLower,
	}).ParseFiles(templatePath)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	// Create README.md
	f, err := os.Create(readmePath)
	if err != nil {
		return fmt.Errorf("failed to create README.md: %w", err)
	}
	defer f.Close()

	// Prepare template data
	data := struct {
		ID        int
		Title     string
		Tags      string   // Comma-separated for YAML frontmatter
		TagsList  []string // Slice for linking
		Source    string
		Extension string // Solution file extension
	}{
		ID:        problem.ID,
		Title:     problem.Title,
		Tags:      strings.Join(problem.Tags, ", "),
		TagsList:  problem.Tags,
		Source:    problem.Source,
		Extension: extension,
	}

	if err := tmpl.Execute(f, data); err != nil {
		return fmt.Errorf("failed to write README.md: %w", err)
	}

	// Create solution file if extension provided
	if extension != "" {
		solutionPath := filepath.Join(problemDir, fmt.Sprintf("solution.%s", extension))
		solutionFile, err := os.Create(solutionPath)
		if err != nil {
			return fmt.Errorf("failed to create solution file: %w", err)
		}
		defer solutionFile.Close()

		// Write a basic comment based on the language
		comment := getCommentForExtension(extension)
		if comment != "" {
			githubURL := fmt.Sprintf("https://github.com/rafaeljc/code-dojo/tree/main/problems/%s", problem.GetFolder())
			fmt.Fprintf(solutionFile, "%s %s\n%s Source: %s\n%s Problem: %s\n\n",
				comment, problem.Title,
				comment, problem.Source,
				comment, githubURL)
		}
	}

	return nil
}

// getCommentForExtension returns the comment syntax for common languages
func getCommentForExtension(ext string) string {
	comments := map[string]string{
		"py":    "#",
		"js":    "//",
		"ts":    "//",
		"java":  "//",
		"cpp":   "//",
		"c":     "//",
		"go":    "//",
		"rs":    "//",
		"rb":    "#",
		"sh":    "#",
		"php":   "//",
		"swift": "//",
		"kt":    "//",
		"scala": "//",
	}

	if comment, ok := comments[ext]; ok {
		return comment
	}
	return "#" // default
}

// GetNextProblemID scans existing problems and returns the next available ID
func GetNextProblemID(repoRoot string) (int, error) {
	problemsDir := filepath.Join(repoRoot, "problems")

	entries, err := os.ReadDir(problemsDir)
	if err != nil {
		if os.IsNotExist(err) {
			return 1, nil
		}
		return 0, fmt.Errorf("failed to read problems directory: %w", err)
	}

	maxID := 0
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		// Parse ID from folder name (format: 0001-problem-name)
		name := entry.Name()
		if len(name) < 4 {
			continue
		}

		var id int
		_, err := fmt.Sscanf(name[:4], "%d", &id)
		if err == nil && id > maxID {
			maxID = id
		}
	}

	return maxID + 1, nil
}
