package ghapi

import (
	"testing"
	"time"
)

func TestNewGitHubClient(t *testing.T) {
	// Test with minimal configuration (no token for CI)
	client, err := NewGitHubClient("", "", "test-version")
	if err != nil {
		t.Fatalf("Failed to create GitHub client: %v", err)
	}

	if client == nil {
		t.Fatal("Expected non-nil GitHub client")
	}

	// Verify the client is properly configured
	if client.BaseURL == nil {
		t.Fatal("Expected BaseURL to be set")
	}

	// Test that the client uses the same base structure as go-github
	if client.BaseURL.String() != "https://api.github.com/" {
		t.Errorf("Expected BaseURL to be https://api.github.com/, got %s", client.BaseURL.String())
	}
}

func TestHTTPClientCreation(t *testing.T) {
	opts := HTTPClientOptions{
		AppVersion:  "test-1.0.0",
		Token:       "test-token",
		EnableCache: true,
		CacheTTL:    time.Hour,
	}

	client, err := NewHTTPClient(opts)
	if err != nil {
		t.Fatalf("Failed to create HTTP client: %v", err)
	}

	if client == nil {
		t.Fatal("Expected non-nil HTTP client")
	}

	// Verify that the transport is properly configured
	if client.Transport == nil {
		t.Fatal("Expected Transport to be set")
	}
}

func TestGitHubClientWithToken(t *testing.T) {
	// Test with token (simulates real usage)
	client, err := NewGitHubClient("", "test-token", "test-version")
	if err != nil {
		t.Fatalf("Failed to create GitHub client with token: %v", err)
	}

	if client == nil {
		t.Fatal("Expected non-nil GitHub client")
	}

	// Verify that we can access basic client properties
	if client.BaseURL == nil {
		t.Fatal("Expected BaseURL to be set")
	}
}

func TestEnterpriseGitHubClient(t *testing.T) {
	// Test with enterprise hostname
	client, err := NewGitHubClient("enterprise.github.com", "test-token", "test-version")
	if err != nil {
		t.Fatalf("Failed to create GitHub client for enterprise: %v", err)
	}

	if client == nil {
		t.Fatal("Expected non-nil GitHub client")
	}

	// For enterprise, base URL should be different
	expectedURL := "https://enterprise.github.com/api/v3/"
	if client.BaseURL.String() != expectedURL {
		t.Errorf("Expected BaseURL to be %s, got %s", expectedURL, client.BaseURL.String())
	}
}