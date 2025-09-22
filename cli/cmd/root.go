package cmd
// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "plate",
	Short: "Developer portal CLI for seamless app deployment",
	Long: `Plate is a developer-friendly CLI tool that simplifies application deployment.

Focus on your code while Plate handles the deployment complexity:
- Import your projects with automatic runtime detection
- Deploy to multiple environments with a single command
- Monitor your applications across development, staging, and production
- Get instant feedback on deployment status and application health

No infrastructure knowledge required - just build, deploy, and ship!`,mt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "plate",
	Short: "Plate CLI - Deploy your code to Kubernetes with ease",
	Long: `Plate is a CLI tool that allows you to import your projects 
and deploy them to Kubernetes clusters using GitOps workflows.

Features:
- Import existing projects
- Deploy to Kubernetes
- GitOps integration with ArgoCD
- Helm chart management
- Environment management`,
	Version: "1.0.0",
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	// Global flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.plate.yaml)")
	rootCmd.PersistentFlags().String("api-url", "http://localhost:8080", "Plate API server URL")
	rootCmd.PersistentFlags().String("token", "", "Authentication token for Plate API")

	// Bind flags to viper
	viper.BindPFlag("api-url", rootCmd.PersistentFlags().Lookup("api-url"))
	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".plate")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}