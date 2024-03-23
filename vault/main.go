package main

import (
	"errors"
	"fmt"
	"strings"
)

// ExtractSubstring extracts the substring from an email-like string between '@' and the first '.'.
// Returns an error if the input string does not contain these characters in the expected order.
func ExtractSubstring(input string) (string, error) {
	if !strings.Contains(input, "@") || !strings.Contains(input, ".") {
		return "", errors.New("input string does not contain the required characters")
	}

	parts := strings.Split(input, "@")
	if len(parts) < 2 || len(parts[1]) == 0 {
		return "", errors.New("invalid format after '@'")
	}

	subParts := strings.Split(parts[1], ".")
	if len(subParts) == 0 {
		return "", errors.New("invalid format before '.'")
	}

	return subParts[0], nil
}

func main() {
	input := "svc-vault679@vault679.iam.gserviceaccount.com"
	result, err := ExtractSubstring(input)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Extracted substring:", result)
	}
}