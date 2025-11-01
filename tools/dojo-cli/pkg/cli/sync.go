package cli

import (
	"fmt"

	"github.com/rafaeljc/code-dojo/tools/dojo-cli/internal/generator"
	"github.com/rafaeljc/code-dojo/tools/dojo-cli/internal/models"
	"github.com/rafaeljc/code-dojo/tools/dojo-cli/internal/parser"
	"github.com/spf13/cobra"
)

var syncTag string

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Synchronize tag indexes with problem metadata",
	Long: `Parse all problem README files and regenerate tag index files.
	
This command scans all problems, extracts their metadata, and updates
the tag index files in tags/ directory and the root README.md with the
current list of tags.`,
	RunE: runSync,
}

func init() {
	rootCmd.AddCommand(syncCmd)

	syncCmd.Flags().StringVarP(&syncTag, "tag", "t", "", "sync only this specific tag")
}

func runSync(cmd *cobra.Command, args []string) error {
	// Parse all problems
	fmt.Println("Parsing problems...")
	problems, err := parser.ParseAllProblems(repoRoot)
	if err != nil {
		return fmt.Errorf("failed to parse problems: %w", err)
	}

	if len(problems) == 0 {
		fmt.Println("⚠ No problems found")
		return nil
	}

	fmt.Printf("✓ Parsed %d problems\n", len(problems))

	// Generate tag indexes
	if syncTag != "" {
		// Filter problems by tag
		var filteredProblems []*models.Problem
		for _, p := range problems {
			for _, t := range p.Tags {
				if t == syncTag {
					filteredProblems = append(filteredProblems, p)
					break
				}
			}
		}

		if len(filteredProblems) == 0 {
			fmt.Printf("⚠ No problems found with tag '%s'\n", syncTag)
			return nil
		}

		if err := generator.GenerateTagIndexes(repoRoot, filteredProblems); err != nil {
			return fmt.Errorf("failed to generate tag index: %w", err)
		}

		fmt.Printf("✓ Updated tag: %s\n", syncTag)
	} else {
		// Generate all tags
		if err := generator.GenerateTagIndexes(repoRoot, problems); err != nil {
			return fmt.Errorf("failed to generate tag indexes: %w", err)
		}

		tags := generator.GetAllTags(problems)
		fmt.Printf("✓ Updated %d tag files\n", len(tags))

		// Update root README with tags
		fmt.Println("Updating root README...")
		if err := generator.UpdateRootREADME(repoRoot, tags); err != nil {
			return fmt.Errorf("failed to update root README: %w", err)
		}
		fmt.Println("✓ Updated root README")
	}

	return nil
}
