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

## Usage

### Building the MCP Server

```bash
go build -o mcpserver ./cmd/mcpserver
```

### Running the MCP Server

```bash
# Run with default settings (port 3000, path /mcp, baseURL http://localhost:8080)
./mcpserver

# Run with custom settings
./mcpserver --port 8000 --baseurl https://1scan.dev
```

### Configuration Parameters

- `--port`: Port to run the MCP server on (default: "3000")
- `--baseurl`: Base URL of the 1scan API (default: "http://localhost:8080")

## Integration with Cursor IDE

The MCP server can be used with Cursor IDE. To use it:

1. Start the MCP server using the instructions above
2. Configure Cursor to use the MCP server at `http://localhost:3000/mcp` (or your custom port)

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