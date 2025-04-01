# Etherscan API Integration for MCP

This document explains how the 1scan MCP server integrates with Etherscan APIs to provide blockchain data.

## Configuration

The MCP server uses a `config.json` file to configure the Etherscan API endpoints and API keys. The configuration file has the following structure:

```json
{
  "1": {
    "endpoint": "api.etherscan.io",
    "keys": {
      "YOUR_ETHERSCAN_API_KEY_1": 5,
      "YOUR_ETHERSCAN_API_KEY_2": 5
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

Each blockchain is identified by its chain ID:
- `1`: Ethereum Mainnet
- `56`: Binance Smart Chain (BSC)
- etc.

For each chain, you need to specify:
- `endpoint`: The API endpoint for that chain's block explorer
- `keys`: A map of API keys and their rate limits (calls per second)

## Supported APIs

The integration supports the following Etherscan API methods:

### Account APIs
- `getAccountBalance`: Get the Ether balance of an address
- `getTokenBalance`: Get the token balance of an address for a specific token
- `getTransactionsByAddress`: Get the list of transactions for an address
- `getInternalTransactionsByAddress`: Get the list of internal transactions for an address
- `getTokenTransfersByAddress`: Get the list of ERC20 token transfers for an address
- `getERC721Transfers`: Get the list of ERC721 token transfers for an address

### Transaction APIs
- `getTransactionByHash`: Get transaction details by transaction hash

### Block APIs
- `getBlockByNumber`: Get block information by block number
- `getBlockRewards`: Get block rewards by block number

### Contract APIs
- `getContractABI`: Get the ABI for a verified contract
- `getContractSourceCode`: Get the source code of a verified contract
- `executeContractMethod`: Execute a read-only contract method call

### Token APIs
- `getTokenInfo`: Get information about an ERC20 token

### Gas APIs
- `getGasOracle`: Get current gas oracle price estimates

## Usage Examples

Here are examples of how to use each API method:

### Get Account Balance
```
{
  "chainID": "1",
  "address": "0x742d35Cc6634C0532925a3b844Bc454e4438f44e"
}
```

### Get Token Balance
```
{
  "chainID": "1",
  "contractAddress": "0xdac17f958d2ee523a2206206994597c13d831ec7",
  "address": "0x742d35Cc6634C0532925a3b844Bc454e4438f44e"
}
```

### Get Transactions By Address
```
{
  "chainID": "1",
  "address": "0x742d35Cc6634C0532925a3b844Bc454e4438f44e",
  "startBlock": "0",
  "endBlock": "99999999",
  "page": "1",
  "offset": "10"
}
```

### Get Transaction By Hash
```
{
  "chainID": "1",
  "txHash": "0x1e2910a262b1008d0616a0beb24c1a491d78771baa54a33e66065e03b1f46bc1"
}
```

### Get Block By Number
```
{
  "chainID": "1",
  "blockNumber": "15000000"
}
```

### Get Contract ABI
```
{
  "chainID": "1",
  "contractAddress": "0xdac17f958d2ee523a2206206994597c13d831ec7"
}
```

## Error Handling

The API integration includes error handling for:
- Invalid inputs
- Rate limiting
- Connection issues
- Invalid API keys

When an error occurs, it will return a JSON object with a status and message indicating the error.

## Rate Limiting

The integration includes built-in rate limiting to respect the Etherscan API limits. It will automatically distribute requests among the available API keys to maximize throughput while staying within the rate limits.

## Adding More Chains

To add support for more chains, update the `config.json` file with the appropriate chain ID, endpoint, and API keys.

## Getting API Keys

To get API keys:

1. Create an account on the appropriate block explorer:
   - Ethereum: [etherscan.io](https://etherscan.io)
   - BSC: [bscscan.com](https://bscscan.com)
   - Polygon: [polygonscan.com](https://polygonscan.com)
   - etc.

2. Navigate to your account's API Keys section
3. Create a new API key
4. Add the API key to your `config.json` file 