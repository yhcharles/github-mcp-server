# GitHub CLI Integration

This document describes the modifications made to github-mcp-server to use the same GitHub API access patterns as the GitHub CLI (gh).

## Overview

The github-mcp-server has been modified to use the same authentication and HTTP client infrastructure as the GitHub CLI (`gh`). This ensures consistent behavior, enhanced security, and better error handling when interacting with the GitHub API.

## Key Changes

### 1. New API Client Package (`pkg/ghapi/`)

A new package has been created that mirrors the gh CLI's API client functionality:

- **`client.go`**: Core API client with GraphQL and REST functionality
- **`http_client.go`**: HTTP client configuration with authentication and caching
- **`integration_test.go`**: Tests to verify the new functionality

### 2. Dependencies Added

- `github.com/cli/go-gh/v2`: Official GitHub CLI library for authentication and API access

### 3. Modified Files

#### `internal/ghmcp/server.go`
- Updated to use the new `ghapi` package
- Replaced simple token authentication with gh CLI's sophisticated HTTP client
- Added proper caching and user agent handling
- Removed custom `bearerAuthTransport` in favor of gh CLI's authentication

#### `go.mod`
- Added `github.com/cli/go-gh/v2` dependency

## Benefits of This Approach

### 1. **Consistent Authentication**
- Uses the same token resolution logic as gh CLI
- Supports environment variables, config files, and system keyring
- Proper hostname normalization and enterprise support

### 2. **Enhanced HTTP Client**
- Built-in caching with configurable TTL
- Request/response logging capabilities
- Proper user agent handling
- Round-trip interceptors for custom headers

### 3. **Better Error Handling**
- OAuth scope suggestions for permission errors
- Detailed HTTP error information
- GraphQL error handling with partial data population

### 4. **Enterprise Support**
- Proper hostname parsing for GitHub Enterprise Server (GHES)
- GitHub Enterprise Cloud (GHEC) support
- Automatic URL construction for different GitHub instances

## API Compatibility

The modified server maintains full backward compatibility with the existing MCP interface. All existing tools and functionality continue to work unchanged.

## Technical Details

### Authentication Flow
1. HTTP client created with gh CLI's `ghAPI.NewHTTPClient()`
2. Authentication token handled by gh CLI's transport layer
3. Automatic header injection for Authorization and User-Agent
4. Support for redirect handling and hostname validation

### GraphQL Client
- Uses `shurcooL/githubv4` with gh CLI's HTTP client
- Maintains enterprise client configuration
- Proper authentication token handling

### REST Client
- Uses `google/go-github/v74` with gh CLI's HTTP client
- Automatic base URL configuration for different GitHub instances
- Upload URL handling for GitHub releases and artifacts

## Usage

The server configuration remains the same. Users provide:
- `Host`: GitHub hostname (github.com, enterprise.github.com, etc.)
- `Token`: GitHub personal access token or GitHub App token
- `Version`: Server version for user agent strings

The new implementation automatically handles:
- Proper API endpoint discovery
- Authentication header formatting
- Caching and performance optimization
- Error handling and scope suggestions

## Testing

The integration includes comprehensive tests in `pkg/ghapi/integration_test.go` that verify:
- HTTP client creation with various configurations
- GitHub client initialization for different hostnames
- Enterprise GitHub support
- Token handling and authentication setup

Run tests with:
```bash
go test ./pkg/ghapi/ -v
```

## Backwards Compatibility

All existing functionality remains unchanged:
- Same MCP server interface
- Same tool definitions and behaviors
- Same configuration options
- Same REST and GraphQL API access patterns

The changes are internal implementation improvements that enhance reliability and consistency with the GitHub CLI ecosystem.