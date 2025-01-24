package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

var (
	templateRepo = "https://github.com/sudo-which-qp/my_flutter_template_v2.git"
	packageName  string
)

func main() {
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

			fmt.Printf("Cloning template into %s...\n", appName)
			_, err := git.PlainClone(appName, false, &git.CloneOptions{
				URL: templateRepo,
			})
			if err != nil {
				return err
			}
			fmt.Println("Template cloned successfully!")

			fmt.Println("Updating package names...")
			if err := updatePackageName(appName); err != nil {
				return err
			}
			fmt.Println("Package names updated successfully!")

			fmt.Println("Running flutter pub get...")
			flutterCmd := exec.Command("flutter", "pub", "get")
			flutterCmd.Dir = appName
			if err := flutterCmd.Run(); err != nil {
				return err
			}
			fmt.Println("Setup complete! Your app is ready.")
			return nil
		},
	}

	createCmd.Flags().StringVarP(&packageName, "package", "p", "", "Package name (e.g. com.company.app)")
	createCmd.MarkFlagRequired("package")

	rootCmd.AddCommand(createCmd)
	rootCmd.Execute()
}

func updatePackageName(appDir string) error {
	pubspecPath := filepath.Join(appDir, "pubspec.yaml")
	content, err := os.ReadFile(pubspecPath)
	if err != nil {
		return err
	}
	newContent := strings.Replace(string(content), "my_flutter_template_v2", appDir, -1)
	if err := os.WriteFile(pubspecPath, []byte(newContent), 0644); err != nil {
		return err
	}

	return filepath.Walk(filepath.Join(appDir, "lib"), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".dart") {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			newContent := strings.Replace(string(content),
				"package:my_flutter_template_v2",
				"package:"+appDir, -1)
			if err := os.WriteFile(path, []byte(newContent), 0644); err != nil {
				return err
			}
		}
		return nil
	})
}
