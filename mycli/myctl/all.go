package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

// KubeConfigTemplate holds the data for the kubeconfig template
type KubeConfigTemplate struct {
	APIServer string
	Token     string
}

func main() {
	// Define the API server URL (replace with your cluster's API server)
	apiServer := "https://your-k8s-api-server" // Replace with the actual Kubernetes API server

	// Step 1: Check if ~/.kube directory exists
	err := ensureKubeDir()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Step 2: Run the kubelogin command to get the token
	token, err := getKubeloginToken()
	if err != nil {
		fmt.Println("Error getting token:", err)
		return
	}

	// Step 3: Generate kubeconfig file and save it to ~/.kube/config
	err = createKubeConfigWithToken(token, apiServer)
	if err != nil {
		fmt.Println("Error creating kubeconfig:", err)
		return
	}

	fmt.Println("Kubeconfig file created successfully in ~/.kube/config")
}

// ensureKubeDir ensures that the ~/.kube directory exists
func ensureKubeDir() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("unable to find home directory: %w", err)
	}

	kubeDir := filepath.Join(homeDir, ".kube")

	// Check if ~/.kube directory exists
	if _, err := os.Stat(kubeDir); os.IsNotExist(err) {
		fmt.Println("~/.kube directory not found, creating it...")
		err := os.MkdirAll(kubeDir, 0755)
		if err != nil {
			return fmt.Errorf("failed to create ~/.kube directory: %w", err)
		}
	}
	return nil
}

// getKubeloginToken runs the kubelogin get-token command and returns the token
func getKubeloginToken() (string, error) {
	// Replace "your-parameters" with the actual parameters required for kubelogin
	cmd := exec.Command("kubelogin", "get-token", "--some-flags") // Replace with the actual flags for kubelogin
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to run kubelogin: %w", err)
	}

	// Convert the output (token) to string and return it
	token := string(output)
	return token, nil
}

// createKubeConfigWithToken generates the kubeconfig file with the token and saves it to ~/.kube/config
func createKubeConfigWithToken(token, apiServer string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("unable to find home directory: %w", err)
	}

	kubeConfigPath := filepath.Join(homeDir, ".kube", "config")

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

	// Open the file for writing
	file, err := os.Create(kubeConfigPath)
	if err != nil {
		return fmt.Errorf("failed to create ~/.kube/config file: %w", err)
	}
	defer file.Close()

	// Execute the template and write to the file
	err = tmpl.Execute(file, data)
	if err != nil {
		return fmt.Errorf("failed to execute kubeconfig template: %w", err)
	}

	return nil
}
