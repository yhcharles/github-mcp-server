# GitHub Authentication Guide

This document explains how to authenticate with GitHub using the updated github-mcp-server that supports both manual tokens and gh CLI automatic authentication.

## Recommended: Use GitHub CLI Authentication

The preferred method is to use GitHub CLI's built-in authentication, which eliminates the need to manage tokens manually.

### Setup Steps

1. **Install GitHub CLI** (if not already installed):
   ```bash
   # macOS
   brew install gh
   
   # Ubuntu/Debian
   sudo apt install gh
   
   # Windows
   winget install GitHub.cli
   ```

2. **Authenticate with GitHub**:
   ```bash
   gh auth login
   ```
   
   Follow the interactive prompts to:
   - Choose GitHub.com or GitHub Enterprise Server
   - Select your preferred authentication method (web browser or token)
   - Complete the authentication flow

3. **Verify authentication**:
   ```bash
   gh auth status
   ```

4. **Configure github-mcp-server**:
   
   When configuring the MCP server, simply omit the `token` parameter or set it to an empty string:
   
   ```json
   {
     "mcpServers": {
       "github": {
         "command": "github-mcp-server",
         "args": ["--host", "github.com"],
         "env": {
           "GITHUB_HOST": "github.com"
         }
       }
     }
   }
   ```

   **No need to set `GITHUB_TOKEN` environment variable!**

### Benefits of gh CLI Authentication

- **Secure**: Uses the same secure token storage as GitHub CLI
- **Convenient**: No need to create or manage personal access tokens
- **Multi-account**: Supports multiple GitHub accounts
- **Enterprise**: Works with GitHub Enterprise Server and GitHub Enterprise Cloud
- **Automatic refresh**: Handles token refresh automatically

## Alternative: Manual Token Authentication

If you prefer to use a manual token or gh CLI is not available, you can still use the traditional token-based authentication.

### Setup Steps

1. **Create a Personal Access Token**:
   - Go to GitHub Settings → Developer settings → Personal access tokens
   - Generate a new token with required scopes
   - Copy the token

2. **Configure the server with token**:
   ```json
   {
     "mcpServers": {
       "github": {
         "command": "github-mcp-server",
         "env": {
           "GITHUB_TOKEN": "your-personal-access-token",
           "GITHUB_HOST": "github.com"
         }
       }
     }
   }
   ```

## Authentication Priority

The server follows this authentication priority:

1. **Explicit token**: If a token is provided via `GITHUB_TOKEN` environment variable, it will be used
2. **gh CLI authentication**: If no token is provided, the server will attempt to use gh CLI's authentication
3. **Error**: If neither is available, authentication will fail

## Enterprise GitHub Support

### GitHub Enterprise Server (GHES)

```bash
# Authenticate with your enterprise server
gh auth login --hostname your-enterprise.com

# Configure the server
# Set GITHUB_HOST to your enterprise hostname
```

### GitHub Enterprise Cloud (GHEC)

```bash
# Authenticate with GHEC
gh auth login --hostname ghe.your-company.com

# Configure the server  
# Set GITHUB_HOST to your GHEC hostname
```

## Troubleshooting

### Authentication Issues

1. **Check gh CLI status**:
   ```bash
   gh auth status
   ```

2. **Re-authenticate if needed**:
   ```bash
   gh auth refresh
   ```

3. **Check token scopes** (for manual tokens):
   Ensure your token has the required scopes for the operations you want to perform.

### Common Error Messages

- **"No token found"**: Either set up gh CLI authentication or provide a manual token
- **"Authentication failed"**: Your token may be expired or have insufficient permissions
- **"Host not found"**: Check your `GITHUB_HOST` configuration

### Debug Authentication

You can check which authentication method is being used by looking at the server logs. The server will indicate whether it's using:
- Manual token authentication
- gh CLI automatic authentication

## Migration from Manual Tokens

If you're currently using manual tokens and want to switch to gh CLI authentication:

1. Set up gh CLI authentication as described above
2. Remove the `GITHUB_TOKEN` environment variable from your configuration
3. Restart the github-mcp-server

The server will automatically detect and use gh CLI authentication.

## Security Best Practices

1. **Use gh CLI authentication** when possible for better security
2. **Rotate tokens regularly** if using manual tokens
3. **Use minimal scopes** - only grant permissions your application needs
4. **Store tokens securely** - never commit tokens to version control
5. **Use organization-scoped tokens** when working with organization repositories

## Multiple GitHub Accounts

gh CLI supports multiple GitHub accounts:

```bash
# Add additional accounts
gh auth login --hostname github.com --account work
gh auth login --hostname github.com --account personal

# Switch between accounts
gh auth switch --hostname github.com --user work
```

The github-mcp-server will use the currently active account for the specified hostname.