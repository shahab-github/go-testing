package main

import (
	"testing"
)

func TestExtractSubdomain(t *testing.T) {
	cases := []struct {
		input          string
		expectedResult string
		expectError    bool
	}{
		{"svc-vault679@vault679.iam.gserviceaccount.com", "vault679", false},
		{"no-at-sign.com", "", true},
		{"no-dot-after-at@com", "", true},
		{"correct@format.com", "format", false},
	}

	for _, c := range cases {
		result, err := ExtractSubstring(c.input)
		if c.expectError && err == nil {
			t.Errorf("Expected an error for input %s, but got none", c.input)
		}
		if !c.expectError && err != nil {
			t.Errorf("Did not expect an error for input %s, but got %s", c.input, err)
		}
		if result != c.expectedResult {
			t.Errorf("Expected result %s for input %s, but got %s", c.expectedResult, c.input, result)
		}
	}
}


