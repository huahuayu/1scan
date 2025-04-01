# 1scan - Blockchain Explorer API Gateway & MCP Server

A unified API gateway for integrating multiple etherscan-like blockchain explorer APIs with Model Context Protocol (MCP) support for AI assistants.

## Overview

1scan provides two main components:

1. **API Gateway**: A unified endpoint for accessing multiple blockchain explorer APIs
2. **MCP Server**: Allows AI models (like Claude in Cursor IDE) to directly query blockchain data

## Supported Blockchain Networks

| Network Name        | Chain ID | Explorer Link                   | Explorer API Endpoint                |
| ------------------- | -------- | ------------------------------- | ------------------------------------ |
| Ethereum            | 1        | https://etherscan.io            | api.etherscan.io                     |
| Binance Smart Chain | 56       | https://bscscan.com             | api.bscscan.com                      |
| Arbitrum One        | 42161    | https://arbiscan.io             | api.arbiscan.io                      |
| Polygon             | 137      | https://polygonscan.com         | api.polygonscan.com                  |
| Optimism            | 10       | https://optimistic.etherscan.io | api-optimistic.etherscan.io          |
| Avalanche C-Chain   | 43114    | https://snowtrace.io            | api.snowtrace.io                     |
| Base                | 8453     | https://basescan.org            | api.basescan.org                     |
| zkSync Era          | 324      | https://explorer.zksync.io      | block-explorer-api.mainnet.zksync.io |
| Gnosis              | 100      | https://gnosisscan.io           | api.gnosisscan.io                    |
| Fantom              | 250      | https://ftmscan.com             | api.ftmscan.com                      |
| Mantle              | 5000     | https://mantlescan.info         | api.mantlescan.info                  |
| Cronos              | 25       | https://cronoscan.com           | api.cronoscan.com                    |
| Polygon ZkEVM       | 1101     | https://zkevm.polygonscan.com   | api-zkevm.polygonscan.com            |
| Linea               | 59144    | https://lineascan.build         | api.lineascan.build                  |
| Moonbeam            | 1284     | https://moonscan.io             | api.moonscan.io                      |
| Celo                | 42220    | https://celoscan.io             | api.celoscan.io                      |
| Scroll              | 534352   | https://scrollscan.com          | api.scrollscan.com                   |
| OpBNB               | 204      | https://opbnbscan.com           | api.opbnbscan.com                    |
| Moonriver           | 1285     | https://moonriver.moonscan.io   | api-moonriver.moonscan.io            |
| Arbitrum Nova       | 42170    | https://nova.arbiscan.io        | api-nova.arbiscan.io                 |
| Blast               | 81457    | https://blastscan.io            | api.blastscan.io                     |
| Fraxtal             | 252      | https://fraxscan.com            | api.fraxscan.com                     |
| Wemix               | 1111     | https://wemixscan.com           | api.wemixscan.com                    |
| Xai                 | 660279   | https://xaiscan.io              | api.xaiscan.io                       |
| World Chain         | 480      | https://worldscan.org           | api.worldscan.org                    |
| Ape                 | 33139    | https://apescan.io/             | api.apescan.io                       |
| Kroma               | 255      | https://kromascan.com           | api.kromascan.com                    |
| Taiko               | 167000   | https://taikoscan.io            | api.taikoscan.io                     |
| Bittorrent          | 199      | https://bttcscan.com            | api.bttcscan.com                     |
| Xdc                 | 50       | https://xdcscan.io              | api.xdcscan.io                       |

## Features

- 🔄 Unified API endpoint for multiple blockchain explorers
- ⚖️ API load balancing
- 🔑 API rate limit management
- 🎯 Custom API key support via URL parameters
- 🤖 MCP server for AI assistants to query blockchain data

## Installation

### From Go

```bash
go install github.com/huahuayu/1scan@latest
```

### From Source

```bash
# Clone the repository
git clone https://github.com/huahuayu/1scan.git
cd 1scan

# Build both 1scan API server and MCP server
make build
```

## Quick Start

### 1. Create Configuration File

Create a `config.json` file with your API keys:

```json
{
  "1": {
    "endpoint": "api.etherscan.io",
    "keys": {
      "YOUR_ETHERSCAN_API_KEY_1": 5,
      "YOUR_ETHERSCAN_API_KEY_2": 10
    }
  },
  "56": {
    "endpoint": "api.bscscan.com",
    "keys": {
      "YOUR_BSCSCAN_API_KEY": 5
    }
  }
}
```

Each chain entry contains:

- `chainID`: The blockchain network ID
- `endpoint`: The API endpoint for the blockchain explorer
- `keys`: Map of API keys and their rate limits (requests per second)

### 2. Run the API Server

```bash
# Run with default settings
1scan -config /path/to/config.json

# Or using make
make run-1scan
```

### 3. Run the MCP Server

```bash
# Run with default settings
1scanmcp -config /path/to/config.json # default serve at port 3000

# Or using make
make run-1scanmcp

# With custom settings
1scanmcp -config /path/to/config.json -port 3000 -path /mcp
```

## Integrating with Cursor IDE

To use 1scan MCP in Cursor IDE:

1. Start both the API server and MCP server

   ```bash
   # Terminal 1
   make run-1scan

   # Terminal 2
   make run-1scanmcp
   ```

2. In Cursor IDE, go to Settings → AI → MCP Servers

3. Add a new MCP server with the URL:

```json
{
  "mcpServers": {
    "1scan": {
      "url": "http://localhost:3000/mcp/sse"
    }
  }
}
```

4. Restart Cursor IDE if necessary

5. Now you can use blockchain data in your AI conversations:
   ```
   Can you check the balance of address 0x742d35Cc6634C0532925a3b844Bc454e4438f44e on Ethereum?
   ```

## MCP Server Configuration

When running the MCP server, you can customize it with these parameters:

- `-port`: Port to run the MCP server on (default: "3000")
- `-path`: Path for the MCP server endpoint (default: "/mcp")
- `-config`: Path to configuration file (default: "config.json")
- `-transport`: Transport type (sse or stdio) (default: "sse")

Example:

```bash
1scanmcp -port 3000 -config /path/to/config.json
```

## Available MCP Tools

The MCP server exposes these tools to AI models:

- `getAccountBalance` - Get the balance of an account on a specific blockchain
- `getTokenBalance` - Get the token balance of an account on a specific blockchain
- `getTransactionByHash` - Get transaction details by hash
- `getBlockByNumber` - Get block information by block number
- `getContractABI` - Get the ABI for a verified contract
- `getContractSourceCode` - Get the source code of a verified contract
- `getTokenInfo` - Get information about an ERC20 token
- `getGasPrice` - Get current gas price on a specific blockchain
- `getLogs` - Get event logs matching specified parameters
- `getTransaction` - Get details of a transaction by its hash
- `getTransactionCount` - Get the number of transactions sent from an address
- `getTransactionReceipt` - Get transaction receipt by transaction hash
- `getBlockTransactionCountByNumber` - Get the number of transactions in a block
- `getBlockTransactionCountByHash` - Get the number of transactions in a block by block hash
- `getBlockByHash` - Get information about a block by its hash
- `getNormalTransactions` - Get normal transactions by address
- `getInternalTransactions` - Get internal transactions by address or transaction hash
- `getERC20Transfers` - Get ERC-20 token transfer events by address or contract
- `getERC721Transfers` - Get ERC-721 NFT transfer events by address or contract
- `getValidators` - Get validator information (for networks with validator sets)
- `getContractCreation` - Get contract creation information
- `getContractVerificationStatus` - Check contract verification status

## Example MCP Interactions

Here are examples of how to interact with blockchain data through the MCP interface:

```
# Get account balance
What's the ETH balance of vitalik.eth (0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045) on Ethereum?

# Check a token balance
What's the USDC balance of address 0x28C6c06298d514Db089934071355E5743bf21d60 on Ethereum?

# Examine a transaction
Can you analyze transaction 0xdd6d7f687c9821404ae8c2ea7de5cfb5a23fc4c01ef1e7f535748cb147aa76dd on Ethereum?

# Look up block information
What transactions were in Ethereum block 18700000?

# Get contract ABI
Can you show me the ABI for the USDC contract (0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48) on Ethereum?

# Analyze contract source code
I need to understand how the Uniswap V3 router works. Can you show me the source code for 0xE592427A0AEce92De3Edee1F18E0157C05861564 on Ethereum?

# Check multiple chains
What's the difference in gas fees between Ethereum and Arbitrum right now?

# Examine token information
Tell me about the tokenomics of the APE token on Ethereum.

## Complete List of MCP Tools

The MCP server exposes these tools to AI models:

- `getAccountBalance` - Get the balance of an account on a specific blockchain
- `getTokenBalance` - Get the token balance of an account on a specific blockchain
- `getTransactionByHash` - Get transaction details by hash
- `getBlockByNumber` - Get block information by block number
- `getContractABI` - Get the ABI for a verified contract
- `getContractSourceCode` - Get the source code of a verified contract
- `getTokenInfo` - Get information about an ERC20 token
- `getGasPrice` - Get current gas price on a specific blockchain
- `getLogs` - Get event logs matching specified parameters
- `getTransaction` - Get details of a transaction by its hash
- `getTransactionCount` - Get the number of transactions sent from an address
- `getTransactionReceipt` - Get transaction receipt by transaction hash
- `getBlockTransactionCountByNumber` - Get the number of transactions in a block
- `getBlockTransactionCountByHash` - Get the number of transactions in a block by block hash
- `getBlockByHash` - Get information about a block by its hash
- `getNormalTransactions` - Get normal transactions by address
- `getInternalTransactions` - Get internal transactions by address or transaction hash
- `getERC20Transfers` - Get ERC-20 token transfer events by address or contract
- `getERC721Transfers` - Get ERC-721 NFT transfer events by address or contract
- `getValidators` - Get validator information (for networks with validator sets)
- `getContractCreation` - Get contract creation information
- `getContractVerificationStatus` - Check contract verification status

## Troubleshooting MCP Integration

If you encounter issues connecting to the MCP server from Cursor:

1. Verify both servers are running:

   ```bash
   # Check if servers are running
   ps aux | grep 1scanmcp
   ```

2. Ensure the correct URL is configured in Cursor

## License

MIT License
