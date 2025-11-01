package cli

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/rafaeljc/code-dojo/tools/dojo-cli/internal/generator"
	"github.com/rafaeljc/code-dojo/tools/dojo-cli/internal/models"
	"github.com/spf13/cobra"
)

var (
	addTitle     string
	addTags      string
	addSource    string
	addExtension string
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new problem",
	Long:  `Add a new problem with automatic ID assignment, folder creation, and solution file.`,
	RunE:  runAdd,
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&addTitle, "title", "t", "", "problem title")
	addCmd.Flags().StringVarP(&addTags, "tags", "", "", "comma-separated tags")
	addCmd.Flags().StringVarP(&addSource, "source", "s", "", "problem source (e.g., LeetCode, HackerRank)")
	addCmd.Flags().StringVarP(&addExtension, "ext", "e", "", "solution file extension (e.g., py, cpp, js, go)")
}

func runAdd(cmd *cobra.Command, args []string) error {
	var err error

	// Get next problem ID
	nextID, err := generator.GetNextProblemID(repoRoot)
	if err != nil {
		return fmt.Errorf("failed to get next problem ID: %w", err)
	}

	// Interactive mode if flags not provided
	if addTitle == "" {
		prompt := promptui.Prompt{
			Label: "Problem title",
		}
		addTitle, err = prompt.Run()
		if err != nil {
			return err
		}
	}

	if addTags == "" {
		prompt := promptui.Prompt{
			Label: "Tags (comma-separated)",
		}
		addTags, err = prompt.Run()
		if err != nil {
			return err
		}
	}

	if addSource == "" {
		prompt := promptui.Prompt{
			Label:   "Source",
			Default: "LeetCode",
		}
		addSource, err = prompt.Run()
		if err != nil {
			return err
		}
	}

	// Ask for solution file extension
	if addExtension == "" {
		prompt := promptui.Prompt{
			Label:   "Solution file extension (e.g., py, cpp, js, go)",
			Default: "py",
		}
		addExtension, err = prompt.Run()
		if err != nil {
			return err
		}
	}

	// Parse tags
	tags := strings.Split(addTags, ",")
	for i, tag := range tags {
		tags[i] = strings.TrimSpace(tag)
	}

	// Create problem
	problem := &models.Problem{
		ID:     nextID,
		Title:  addTitle,
		Tags:   tags,
		Source: addSource,
	}

	if err := generator.GenerateProblem(repoRoot, problem, addExtension); err != nil {
		return fmt.Errorf("failed to generate problem: %w", err)
	}

	fmt.Printf("✓ Created problem %s in problems/%s/\n", problem.Title, problem.GetFolder())
	fmt.Printf("✓ Created solution.%s\n", addExtension)
	return nil
}
