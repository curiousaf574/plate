package cmd

import (
	"fmt"
	"os"

	"github.com/plate/cli/internal/client"
	"github.com/spf13/cobra"
)

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy your application to an environment",
	Long: `Deploy your application to the specified environment.

This command will:
- Package your application code
- Deploy to the target environment (development, staging, or production)
- Provide real-time deployment feedback
- Generate a live URL for your application

Examples:
  # Deploy to development environment
  plate deploy
  
  # Deploy to production
  plate deploy --env production
  
  # Deploy specific version
  plate deploy --env staging --version v1.2.0`,
	Run: func(cmd *cobra.Command, args []string) {
		env, _ := cmd.Flags().GetString("env")
		watch, _ := cmd.Flags().GetBool("watch")

		client := client.NewAPIClient()
		
		fmt.Printf("Deploying to environment: %s\n", env)
		
		if err := client.Deploy(env); err != nil {
			fmt.Fprintf(os.Stderr, "Error deploying project: %v\n", err)
			os.Exit(1)
		}

		if watch {
			fmt.Println("Watching deployment status...")
			if err := client.WatchDeployment(env); err != nil {
				fmt.Fprintf(os.Stderr, "Error watching deployment: %v\n", err)
				os.Exit(1)
			}
		}

		fmt.Println("Deployment initiated successfully!")
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)

	deployCmd.Flags().StringP("env", "e", "development", "Environment to deploy to")
	deployCmd.Flags().BoolP("watch", "w", false, "Watch deployment progress")
}