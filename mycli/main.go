package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/confidential"
)

// get_token retrieves an access token using client credentials
func get_token() (string, error) {
	// Load sensitive data from environment variables
	clientSecret := os.Getenv("AZURE_CLIENT_SECRET")
	clientID := os.Getenv("AZURE_CLIENT_ID")
	tenantID := os.Getenv("AZURE_TENANT_ID")
	if clientSecret == "" || clientID == "" || tenantID == "" {
		return "", fmt.Errorf("Azure AD credentials not set in environment variables")
	}

	// Create confidential credential from client secret
	cred, err := confidential.NewCredFromSecret(clientSecret)
	if err != nil {
		return "", fmt.Errorf("failed to create credential: %w", err)
	}

	// Create confidential client with tenant ID
	authority := fmt.Sprintf("https://login.microsoftonline.com/%s", tenantID)
	client, err := confidential.New(authority, clientID, cred)
	if err != nil {
		return "", fmt.Errorf("failed to create confidential client: %w", err)
	}

	// Define the required scopes
	scopes := []string{"https://graph.microsoft.com/.default"}

	// Try to acquire token silently (from cache)
	result, err := client.AcquireTokenSilent(context.TODO(), scopes)
	if err != nil {
		// Cache miss, acquire token using client credentials
		result, err = client.AcquireTokenByCredential(context.TODO(), scopes)
		if err != nil {
			return "", fmt.Errorf("failed to acquire token: %w", err)
		}
	}

	// Return the acquired access token
	return result.AccessToken, nil
}

func main() {
	// Call get_token and handle error
	token, err := get_token()
	if err != nil {
		log.Fatalf("Error getting token: %v", err)
	}

	// Use the acquired access token
	fmt.Printf("Access Token: %s\n", token)
}
