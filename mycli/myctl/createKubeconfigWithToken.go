package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

// KubeConfigTemplate holds the data for the kubeconfig template
type KubeConfigTemplate struct {
	APIServer string
	Token     string
}

func main() {
	token := "your-token-here"                 // Replace with the actual token
	apiServer := "https://your-k8s-api-server" // Replace with the actual Kubernetes API server

	err := createKubeConfigWithInsecureVerify(token, apiServer)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Kubeconfig created successfully.")
}

// createKubeConfigWithInsecureVerify creates the ~/.kube/config file with insecure TLS verification
func createKubeConfigWithInsecureVerify(token, apiServer string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("unable to find home directory: %w", err)
	}

	kubeDir := filepath.Join(homeDir, ".kube")
	kubeConfigPath := filepath.Join(kubeDir, "config")

	// Ensure ~/.kube directory exists
	if _, err := os.Stat(kubeDir); os.IsNotExist(err) {
		fmt.Println("~/.kube directory not found, creating it...")
		if err := os.MkdirAll(kubeDir, 0755); err != nil {
			return fmt.Errorf("failed to create ~/.kube directory: %w", err)
		}
	}

	// Open the file for writing
	file, err := os.Create(kubeConfigPath)
	if err != nil {
		return fmt.Errorf("failed to create ~/.kube/config file: %w", err)
	}
	defer file.Close()

	// Define the kubeconfig template
	kubeConfigTemplate := `
apiVersion: v1
clusters:
- cluster:
    server: {{.APIServer}}
    insecure-skip-tls-verify: true
  name: my-cluster
contexts:
- context:
    cluster: my-cluster
    user: my-user
  name: my-context
current-context: my-context
kind: Config
preferences: {}
users:
- name: my-user
  user:
    token: {{.Token}}
`

	// Create a new template and parse the kubeconfig template into it
	tmpl, err := template.New("kubeconfig").Parse(kubeConfigTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse kubeconfig template: %w", err)
	}

	// Create the data structure for the template
	data := KubeConfigTemplate{
		APIServer: apiServer,
		Token:     token,
	}

	// Execute the template and write to the file
	err = tmpl.Execute(file, data)
	if err != nil {
		return fmt.Errorf("failed to execute kubeconfig template: %w", err)
	}

	return nil
}
