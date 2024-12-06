# 1scan

A unified API gateway that for integrating multiple etherscan-like blockchain explorer APIs.

Now you can send request to https://1scan.dev/$chainID for all these supported chains:

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

## Installation

```bash
go install github.com/huahuayu/1scan@latest
```

## Quick Start

1. Create a configuration file `config.json`:

```json
{
  "1": {
    "endpoint": "api.etherscan.io",
    "keys": {
      "yourAPIKey1GetFromEtherscan": 5,
      "yourAPIKey2GetFromEtherscan": 10
    }
  },
  "56": {
    "endpoint": "api.bscscan.com",
    "keys": {
      "yourAPIKey1GetFromBscscan": 5,
      "yourAPIKey2GetFromBscscan": 10
    }
  }
}
```

In the above example, your config contains two chains, `1` and `56` are the chainIDs, and the corresponding explorer api endpoints are `api.etherscan.io` and `api.bscscan.com`. For chainID `1`, you have two API keys, `yourAPIKey1GetFromEtherscan` and `yourAPIKey2GetFromEtherscan`, and the rate limit for each API key is 5 and 10 respectively.

1. Run the service:

```bash
1scan -config /path/to/config.json
```

## Usage

### Basic Request

Send requests to `http://localhost:8080/{chainID}` with your parameters:

```bash
curl "http://localhost:8080/1?module=account&action=balance&address=0x742d35Cc6634C0532925a3b844Bc454e4438f44e"
```

### Using Custom API Key

You can provide your own API key in the request:

```bash
curl "http://localhost:8080/1?apikey=YOUR_CUSTOM_KEY&module=account&action=balance&address=0x742d35Cc6634C0532925a3b844Bc454e4438f44e"
```

## Configuration

### Command Line Flags

- `-config`: Path to configuration file (default: "config.json")
- `-port`: Server port (default: 8080)
- `-loglevel`: Logging level (default: "info", options: "debug", "info", "warn", "error")

### Configuration File Format

```json
{
  "chainID": {
    "endpoint": "api.fooscan.io",
    "keys": {
      "apiKey1": rateLimit1,
      "apiKey2": rateLimit2
    }
  }
}
```

- `chainID`: The blockchain network ID (e.g., 1 for Ethereum mainnet, 56 for BSC)
- `endpoint`: The API endpoint for the blockchain explorer
- `keys`: Map of API keys and their rate limits (requests per second)

## License

MIT License
