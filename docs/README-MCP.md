# 1scan MCP Server

This is a Model Context Protocol (MCP) server for the 1scan project that allows AI models to query blockchain data through the 1scan API.

## Overview

The 1scan MCP server exposes a set of tools that AI models can use to query data from multiple blockchain networks through a unified API. It acts as a bridge between AI assistants and blockchain data.

## Available Tools

- `getAccountBalance` - Get the balance of an account on a specific blockchain
- `getTokenBalance` - Get the token balance of an account on a specific blockchain
- `getTransactionByHash` - Get transaction details by hash
- `getBlockByNumber` - Get block information by block number
- `getContractABI` - Get the ABI for a verified contract

## Prerequisites

- Go 1.21 or higher
- Running 1scan API server (or other compatible blockchain API)

## Installation

### Building the MCP Server

```bash
go build -o mcpserver ./cmd/mcpserver
```

### Running the MCP Server

```bash
# Run with default settings (port 3000, path /mcp, baseURL http://localhost:8080)
./mcpserver

# Run with custom settings
./mcpserver --port 8000 --path /mcp --baseurl https://1scan.dev --transport sse
```

### Configuration Parameters

- `--port`: Port to run the MCP server on (default: "3000")
- `--path`: Path for the MCP server endpoint (default: "/mcp")
- `--baseurl`: Base URL of the 1scan API (default: "http://localhost:8080")
- `--transport`: Transport type (sse or stdio) (default: "sse")

## Integration with Cursor IDE

The MCP server can be used with Cursor IDE. To use it:

1. Start the MCP server using the instructions above
2. Configure Cursor to use the MCP server at `http://localhost:3000/mcp` (or your custom port/path)

## Example Queries

AI models can query blockchain data through the MCP server using the following tools:

```
// Get an account balance
{
  "chainID": "1",
  "address": "0x742d35Cc6634C0532925a3b844Bc454e4438f44e"
}

// Get a token balance
{
  "chainID": "1",
  "contractAddress": "0xdac17f958d2ee523a2206206994597c13d831ec7",
  "address": "0x742d35Cc6634C0532925a3b844Bc454e4438f44e"
}

// Get transaction details
{
  "chainID": "1",
  "txHash": "0xb2fea9c4b24775af6990237aa90228e5e092c56bdaee7c3054e5471fa3b5d38b"
}

// Get block information
{
  "chainID": "1",
  "blockNumber": "latest"
}

// Get contract ABI
{
  "chainID": "1",
  "contractAddress": "0xdac17f958d2ee523a2206206994597c13d831ec7"
}
```

## Troubleshooting

### Connection Errors

If you see "SSE error: connect ECONNREFUSED" errors:

1. Make sure your 1scan API server is running at the specified baseURL
2. You can run the MCP server with mock data by modifying the `makeRequest` function in `internal/mcp/tools.go` to return mock responses instead of real API calls

### 404 Errors

If you see "SSE error: 404" errors:

1. Ensure that the server paths are correctly configured
2. Check the server logs to see the actual endpoints being used
3. Make sure Cursor is pointing to the exact URL shown in the logs
4. Try restarting both the server and Cursor

### Server Logs

The server logs will show:
- The URL and path configuration
- The SSE and message endpoints
- Each request made to the blockchain API
- Any errors encountered

Example:
```
Setting up SSE server with base URL: http://localhost:3000 and path: /mcp
Starting MCP server with SSE transport on http://localhost:3000/mcp
SSE endpoint: http://localhost:3000/mcp/sse
Message endpoint: http://localhost:3000/mcp/message
MCP server connecting to 1scan API at http://localhost:8080
```

## Architecture

The MCP server consists of:

1. **BlockchainTool** - Handles API requests to the 1scan service
2. **Tool Handlers** - Process tool calls from AI models and return formatted responses
3. **SSE Server** - Handles the Server-Sent Events protocol for real-time communication
4. **MCP Server** - Core MCP implementation that manages tools and requests

## How It Works

1. AI model connects to the MCP server via SSE
2. Model requests available tools
3. Model makes a tool call with parameters
4. MCP server forwards request to 1scan API
5. Response is formatted and returned to the model

## License

MIT License 