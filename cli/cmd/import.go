package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/plate/cli/internal/project"
	"github.com/spf13/cobra"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import [path]",
	Short: "Import your project into the Plate platform",
	Long: `Import your project into the Plate platform for easy deployment.

Plate automatically detects your project type and sets up everything needed:
- Detects runtime (Node.js, Python, Go, Java, etc.)
- Configures build and start commands
- Sets up environment variables
- Prepares deployment configuration

After importing, you can deploy to any environment with a single command.

Examples:
  # Import current directory
  plate import
  
  # Import specific directory
  plate import /path/to/my-app
  
  # Import with custom name
  plate import --name my-awesome-app`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectPath := "."
		if len(args) > 0 {
			projectPath = args[0]
		}

		// Get absolute path
		absPath, err := filepath.Abs(projectPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting absolute path: %v\n", err)
			os.Exit(1)
		}

		// Get flags
		name, _ := cmd.Flags().GetString("name")
		env, _ := cmd.Flags().GetString("env")
		runtime, _ := cmd.Flags().GetString("runtime")

		// Import project
		importer := project.NewImporter()
		if err := importer.Import(absPath, name, env, runtime); err != nil {
			fmt.Fprintf(os.Stderr, "Error importing project: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Successfully imported project from %s\n", absPath)
	},
}

func init() {
	rootCmd.AddCommand(importCmd)

	importCmd.Flags().StringP("name", "n", "", "Project name (default: directory name)")
	importCmd.Flags().StringP("env", "e", "development", "Environment name")
	importCmd.Flags().StringP("runtime", "r", "", "Runtime type (auto-detect if not specified)")
}