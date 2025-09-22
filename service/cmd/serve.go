package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/plate/service/internal/api"
	"github.com/plate/service/internal/config"
	"github.com/plate/service/internal/services"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the Plate API server",
	Long: `Start the Plate API server to handle deployment requests.
This will start the HTTP server and initialize all required services.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Load()
		
		// Initialize services
		serviceManager := services.NewManager(cfg, nil)
		if err := serviceManager.Initialize(); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to initialize services: %v\n", err)
			os.Exit(1)
		}

		// Initialize API server
		server := api.NewServer(cfg, serviceManager)
		
		// Start server
		srv := &http.Server{
			Addr:    ":" + cfg.Port,
			Handler: server.Router(),
		}

		// Graceful shutdown
		go func() {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				fmt.Fprintf(os.Stderr, "Failed to start server: %v\n", err)
				os.Exit(1)
			}
		}()

		fmt.Printf("Plate service started on port %s\n", cfg.Port)

		// Wait for interrupt signal to gracefully shutdown
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		fmt.Println("Shutting down server...")

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			fmt.Fprintf(os.Stderr, "Server forced to shutdown: %v\n", err)
		}

		fmt.Println("Server exited")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}