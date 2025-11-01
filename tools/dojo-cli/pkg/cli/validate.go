package cli

import (
	"fmt"

	"github.com/rafaeljc/code-dojo/tools/dojo-cli/internal/parser"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate problem metadata",
	Long:  `Check all problem README files for valid and complete metadata.`,
	RunE:  runValidate,
}

func init() {
	rootCmd.AddCommand(validateCmd)
}

func runValidate(cmd *cobra.Command, args []string) error {
	// Parse all problems
	problems, err := parser.ParseAllProblems(repoRoot)
	if err != nil {
		return fmt.Errorf("failed to parse problems: %w", err)
	}

	if len(problems) == 0 {
		fmt.Println("⚠ No problems found")
		return nil
	}

	// Validate each problem
	hasErrors := false
	for _, problem := range problems {
		errors := parser.ValidateProblem(problem)
		if len(errors) > 0 {
			hasErrors = true
			fmt.Printf("⚠ Problem %s:\n", problem.Folder)
			for _, err := range errors {
				fmt.Printf("  - %s\n", err)
			}
		}
	}

	if !hasErrors {
		fmt.Printf("✓ All %d problems are valid\n", len(problems))
	}

	return nil
}
