package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var repoRoot string

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "dojo",
	Short: "Code Dojo - Manage your coding practice repository",
	Long: `Code Dojo CLI helps you manage your coding practice repository.

Features:
  - Create new problems with automatic ID assignment
  - Parse problem metadata and generate tag indexes
  - Validate problem metadata
  - Update root README with tag links`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Find repo root by looking for problems/ directory
		if repoRoot == "" {
			cwd, err := os.Getwd()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			repoRoot = findRepoRoot(cwd)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&repoRoot, "root", "r", "", "repository root directory (auto-detected if not specified)")
}

// findRepoRoot walks up the directory tree to find the repo root
func findRepoRoot(start string) string {
	dir := start
	for {
		// Check if problems/ directory exists
		if _, err := os.Stat(filepath.Join(dir, "problems")); err == nil {
			return dir
		}

		// Move up one directory
		parent := filepath.Dir(dir)
		if parent == dir {
			// Reached filesystem root
			fmt.Fprintf(os.Stderr, "Error: Could not find repository root (no problems/ directory found)\n")
			os.Exit(1)
		}
		dir = parent
	}
}
