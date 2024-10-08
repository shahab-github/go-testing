package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	err := ensureKubeConfig()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("~/.kube/config check completed successfully.")
}

// ensureKubeConfig ensures that the ~/.kube/config file exists
func ensureKubeConfig() error {
	// Get the home directory of the user
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("unable to find home directory: %w", err)
	}

	// Define the paths
	kubeDir := filepath.Join(homeDir, ".kube")
	kubeConfigPath := filepath.Join(kubeDir, "config")

	// Ensure ~/.kube directory exists
	if _, err := os.Stat(kubeDir); os.IsNotExist(err) {
		fmt.Println("~/.kube directory not found, creating it...")
		if err := os.MkdirAll(kubeDir, 0755); err != nil {
			return fmt.Errorf("failed to create ~/.kube directory: %w", err)
		}
	}

	// Check if ~/.kube/config exists
	if _, err := os.Stat(kubeConfigPath); os.IsNotExist(err) {
		fmt.Println("~/.kube/config file not found, creating it...")

		// Create the config file
		file, err := os.Create(kubeConfigPath)
		if err != nil {
			return fmt.Errorf("failed to create ~/.kube/config file: %w", err)
		}
		defer file.Close()

		// Write default content to the config file (optional)
		defaultConfig := "# Kubernetes config file\n"
		if _, err = file.WriteString(defaultConfig); err != nil {
			return fmt.Errorf("failed to write default content to config file: %w", err)
		}

		fmt.Println("~/.kube/config file created successfully.")
	} else {
		fmt.Println("~/.kube/config file already exists.")
	}

	return nil
}
