package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

const templateRepo = "https://github.com/sudo-which-qp/my_flutter_template_v2.git"

func createFlutterApp(appName, githubToken, packageName string) error {
	fmt.Printf("Cloning template into %s...\n", appName)

	// Clone the repository with authentication
	if err := cloneRepository(appName, githubToken); err != nil {
		return fmt.Errorf("failed to clone repository: %v", err)
	}

	fmt.Println("Template cloned successfully!")

	// Update package names
	fmt.Println("Updating package names...")
	if err := updatePackageName(appName, packageName); err != nil {
		return fmt.Errorf("failed to update package names: %v", err)
	}

	fmt.Println("Package names updated successfully!")

	// Run flutter pub get
	fmt.Println("Running flutter pub get...")
	if err := runFlutterPubGet(appName); err != nil {
		return fmt.Errorf("failed to run flutter pub get: %v", err)
	}

	fmt.Println("Setup complete! Your app is ready.")
	return nil
}

func cloneRepository(appName, githubToken string) error {
	cloneOptions := &git.CloneOptions{
		URL: templateRepo,
		Auth: &http.BasicAuth{
			Username: "x-access-token",
			Password: githubToken,
		},
	}

	_, err := git.PlainClone(appName, false, cloneOptions)
	return err
}

func updatePackageName(appDir, packageName string) error {
	// Update pubspec.yaml
	pubspecPath := filepath.Join(appDir, "pubspec.yaml")
	content, err := os.ReadFile(pubspecPath)
	if err != nil {
		return err
	}

	newContent := strings.Replace(string(content), "my_flutter_template_v2", appDir, -1)
	if err := os.WriteFile(pubspecPath, []byte(newContent), 0644); err != nil {
		return err
	}

	// Update dart files in lib directory
	libDir := filepath.Join(appDir, "lib")
	return filepath.Walk(libDir, func(path string, info os.FileInfo, err error) error {
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

func runFlutterPubGet(appDir string) error {
	flutterCmd := exec.Command("flutter", "pub", "get")
	flutterCmd.Dir = appDir
	flutterCmd.Stdout = os.Stdout
	flutterCmd.Stderr = os.Stderr
	return flutterCmd.Run()
}
