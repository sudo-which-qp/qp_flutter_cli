package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	packageName string
	token       string
)

var rootCmd = &cobra.Command{
	Use:   "qp_flutter",
	Short: "QP Flutter CLI tool",
}

var createCmd = &cobra.Command{
	Use:   "create [app_name]",
	Short: "Create a new Flutter app from template",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		appName := args[0]

		// Get GitHub token from environment
		githubToken := os.Getenv("GITHUB_TOKEN")
		if githubToken == "" {
			return fmt.Errorf("GITHUB_TOKEN environment variable is required")
		}

		// Create the app
		return createFlutterApp(appName, githubToken, packageName)
	},
}

func init() {
	createCmd.Flags().StringVarP(&packageName, "package", "p", "", "Package name (e.g. com.company.app)")
	createCmd.Flags().StringVarP(&token, "token", "t", "", "GitHub personal access token")
	createCmd.MarkFlagRequired("package")

	rootCmd.AddCommand(createCmd)
}
