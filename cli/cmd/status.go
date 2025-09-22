package cmd

import (
	"fmt"
	"os"

	"github.com/plate/cli/internal/client"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check your application deployment status",
	Long: `Check the status of your deployed applications across all environments.

This command shows:
- Application health and status
- Live URLs for each environment
- Deployment history and versions
- Build and runtime information

Use this to monitor your applications and troubleshoot any issues.

Examples:
  # Check status of all applications
  plate status
  
  # Get detailed status information
  plate status --detailed
  
  # Check specific environment
  plate status --env production`,
	Run: func(cmd *cobra.Command, args []string) {
		env, _ := cmd.Flags().GetString("env")
		detailed, _ := cmd.Flags().GetBool("detailed")

		client := client.NewAPIClient()
		
		status, err := client.GetStatus(env, detailed)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting status: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(status)
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	statusCmd.Flags().StringP("env", "e", "", "Environment to check (default: all)")
	statusCmd.Flags().BoolP("detailed", "d", false, "Show detailed status information")
}