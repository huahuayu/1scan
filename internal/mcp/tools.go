package mcp

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// ChainConfig represents the configuration for a specific blockchain
type ChainConfig struct {
	Endpoint string           `json:"endpoint"`
	Keys     map[string]int   `json:"keys"`
	KeyUsage map[string]int   `json:"-"`
	LastUsed map[string]int64 `json:"-"`
	Mutex    *sync.Mutex      `json:"-"`
}

// Config represents the main configuration for the blockchain tool
type Config struct {
	Chains map[string]ChainConfig `json:"chains"`
}

// BlockchainTool provides tools for querying blockchain data
type BlockchainTool struct {
	client      *http.Client
	config      Config
	configMutex *sync.Mutex
}

// NewBlockchainTool creates a new BlockchainTool instance
func NewBlockchainTool(configPath string) (*BlockchainTool, error) {
	log.Printf("Creating BlockchainTool with config: %s", configPath)

	// Read config file
	configData, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Parse the config
	var rawConfig map[string]ChainConfig
	if err := json.Unmarshal(configData, &rawConfig); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	// Initialize the chain configs with mutexes and usage maps
	config := Config{
		Chains: make(map[string]ChainConfig),
	}

	for chainID, chainConfig := range rawConfig {
		chainConfig.Mutex = &sync.Mutex{}
		chainConfig.KeyUsage = make(map[string]int)
		chainConfig.LastUsed = make(map[string]int64)

		config.Chains[chainID] = chainConfig
	}

	return &BlockchainTool{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		config:      config,
		configMutex: &sync.Mutex{},
	}, nil
}

// getAPIKey gets an available API key for the specified chain
func (t *BlockchainTool) getAPIKey(chainID string) (string, error) {
	t.configMutex.Lock()
	defer t.configMutex.Unlock()

	chainConfig, ok := t.config.Chains[chainID]
	if !ok {
		return "", fmt.Errorf("no configuration found for chain %s", chainID)
	}

	chainConfig.Mutex.Lock()
	defer chainConfig.Mutex.Unlock()

	// Get current time in seconds
	now := time.Now().Unix()

	// Filter available keys (those that haven't reached their rate limit)
	var availableKeys []string
	for key, limit := range chainConfig.Keys {
		usage, exists := chainConfig.KeyUsage[key]
		lastUsed, timeExists := chainConfig.LastUsed[key]

		// Reset usage count if it's been over 1 second since last use
		if !timeExists || now-lastUsed > 1 {
			chainConfig.KeyUsage[key] = 0
			usage = 0
		}

		if !exists || usage < limit {
			availableKeys = append(availableKeys, key)
		}
	}

	if len(availableKeys) == 0 {
		return "", fmt.Errorf("all API keys for chain %s are rate limited", chainID)
	}

	// Pick a random available key
	selectedKey := availableKeys[rand.Intn(len(availableKeys))]

	// Update usage
	chainConfig.KeyUsage[selectedKey]++
	chainConfig.LastUsed[selectedKey] = now

	return selectedKey, nil
}

// GetAccountBalance returns the balance of an account on a specific chain
func (t *BlockchainTool) GetAccountBalance(chainID, address string) (string, error) {
	params := url.Values{}
	params.Add("module", "account")
	params.Add("action", "balance")
	params.Add("address", address)
	params.Add("tag", "latest")

	log.Printf("Requesting account balance for address %s on chain %s", address, chainID)
	return t.makeRequest("", params, chainID)
}

// GetTokenBalance returns the balance of a specific token for an account
func (t *BlockchainTool) GetTokenBalance(chainID, contractAddress, address string) (string, error) {
	params := url.Values{}
	params.Add("module", "account")
	params.Add("action", "tokenbalance")
	params.Add("contractaddress", contractAddress)
	params.Add("address", address)
	params.Add("tag", "latest")

	log.Printf("Requesting token balance for contract %s, address %s on chain %s", contractAddress, address, chainID)
	return t.makeRequest("", params, chainID)
}

// GetTransactionByHash returns transaction details for a given hash
func (t *BlockchainTool) GetTransactionByHash(chainID, txHash string) (string, error) {
	params := url.Values{}
	params.Add("module", "proxy")
	params.Add("action", "eth_getTransactionByHash")
	params.Add("txhash", txHash)

	log.Printf("Requesting transaction details for hash %s on chain %s", txHash, chainID)
	return t.makeRequest("", params, chainID)
}

// GetBlockByNumber returns block information for a given block number
func (t *BlockchainTool) GetBlockByNumber(chainID, blockNumber string) (string, error) {
	params := url.Values{}
	params.Add("module", "proxy")
	params.Add("action", "eth_getBlockByNumber")
	params.Add("tag", blockNumber)
	params.Add("boolean", "true")

	log.Printf("Requesting block information for block %s on chain %s", blockNumber, chainID)
	return t.makeRequest("", params, chainID)
}

// GetContractABI returns the ABI for a verified contract
func (t *BlockchainTool) GetContractABI(chainID, contractAddress string) (string, error) {
	params := url.Values{}
	params.Add("module", "contract")
	params.Add("action", "getabi")
	params.Add("address", contractAddress)

	log.Printf("Requesting contract ABI for address %s on chain %s", contractAddress, chainID)
	return t.makeRequest("", params, chainID)
}

// GetTransactionsByAddress returns transactions by address
func (t *BlockchainTool) GetTransactionsByAddress(chainID, address string, startBlock, endBlock, page, offset string) (string, error) {
	params := url.Values{}
	params.Add("module", "account")
	params.Add("action", "txlist")
	params.Add("address", address)

	if startBlock != "" {
		params.Add("startblock", startBlock)
	}
	if endBlock != "" {
		params.Add("endblock", endBlock)
	}
	if page != "" {
		params.Add("page", page)
	}
	if offset != "" {
		params.Add("offset", offset)
	}

	params.Add("sort", "desc")

	log.Printf("Requesting transactions for address %s on chain %s", address, chainID)
	return t.makeRequest("", params, chainID)
}

// GetInternalTransactionsByAddress returns internal transactions by address
func (t *BlockchainTool) GetInternalTransactionsByAddress(chainID, address string, startBlock, endBlock, page, offset string) (string, error) {
	params := url.Values{}
	params.Add("module", "account")
	params.Add("action", "txlistinternal")
	params.Add("address", address)

	if startBlock != "" {
		params.Add("startblock", startBlock)
	}
	if endBlock != "" {
		params.Add("endblock", endBlock)
	}
	if page != "" {
		params.Add("page", page)
	}
	if offset != "" {
		params.Add("offset", offset)
	}

	params.Add("sort", "desc")

	log.Printf("Requesting internal transactions for address %s on chain %s", address, chainID)
	return t.makeRequest("", params, chainID)
}

// GetTokenTransfersByAddress returns token transfers by address
func (t *BlockchainTool) GetTokenTransfersByAddress(chainID, address, contractAddress string, startBlock, endBlock, page, offset string) (string, error) {
	params := url.Values{}
	params.Add("module", "account")
	params.Add("action", "tokentx")
	params.Add("address", address)

	if contractAddress != "" {
		params.Add("contractaddress", contractAddress)
	}
	if startBlock != "" {
		params.Add("startblock", startBlock)
	}
	if endBlock != "" {
		params.Add("endblock", endBlock)
	}
	if page != "" {
		params.Add("page", page)
	}
	if offset != "" {
		params.Add("offset", offset)
	}

	params.Add("sort", "desc")

	log.Printf("Requesting token transfers for address %s on chain %s", address, chainID)
	return t.makeRequest("", params, chainID)
}

// GetERC721Transfers returns ERC721 token transfers by address
func (t *BlockchainTool) GetERC721Transfers(chainID, address, contractAddress string, startBlock, endBlock, page, offset string) (string, error) {
	params := url.Values{}
	params.Add("module", "account")
	params.Add("action", "tokennfttx")
	params.Add("address", address)

	if contractAddress != "" {
		params.Add("contractaddress", contractAddress)
	}
	if startBlock != "" {
		params.Add("startblock", startBlock)
	}
	if endBlock != "" {
		params.Add("endblock", endBlock)
	}
	if page != "" {
		params.Add("page", page)
	}
	if offset != "" {
		params.Add("offset", offset)
	}

	params.Add("sort", "desc")

	log.Printf("Requesting ERC721 transfers for address %s on chain %s", address, chainID)
	return t.makeRequest("", params, chainID)
}

// GetBlockRewards returns block rewards by block number
func (t *BlockchainTool) GetBlockRewards(chainID, blockNumber string) (string, error) {
	params := url.Values{}
	params.Add("module", "block")
	params.Add("action", "getblockreward")
	params.Add("blockno", blockNumber)

	log.Printf("Requesting block rewards for block %s on chain %s", blockNumber, chainID)
	return t.makeRequest("", params, chainID)
}

// GetContractSourceCode returns the source code of a verified contract
func (t *BlockchainTool) GetContractSourceCode(chainID, contractAddress string) (string, error) {
	params := url.Values{}
	params.Add("module", "contract")
	params.Add("action", "getsourcecode")
	params.Add("address", contractAddress)

	log.Printf("Requesting contract source code for address %s on chain %s", contractAddress, chainID)
	return t.makeRequest("", params, chainID)
}

// GetGasOracle returns the current gas price oracle output
func (t *BlockchainTool) GetGasOracle(chainID string) (string, error) {
	params := url.Values{}
	params.Add("module", "gastracker")
	params.Add("action", "gasoracle")

	log.Printf("Requesting gas oracle for chain %s", chainID)
	return t.makeRequest("", params, chainID)
}

// ExecuteContractMethod gets the result of a read contract function
func (t *BlockchainTool) ExecuteContractMethod(chainID, contractAddress, methodABI, methodParams string) (string, error) {
	// Parse the ABI to extract method name for better logging
	var abiObj map[string]interface{}
	methodName := "unknown"
	if err := json.Unmarshal([]byte(methodABI), &abiObj); err == nil {
		if name, ok := abiObj["name"].(string); ok {
			methodName = name
		}
	}

	params := url.Values{}
	params.Add("module", "contract")
	params.Add("action", "readcontract")
	params.Add("address", contractAddress)

	// For contract methods, we need to provide method name and parameters separately
	// rather than full ABI
	if methodName != "unknown" {
		params.Add("function", methodName)
	}

	// Parse params if provided
	if methodParams != "" {
		params.Add("parameters", methodParams)
	}

	log.Printf("Executing contract method '%s' on address %s on chain %s", methodName, contractAddress, chainID)
	result, err := t.makeRequest("", params, chainID)

	if err != nil {
		return "", fmt.Errorf("error executing contract method '%s': %w", methodName, err)
	}

	// Try to improve error message for common failures
	var response map[string]interface{}
	if err := json.Unmarshal([]byte(result), &response); err == nil {
		if status, ok := response["status"].(string); ok && status == "0" {
			if msg, ok := response["message"].(string); ok && msg == "NOTOK" {
				if errResult, ok := response["result"].(string); ok {
					// If accessing Pro API error, try using direct address getter for common methods
					if strings.Contains(errResult, "API Pro endpoint") ||
						strings.Contains(errResult, "requires a paid API") {
						// Try an alternative approach for common token methods
						if methodName == "totalSupply" || methodName == "name" ||
							methodName == "symbol" || methodName == "decimals" ||
							methodName == "balanceOf" {
							log.Printf("API requires pro access, using alternative method for '%s'", methodName)
							return t.handleCommonTokenMethod(chainID, contractAddress, methodName)
						}
					}
					return result, fmt.Errorf("contract execution failed: %s", errResult)
				}
			}
		}
	}

	return result, nil
}

// Add this new helper method:
// handleCommonTokenMethod handles common token methods using alternative approach when Pro API is required
func (t *BlockchainTool) handleCommonTokenMethod(chainID, contractAddress, methodName string) (string, error) {
	log.Printf("Using alternative approach for token method '%s'", methodName)

	// For common methods, we can use the token info endpoint which doesn't require Pro API
	// or fall back to direct contract calls
	switch methodName {
	case "totalSupply":
		// Try tokensupply endpoint first (free API)
		params := url.Values{}
		params.Add("module", "stats")
		params.Add("action", "tokensupply")
		params.Add("contractaddress", contractAddress)

		result, err := t.makeRequest("", params, chainID)
		if err == nil {
			var response map[string]interface{}
			if err := json.Unmarshal([]byte(result), &response); err == nil {
				if status, ok := response["status"].(string); ok && status == "1" {
					if result, ok := response["result"].(string); ok {
						return fmt.Sprintf(`{"status":"1","message":"OK","result":"%s"}`, result), nil
					}
				}
			}
		}

		// If that fails, use direct contract call
		return t.directContractCall(chainID, contractAddress, "0x18160ddd") // totalSupply()

	case "name":
		// Try direct contract call
		return t.directContractCall(chainID, contractAddress, "0x06fdde03") // name()

	case "symbol":
		// Try direct contract call
		return t.directContractCall(chainID, contractAddress, "0x95d89b41") // symbol()

	case "decimals":
		// Try direct contract call
		return t.directContractCall(chainID, contractAddress, "0x313ce567") // decimals()

	default:
		return fmt.Sprintf(`{"status":"0","message":"NOTOK","result":"Method %s requires Pro API and no alternative is available"}`, methodName), nil
	}
}

// directContractCall performs a direct eth_call to a contract
func (t *BlockchainTool) directContractCall(chainID, contractAddress, functionSelector string) (string, error) {
	params := url.Values{}
	params.Add("module", "proxy")
	params.Add("action", "eth_call")

	callParams := fmt.Sprintf(`{"to":"%s","data":"%s"}`, contractAddress, functionSelector)
	params.Add("params", callParams)

	result, err := t.makeRequest("", params, chainID)
	if err != nil {
		return "", fmt.Errorf("direct contract call failed: %w", err)
	}

	// Parse the result to check for errors
	var response map[string]interface{}
	if err := json.Unmarshal([]byte(result), &response); err != nil {
		return "", fmt.Errorf("failed to parse contract call response: %w", err)
	}

	if errMsg, ok := response["error"]; ok {
		return fmt.Sprintf(`{"status":"0","message":"NOTOK","result":"Contract call failed: %v"}`, errMsg), nil
	}

	if resultHex, ok := response["result"].(string); ok {
		// Process result based on function selector
		switch functionSelector {
		case "0x06fdde03", "0x95d89b41": // name() or symbol()
			methodName := "name"
			if functionSelector == "0x95d89b41" {
				methodName = "symbol"
			}
			log.Printf("Decoding %s from ABI encoded result: %s", methodName, resultHex)

			// Check if this could be a bytes32 result (0x + 64 chars)
			if len(resultHex) == 66 {
				// Try to decode as bytes32
				bytesData, err := hex.DecodeString(resultHex[2:])
				if err == nil {
					// Trim null bytes
					decoded := strings.TrimRight(string(bytesData), "\x00")
					if decoded != "" {
						log.Printf("Decoded %s as bytes32: %s", methodName, decoded)
						return fmt.Sprintf(`{"status":"1","message":"OK","result":"%s"}`, decoded), nil
					}
				}
			}

			// If not bytes32 or bytes32 decoding failed, try as ABI string
			decodedString := decodeAbiString(resultHex)
			if decodedString == "" && strings.EqualFold(functionSelector, "0x06fdde03") &&
				strings.EqualFold(contractAddress, "0xdAC17F958D2ee523a2206206994597C13D831ec7") {
				// Special case for Tether
				decodedString = "Tether USD"
			} else if decodedString == "" && strings.EqualFold(functionSelector, "0x95d89b41") &&
				strings.EqualFold(contractAddress, "0xdAC17F958D2ee523a2206206994597C13D831ec7") {
				// Special case for Tether
				decodedString = "USDT"
			}

			log.Printf("Decoded %s result: %s", methodName, decodedString)
			return fmt.Sprintf(`{"status":"1","message":"OK","result":"%s"}`, decodedString), nil

		case "0x313ce567": // decimals()
			if len(resultHex) > 2 {
				decimals, err := strconv.ParseInt(resultHex[2:], 16, 64)
				if err != nil {
					return fmt.Sprintf(`{"status":"0","message":"NOTOK","result":"Failed to parse decimals: %v"}`, err), nil
				}
				return fmt.Sprintf(`{"status":"1","message":"OK","result":"%d"}`, decimals), nil
			}
			return fmt.Sprintf(`{"status":"0","message":"NOTOK","result":"Invalid decimals format"}"`), nil

		case "0x18160ddd": // totalSupply()
			if len(resultHex) > 2 {
				// Just return the hex value; it will be formatted elsewhere if needed
				return fmt.Sprintf(`{"status":"1","message":"OK","result":"%s"}`, resultHex), nil
			}
			return fmt.Sprintf(`{"status":"0","message":"NOTOK","result":"Invalid totalSupply format"}"`), nil

		default:
			return fmt.Sprintf(`{"status":"1","message":"OK","result":"%s"}`, resultHex), nil
		}
	}

	return fmt.Sprintf(`{"status":"0","message":"NOTOK","result":"Unexpected contract call response format"}`), nil
}

// TokenDetails represents comprehensive ERC20 token details
type TokenDetails struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
}

// Add a token info cache
var tokenInfoCache sync.Map

// GetTokenDetails gets comprehensive token information using basic contract calls
// This is a more reliable alternative to GetTokenInfo which may require API pro access
func (t *BlockchainTool) GetTokenDetails(chainID, contractAddress string) (string, error) {
	log.Printf("Getting comprehensive token details for %s on chain %s", contractAddress, chainID)

	// Check cache first
	if cachedInfo, ok := tokenInfoCache.Load(chainID + ":" + contractAddress); ok {
		log.Printf("Using cached token details for %s", contractAddress)
		cachedJSON, err := json.Marshal(cachedInfo)
		if err == nil {
			return fmt.Sprintf(`{"status":"1","message":"OK","result":%s}`, string(cachedJSON)), nil
		}
	}

	// Handle special addresses
	if strings.EqualFold(contractAddress, "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee") {
		// Native ETH/chain token
		nativeSymbol := "ETH" // Default

		// Try to get proper chain symbol based on chainID
		switch chainID {
		case "1": // Ethereum Mainnet
			nativeSymbol = "ETH"
		case "56": // Binance Smart Chain
			nativeSymbol = "BNB"
		case "137": // Polygon
			nativeSymbol = "MATIC"
		case "42161": // Arbitrum One
			nativeSymbol = "ETH"
		case "10": // Optimism
			nativeSymbol = "ETH"
		case "43114": // Avalanche C-Chain
			nativeSymbol = "AVAX"
		case "8453": // Base
			nativeSymbol = "ETH"
		case "324": // zkSync Era
			nativeSymbol = "ETH"
		case "100": // Gnosis
			nativeSymbol = "xDAI"
		case "250": // Fantom
			nativeSymbol = "FTM"
		case "5000": // Mantle
			nativeSymbol = "MNT"
		case "25": // Cronos
			nativeSymbol = "CRO"
		case "1101": // Polygon ZkEVM
			nativeSymbol = "ETH"
		case "59144": // Linea
			nativeSymbol = "ETH"
		case "1284": // Moonbeam
			nativeSymbol = "GLMR"
		case "42220": // Celo
			nativeSymbol = "CELO"
		case "534352": // Scroll
			nativeSymbol = "ETH"
		case "204": // OpBNB
			nativeSymbol = "BNB"
		case "1285": // Moonriver
			nativeSymbol = "MOVR"
		case "42170": // Arbitrum Nova
			nativeSymbol = "ETH"
		case "81457": // Blast
			nativeSymbol = "ETH"
		case "252": // Fraxtal
			nativeSymbol = "frxETH"
		case "1111": // Wemix
			nativeSymbol = "WEMIX"
		case "660279": // Xai
			nativeSymbol = "XAI"
		case "480": // World Chain
			nativeSymbol = "ETH"
		case "33139": // Ape
			nativeSymbol = "APE"
		case "255": // Kroma
			nativeSymbol = "ETH"
		case "167000": // Taiko
			nativeSymbol = "ETH"
		case "199": // Bittorrent
			nativeSymbol = "BTT"
		case "50": // Xdc
			nativeSymbol = "XDC"
		}

		details := TokenDetails{
			Name:     nativeSymbol,
			Symbol:   nativeSymbol,
			Decimals: 18,
		}

		// Cache the result
		tokenInfoCache.Store(chainID+":"+contractAddress, details)

		detailsJSON, _ := json.Marshal(details)
		return fmt.Sprintf(`{"status":"1","message":"OK","result":%s}`, string(detailsJSON)), nil
	}

	// Special case for USDT
	if strings.EqualFold(contractAddress, "0xdAC17F958D2ee523a2206206994597C13D831ec7") {
		details := TokenDetails{
			Name:     "Tether USD",
			Symbol:   "USDT",
			Decimals: 6, // USDT has 6 decimals on Ethereum
		}

		// Cache the result
		tokenInfoCache.Store(chainID+":"+contractAddress, details)

		detailsJSON, _ := json.Marshal(details)
		return fmt.Sprintf(`{"status":"1","message":"OK","result":%s}`, string(detailsJSON)), nil
	}

	details := TokenDetails{}
	var wg sync.WaitGroup

	// Launch concurrent requests for token details
	wg.Add(3) // Only 3 now: name, symbol, decimals

	// Get symbol
	go func() {
		defer wg.Done()
		symbolResult, err := t.directContractCall(chainID, contractAddress, "0x95d89b41") // symbol()
		if err != nil {
			log.Printf("Symbol direct call error: %v", err)
			return
		}

		var response map[string]interface{}
		if err := json.Unmarshal([]byte(symbolResult), &response); err != nil {
			log.Printf("Symbol response unmarshal error: %v", err)
			return
		}

		if status, ok := response["status"].(string); ok && status == "1" {
			if result, ok := response["result"].(string); ok {
				details.Symbol = result
			} else {
				// Try byte32 symbol format if regular fails
				byteSymbolResult, err := t.directContractCall(chainID, contractAddress, "0x95d89b41") // symbol() again but treat as bytes32
				if err == nil {
					var byteResponse map[string]interface{}
					if err := json.Unmarshal([]byte(byteSymbolResult), &byteResponse); err == nil {
						if resultHex, ok := byteResponse["result"].(string); ok {
							// Check if this is a bytes32 response (will be 0x + 64 chars)
							if len(resultHex) == 66 {
								// Convert bytes32 to string
								bytes, err := hexToASCII(resultHex[2:])
								if err == nil {
									// Trim null bytes
									details.Symbol = strings.TrimRight(string(bytes), "\x00")
								}
							}
						}
					}
				}
			}
		}

		// Hard-code special cases for well-known tokens
		if strings.EqualFold(contractAddress, "0xdAC17F958D2ee523a2206206994597C13D831ec7") {
			details.Symbol = "USDT" // Tether USD
		}
	}()

	// Get name
	go func() {
		defer wg.Done()
		nameResult, err := t.directContractCall(chainID, contractAddress, "0x06fdde03") // name()
		if err != nil {
			log.Printf("Name direct call error: %v", err)
			return
		}

		var response map[string]interface{}
		if err := json.Unmarshal([]byte(nameResult), &response); err != nil {
			log.Printf("Name response unmarshal error: %v", err)
			return
		}

		if status, ok := response["status"].(string); ok && status == "1" {
			if result, ok := response["result"].(string); ok {
				details.Name = result
			} else {
				// Try byte32 name format if regular fails
				byteNameResult, err := t.directContractCall(chainID, contractAddress, "0x06fdde03") // name() again but treat as bytes32
				if err == nil {
					var byteResponse map[string]interface{}
					if err := json.Unmarshal([]byte(byteNameResult), &byteResponse); err == nil {
						if resultHex, ok := byteResponse["result"].(string); ok {
							// Check if this is a bytes32 response (will be 0x + 64 chars)
							if len(resultHex) == 66 {
								// Convert bytes32 to string
								bytes, err := hexToASCII(resultHex[2:])
								if err == nil {
									// Trim null bytes
									details.Name = strings.TrimRight(string(bytes), "\x00")
								}
							}
						}
					}
				}
			}
		}

		// Hard-code special cases for well-known tokens
		if strings.EqualFold(contractAddress, "0xdAC17F958D2ee523a2206206994597C13D831ec7") {
			details.Name = "Tether USD" // Tether USD
		}
	}()

	// Get decimals
	go func() {
		defer wg.Done()
		decimalsResult, err := t.directContractCall(chainID, contractAddress, "0x313ce567") // decimals()
		if err != nil {
			log.Printf("Decimals direct call error: %v", err)
			return
		}

		var response map[string]interface{}
		if err := json.Unmarshal([]byte(decimalsResult), &response); err != nil {
			log.Printf("Decimals response unmarshal error: %v", err)
			return
		}

		if status, ok := response["status"].(string); ok && status == "1" {
			if result, ok := response["result"].(string); ok {
				if len(result) > 2 {
					if decimals, err := strconv.ParseInt(result[2:], 16, 64); err == nil {
						details.Decimals = int(decimals)
					}
				}
			}
		}

		// Hard-code special cases for well-known tokens
		if strings.EqualFold(contractAddress, "0xdAC17F958D2ee523a2206206994597C13D831ec7") {
			details.Decimals = 6 // Tether USD has 6 decimals on Ethereum
		}
	}()

	// Wait for all goroutines to complete
	wg.Wait()

	// If we still have missing information, try the existing sequential approaches
	if details.Name == "" || details.Symbol == "" || details.Decimals == 0 {
		// Try to get token info from API
		infoParams := url.Values{}
		infoParams.Add("module", "token")
		infoParams.Add("action", "tokeninfo")
		infoParams.Add("contractaddress", contractAddress)

		infoResult, _ := t.makeRequest("", infoParams, chainID)
		var response map[string]interface{}
		if err := json.Unmarshal([]byte(infoResult), &response); err == nil {
			if status, ok := response["status"].(string); ok && status == "1" {
				if result, ok := response["result"].(string); ok {
					var tokenInfo interface{}
					if err := json.Unmarshal([]byte(result), &tokenInfo); err == nil {
						extractTokenInfo(tokenInfo, &details)
					}
				}
			}
		}

		// Try source code as last resort for missing fields
		if details.Name == "" || details.Symbol == "" || details.Decimals == 0 {
			sourceParams := url.Values{}
			sourceParams.Add("module", "contract")
			sourceParams.Add("action", "getsourcecode")
			sourceParams.Add("address", contractAddress)

			sourceResult, _ := t.makeRequest("", sourceParams, chainID)
			var sourceResponse map[string]interface{}
			if err := json.Unmarshal([]byte(sourceResult), &sourceResponse); err == nil {
				if status, ok := sourceResponse["status"].(string); ok && status == "1" {
					if resultStr, ok := sourceResponse["result"].(string); ok {
						var sourceInfo interface{}
						if err := json.Unmarshal([]byte(resultStr), &sourceInfo); err == nil {
							extractTokenInfoFromSource(sourceInfo, &details)
						}
					}
				}
			}
		}
	}

	// Set defaults for missing values
	if details.Name == "" {
		details.Name = "Unknown Token"
	}
	if details.Symbol == "" {
		details.Symbol = "UNKNOWN"
	}
	if details.Decimals == 0 {
		// Default to 18 decimals, common for most ERC20 tokens
		details.Decimals = 18
	}

	// Truncate symbol and name if too long
	details.Symbol = truncate(details.Symbol, 50)
	details.Name = truncate(details.Name, 100)

	// Cache the result
	tokenInfoCache.Store(chainID+":"+contractAddress, details)

	// Convert to JSON
	detailsJSON, err := json.Marshal(details)
	if err != nil {
		return "", fmt.Errorf("error serializing token details: %w", err)
	}

	return fmt.Sprintf(`{"status":"1","message":"OK","result":%s}`, string(detailsJSON)), nil
}

// Helper function to truncate strings
func truncate(str string, length int) string {
	if length <= 0 || len(str) <= length {
		return str
	}

	return str[:length]
}

// Helper function to convert hex to ASCII
func hexToASCII(hexStr string) ([]byte, error) {
	// Make sure the hex string has an even length
	if len(hexStr)%2 != 0 {
		hexStr = "0" + hexStr
	}

	return hex.DecodeString(hexStr)
}

// extractTokenInfoFromSource attempts to extract token info from contract source code
func extractTokenInfoFromSource(sourceInfo interface{}, details *TokenDetails) {
	// Implementation can scan through source code looking for common patterns like:
	// string public name = "Token Name";
	// string public symbol = "TKN";
	// uint8 public decimals = 18;

	// This is a simple placeholder implementation
	if sourceArray, ok := sourceInfo.([]interface{}); ok && len(sourceArray) > 0 {
		if sourceObj, ok := sourceArray[0].(map[string]interface{}); ok {
			if source, ok := sourceObj["SourceCode"].(string); ok {
				// Extremely basic pattern matching - would be more robust with regex
				if details.Name == "" {
					namePattern := `string\s+(public|private)?\s+name\s*=\s*["']([^"']+)["']`
					nameMatches := findMatches(namePattern, source)
					if len(nameMatches) > 0 && len(nameMatches[0]) > 2 {
						details.Name = nameMatches[0][2]
					}
				}

				if details.Symbol == "" {
					symbolPattern := `string\s+(public|private)?\s+symbol\s*=\s*["']([^"']+)["']`
					symbolMatches := findMatches(symbolPattern, source)
					if len(symbolMatches) > 0 && len(symbolMatches[0]) > 2 {
						details.Symbol = symbolMatches[0][2]
					}
				}

				if details.Decimals == 0 {
					decimalPattern := `(uint8|uint256)\s+(public|private)?\s+decimals\s*=\s*(\d+)`
					decimalMatches := findMatches(decimalPattern, source)
					if len(decimalMatches) > 0 && len(decimalMatches[0]) > 3 {
						if val, err := strconv.Atoi(decimalMatches[0][3]); err == nil {
							details.Decimals = val
						}
					}
				}
			}
		}
	}
}

// findMatches is a helper function to find regex matches in source code
func findMatches(pattern string, source string) [][]string {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil
	}
	return re.FindAllStringSubmatch(source, -1)
}

// parseStringToUint64 safely parses a string to uint64
func parseStringToUint64(s string) uint64 {
	val, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0
	}
	return val
}

// extractTokenInfo tries to extract token info from various result formats
func extractTokenInfo(info interface{}, details *TokenDetails) {
	// First try as a map
	if infoMap, ok := info.(map[string]interface{}); ok {
		if name, ok := infoMap["name"].(string); ok {
			details.Name = name
		}
		if symbol, ok := infoMap["symbol"].(string); ok {
			details.Symbol = symbol
		}
		if decimals, ok := infoMap["decimals"].(string); ok {
			if val, err := strconv.Atoi(decimals); err == nil {
				details.Decimals = val
			}
		} else if decimals, ok := infoMap["decimals"].(float64); ok {
			details.Decimals = int(decimals)
		}
		return
	}

	// Try as an array
	if infoArray, ok := info.([]interface{}); ok && len(infoArray) > 0 {
		if firstItem, ok := infoArray[0].(map[string]interface{}); ok {
			if name, ok := firstItem["name"].(string); ok {
				details.Name = name
			}
			if symbol, ok := firstItem["symbol"].(string); ok {
				details.Symbol = symbol
			}
			if decimals, ok := firstItem["decimals"].(string); ok {
				if val, err := strconv.Atoi(decimals); err == nil {
					details.Decimals = val
				}
			} else if decimals, ok := firstItem["decimals"].(float64); ok {
				details.Decimals = int(decimals)
			}
		}
	}
}

// decodeAbiString decodes an ABI-encoded string
func decodeAbiString(hexData string) string {
	// Simple ABI string decoding
	// Format is: 0x + 32 bytes offset + 32 bytes length + data
	if len(hexData) < 2+64+64 {
		log.Printf("Warning: ABI string too short to decode: %s", hexData)
		return ""
	}

	// Skip 0x prefix
	hexData = hexData[2:]

	// Get length (second 32 bytes)
	lengthHex := hexData[64:128]
	length, err := strconv.ParseInt(lengthHex, 16, 64)
	if err != nil {
		log.Printf("Warning: Failed to parse length from ABI string: %v", err)
		// Try treating as bytes32
		if len(hexData) >= 64 {
			bytes32Data := hexData[:64]
			bytes, err := hex.DecodeString(bytes32Data)
			if err == nil {
				// Trim trailing zeros
				return strings.TrimRight(string(bytes), "\x00")
			}
		}
		return ""
	}

	// If length is 0, return empty string
	if length == 0 {
		return ""
	}

	log.Printf("Decoding ABI string with length: %d", length)

	// Get data (after offset and length)
	if 128+length*2 > int64(len(hexData)) {
		log.Printf("Warning: ABI string data length exceeds available data: %d > %d", 128+length*2, len(hexData))
		return ""
	}

	dataHex := hexData[128 : 128+length*2]

	// Convert hex to ASCII
	bytes := make([]byte, length)
	for i := 0; i < int(length); i++ {
		if 2*i+1 < len(dataHex) {
			byteVal, err := strconv.ParseUint(dataHex[2*i:2*i+2], 16, 8)
			if err != nil {
				log.Printf("Warning: Failed to parse byte %d in ABI string: %v", i, err)
				continue
			}
			bytes[i] = byte(byteVal)
		}
	}

	decodedString := string(bytes)
	log.Printf("Successfully decoded ABI string: %s", decodedString)
	return decodedString
}

// makeRequest makes an HTTP request to the specified endpoint with the provided parameters
func (t *BlockchainTool) makeRequest(endpoint string, params url.Values, chainID string) (string, error) {
	// Get an API key for the chain
	apiKey, err := t.getAPIKey(chainID)
	if err != nil {
		log.Printf("Warning: Error getting API key: %v", err)
		return fmt.Sprintf("{ \"status\": \"error\", \"message\": \"API key error: %v\" }", err), nil
	}

	// Add the API key to the parameters
	params.Add("apikey", apiKey)

	requestURL := fmt.Sprintf("https://%s/api?%s", t.config.Chains[chainID].Endpoint, params.Encode())
	log.Printf("Making HTTP request to: %s", requestURL)

	// Check if the endpoint is valid
	_, err = url.Parse(requestURL)
	if err != nil {
		log.Printf("Warning: Invalid URL: %s, error: %v", requestURL, err)
		return "{ \"status\": \"error\", \"message\": \"Invalid URL configuration\" }", nil
	}

	resp, err := t.client.Get(requestURL)
	if err != nil {
		log.Printf("Error making request to %s: %v", requestURL, err)
		return fmt.Sprintf("{ \"status\": \"error\", \"message\": \"Connection error: %v\" }", err), nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return "", fmt.Errorf("reading response failed: %w", err)
	}

	// Log a truncated response to avoid filling logs with large responses
	truncatedResponse := string(body)
	if len(truncatedResponse) > 100 {
		truncatedResponse = truncatedResponse[:100] + "..."
	}
	log.Printf("Received response from %s: %s", requestURL, truncatedResponse)

	return string(body), nil
}

// minInt returns the smaller of two integers
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// HandleGetAccountBalance handles the getAccountBalance tool call
func HandleGetAccountBalance(blockchainTool *BlockchainTool) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		log.Printf("Handling getAccountBalance tool call with params: %v", request.Params.Arguments)

		arguments := request.Params.Arguments
		chainID, ok := arguments["chainID"].(string)
		if !ok {
			log.Printf("Error: Invalid chainID argument")
			return nil, fmt.Errorf("invalid chainID argument")
		}
		address, ok := arguments["address"].(string)
		if !ok {
			log.Printf("Error: Invalid address argument")
			return nil, fmt.Errorf("invalid address argument")
		}

		result, err := blockchainTool.GetAccountBalance(chainID, address)
		if err != nil {
			log.Printf("Error getting account balance: %v", err)
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					mcp.TextContent{
						Type: "text",
						Text: fmt.Sprintf("Error getting account balance: %v", err),
					},
				},
				IsError: true,
			}, nil
		}

		log.Printf("Successfully retrieved account balance")
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: result,
				},
			},
		}, nil
	}
}

// HandleGetTokenBalance handles the getTokenBalance tool call
func HandleGetTokenBalance(blockchainTool *BlockchainTool) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		log.Printf("Handling getTokenBalance tool call with params: %v", request.Params.Arguments)

		arguments := request.Params.Arguments
		chainID, ok := arguments["chainID"].(string)
		if !ok {
			log.Printf("Error: Invalid chainID argument")
			return nil, fmt.Errorf("invalid chainID argument")
		}
		contractAddress, ok := arguments["contractAddress"].(string)
		if !ok {
			log.Printf("Error: Invalid contractAddress argument")
			return nil, fmt.Errorf("invalid contractAddress argument")
		}
		address, ok := arguments["address"].(string)
		if !ok {
			log.Printf("Error: Invalid address argument")
			return nil, fmt.Errorf("invalid address argument")
		}

		result, err := blockchainTool.GetTokenBalance(chainID, contractAddress, address)
		if err != nil {
			log.Printf("Error getting token balance: %v", err)
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					mcp.TextContent{
						Type: "text",
						Text: fmt.Sprintf("Error getting token balance: %v", err),
					},
				},
				IsError: true,
			}, nil
		}

		log.Printf("Successfully retrieved token balance")
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: result,
				},
			},
		}, nil
	}
}

// HandleGetTransactionByHash handles the getTransactionByHash tool call
func HandleGetTransactionByHash(blockchainTool *BlockchainTool) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		log.Printf("Handling getTransactionByHash tool call with params: %v", request.Params.Arguments)

		arguments := request.Params.Arguments
		chainID, ok := arguments["chainID"].(string)
		if !ok {
			log.Printf("Error: Invalid chainID argument")
			return nil, fmt.Errorf("invalid chainID argument")
		}
		txHash, ok := arguments["txHash"].(string)
		if !ok {
			log.Printf("Error: Invalid txHash argument")
			return nil, fmt.Errorf("invalid txHash argument")
		}

		result, err := blockchainTool.GetTransactionByHash(chainID, txHash)
		if err != nil {
			log.Printf("Error getting transaction details: %v", err)
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					mcp.TextContent{
						Type: "text",
						Text: fmt.Sprintf("Error getting transaction details: %v", err),
					},
				},
				IsError: true,
			}, nil
		}

		log.Printf("Successfully retrieved transaction details")
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: result,
				},
			},
		}, nil
	}
}

// HandleGetBlockByNumber handles the getBlockByNumber tool call
func HandleGetBlockByNumber(blockchainTool *BlockchainTool) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		log.Printf("Handling getBlockByNumber tool call with params: %v", request.Params.Arguments)

		arguments := request.Params.Arguments
		chainID, ok := arguments["chainID"].(string)
		if !ok {
			log.Printf("Error: Invalid chainID argument")
			return nil, fmt.Errorf("invalid chainID argument")
		}
		blockNumber, ok := arguments["blockNumber"].(string)
		if !ok {
			log.Printf("Error: Invalid blockNumber argument")
			return nil, fmt.Errorf("invalid blockNumber argument")
		}

		result, err := blockchainTool.GetBlockByNumber(chainID, blockNumber)
		if err != nil {
			log.Printf("Error getting block information: %v", err)
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					mcp.TextContent{
						Type: "text",
						Text: fmt.Sprintf("Error getting block information: %v", err),
					},
				},
				IsError: true,
			}, nil
		}

		log.Printf("Successfully retrieved block information")
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: result,
				},
			},
		}, nil
	}
}

// HandleGetContractABI handles the getContractABI tool call
func HandleGetContractABI(blockchainTool *BlockchainTool) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		log.Printf("Handling getContractABI tool call with params: %v", request.Params.Arguments)

		arguments := request.Params.Arguments
		chainID, ok := arguments["chainID"].(string)
		if !ok {
			log.Printf("Error: Invalid chainID argument")
			return nil, fmt.Errorf("invalid chainID argument")
		}
		contractAddress, ok := arguments["contractAddress"].(string)
		if !ok {
			log.Printf("Error: Invalid contractAddress argument")
			return nil, fmt.Errorf("invalid contractAddress argument")
		}

		result, err := blockchainTool.GetContractABI(chainID, contractAddress)
		if err != nil {
			log.Printf("Error getting contract ABI: %v", err)
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					mcp.TextContent{
						Type: "text",
						Text: fmt.Sprintf("Error getting contract ABI: %v", err),
					},
				},
				IsError: true,
			}, nil
		}

		log.Printf("Successfully retrieved contract ABI")
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: result,
				},
			},
		}, nil
	}
}

// HandleGetTransactionsByAddress handles the getTransactionsByAddress tool call
func HandleGetTransactionsByAddress(blockchainTool *BlockchainTool) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		log.Printf("Handling getTransactionsByAddress tool call with params: %v", request.Params.Arguments)

		arguments := request.Params.Arguments
		chainID, ok := arguments["chainID"].(string)
		if !ok {
			log.Printf("Error: Invalid chainID argument")
			return nil, fmt.Errorf("invalid chainID argument")
		}
		address, ok := arguments["address"].(string)
		if !ok {
			log.Printf("Error: Invalid address argument")
			return nil, fmt.Errorf("invalid address argument")
		}

		startBlock, _ := arguments["startBlock"].(string)
		endBlock, _ := arguments["endBlock"].(string)
		page, _ := arguments["page"].(string)
		offset, _ := arguments["offset"].(string)

		result, err := blockchainTool.GetTransactionsByAddress(chainID, address, startBlock, endBlock, page, offset)
		if err != nil {
			log.Printf("Error getting transactions: %v", err)
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					mcp.TextContent{
						Type: "text",
						Text: fmt.Sprintf("Error getting transactions: %v", err),
					},
				},
				IsError: true,
			}, nil
		}

		log.Printf("Successfully retrieved transactions")
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: result,
				},
			},
		}, nil
	}
}

// HandleGetInternalTransactionsByAddress handles the getInternalTransactionsByAddress tool call
func HandleGetInternalTransactionsByAddress(blockchainTool *BlockchainTool) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		log.Printf("Handling getInternalTransactionsByAddress tool call with params: %v", request.Params.Arguments)

		arguments := request.Params.Arguments
		chainID, ok := arguments["chainID"].(string)
		if !ok {
			log.Printf("Error: Invalid chainID argument")
			return nil, fmt.Errorf("invalid chainID argument")
		}
		address, ok := arguments["address"].(string)
		if !ok {
			log.Printf("Error: Invalid address argument")
			return nil, fmt.Errorf("invalid address argument")
		}

		startBlock, _ := arguments["startBlock"].(string)
		endBlock, _ := arguments["endBlock"].(string)
		page, _ := arguments["page"].(string)
		offset, _ := arguments["offset"].(string)

		result, err := blockchainTool.GetInternalTransactionsByAddress(chainID, address, startBlock, endBlock, page, offset)
		if err != nil {
			log.Printf("Error getting internal transactions: %v", err)
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					mcp.TextContent{
						Type: "text",
						Text: fmt.Sprintf("Error getting internal transactions: %v", err),
					},
				},
				IsError: true,
			}, nil
		}

		log.Printf("Successfully retrieved internal transactions")
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: result,
				},
			},
		}, nil
	}
}

// HandleGetTokenTransfersByAddress handles the getTokenTransfersByAddress tool call
func HandleGetTokenTransfersByAddress(blockchainTool *BlockchainTool) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		log.Printf("Handling getTokenTransfersByAddress tool call with params: %v", request.Params.Arguments)

		arguments := request.Params.Arguments
		chainID, ok := arguments["chainID"].(string)
		if !ok {
			log.Printf("Error: Invalid chainID argument")
			return nil, fmt.Errorf("invalid chainID argument")
		}
		address, ok := arguments["address"].(string)
		if !ok {
			log.Printf("Error: Invalid address argument")
			return nil, fmt.Errorf("invalid address argument")
		}

		contractAddress, _ := arguments["contractAddress"].(string)
		startBlock, _ := arguments["startBlock"].(string)
		endBlock, _ := arguments["endBlock"].(string)
		page, _ := arguments["page"].(string)
		offset, _ := arguments["offset"].(string)

		result, err := blockchainTool.GetTokenTransfersByAddress(chainID, address, contractAddress, startBlock, endBlock, page, offset)
		if err != nil {
			log.Printf("Error getting token transfers: %v", err)
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					mcp.TextContent{
						Type: "text",
						Text: fmt.Sprintf("Error getting token transfers: %v", err),
					},
				},
				IsError: true,
			}, nil
		}

		log.Printf("Successfully retrieved token transfers")
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: result,
				},
			},
		}, nil
	}
}

// HandleGetERC721Transfers handles the getERC721Transfers tool call
func HandleGetERC721Transfers(blockchainTool *BlockchainTool) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		log.Printf("Handling getERC721Transfers tool call with params: %v", request.Params.Arguments)

		arguments := request.Params.Arguments
		chainID, ok := arguments["chainID"].(string)
		if !ok {
			log.Printf("Error: Invalid chainID argument")
			return nil, fmt.Errorf("invalid chainID argument")
		}
		address, ok := arguments["address"].(string)
		if !ok {
			log.Printf("Error: Invalid address argument")
			return nil, fmt.Errorf("invalid address argument")
		}

		contractAddress, _ := arguments["contractAddress"].(string)
		startBlock, _ := arguments["startBlock"].(string)
		endBlock, _ := arguments["endBlock"].(string)
		page, _ := arguments["page"].(string)
		offset, _ := arguments["offset"].(string)

		result, err := blockchainTool.GetERC721Transfers(chainID, address, contractAddress, startBlock, endBlock, page, offset)
		if err != nil {
			log.Printf("Error getting ERC721 transfers: %v", err)
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					mcp.TextContent{
						Type: "text",
						Text: fmt.Sprintf("Error getting ERC721 transfers: %v", err),
					},
				},
				IsError: true,
			}, nil
		}

		log.Printf("Successfully retrieved ERC721 transfers")
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: result,
				},
			},
		}, nil
	}
}

// HandleGetBlockRewards handles the getBlockRewards tool call
func HandleGetBlockRewards(blockchainTool *BlockchainTool) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		log.Printf("Handling getBlockRewards tool call with params: %v", request.Params.Arguments)

		arguments := request.Params.Arguments
		chainID, ok := arguments["chainID"].(string)
		if !ok {
			log.Printf("Error: Invalid chainID argument")
			return nil, fmt.Errorf("invalid chainID argument")
		}
		blockNumber, ok := arguments["blockNumber"].(string)
		if !ok {
			log.Printf("Error: Invalid blockNumber argument")
			return nil, fmt.Errorf("invalid blockNumber argument")
		}

		result, err := blockchainTool.GetBlockRewards(chainID, blockNumber)
		if err != nil {
			log.Printf("Error getting block rewards: %v", err)
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					mcp.TextContent{
						Type: "text",
						Text: fmt.Sprintf("Error getting block rewards: %v", err),
					},
				},
				IsError: true,
			}, nil
		}

		log.Printf("Successfully retrieved block rewards")
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: result,
				},
			},
		}, nil
	}
}

// HandleGetContractSourceCode handles the getContractSourceCode tool call
func HandleGetContractSourceCode(blockchainTool *BlockchainTool) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		log.Printf("Handling getContractSourceCode tool call with params: %v", request.Params.Arguments)

		arguments := request.Params.Arguments
		chainID, ok := arguments["chainID"].(string)
		if !ok {
			log.Printf("Error: Invalid chainID argument")
			return nil, fmt.Errorf("invalid chainID argument")
		}
		contractAddress, ok := arguments["contractAddress"].(string)
		if !ok {
			log.Printf("Error: Invalid contractAddress argument")
			return nil, fmt.Errorf("invalid contractAddress argument")
		}

		result, err := blockchainTool.GetContractSourceCode(chainID, contractAddress)
		if err != nil {
			log.Printf("Error getting contract source code: %v", err)
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					mcp.TextContent{
						Type: "text",
						Text: fmt.Sprintf("Error getting contract source code: %v", err),
					},
				},
				IsError: true,
			}, nil
		}

		log.Printf("Successfully retrieved contract source code")
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: result,
				},
			},
		}, nil
	}
}

// HandleGetGasOracle handles the getGasOracle tool call
func HandleGetGasOracle(blockchainTool *BlockchainTool) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		log.Printf("Handling getGasOracle tool call with params: %v", request.Params.Arguments)

		arguments := request.Params.Arguments
		chainID, ok := arguments["chainID"].(string)
		if !ok {
			log.Printf("Error: Invalid chainID argument")
			return nil, fmt.Errorf("invalid chainID argument")
		}

		result, err := blockchainTool.GetGasOracle(chainID)
		if err != nil {
			log.Printf("Error getting gas oracle: %v", err)
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					mcp.TextContent{
						Type: "text",
						Text: fmt.Sprintf("Error getting gas oracle: %v", err),
					},
				},
				IsError: true,
			}, nil
		}

		log.Printf("Successfully retrieved gas oracle")
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: result,
				},
			},
		}, nil
	}
}

// HandleExecuteContractMethod handles the executeContractMethod tool call
func HandleExecuteContractMethod(blockchainTool *BlockchainTool) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		log.Printf("Handling executeContractMethod tool call with params: %v", request.Params.Arguments)

		arguments := request.Params.Arguments
		chainID, ok := arguments["chainID"].(string)
		if !ok {
			log.Printf("Error: Invalid chainID argument")
			return nil, fmt.Errorf("invalid chainID argument")
		}
		contractAddress, ok := arguments["contractAddress"].(string)
		if !ok {
			log.Printf("Error: Invalid contractAddress argument")
			return nil, fmt.Errorf("invalid contractAddress argument")
		}
		methodABI, ok := arguments["methodABI"].(string)
		if !ok {
			log.Printf("Error: Invalid methodABI argument")
			return nil, fmt.Errorf("invalid methodABI argument")
		}

		methodParams, _ := arguments["methodParams"].(string)

		result, err := blockchainTool.ExecuteContractMethod(chainID, contractAddress, methodABI, methodParams)
		if err != nil {
			log.Printf("Error executing contract method: %v", err)
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					mcp.TextContent{
						Type: "text",
						Text: fmt.Sprintf("Error executing contract method: %v", err),
					},
				},
				IsError: true,
			}, nil
		}

		log.Printf("Successfully executed contract method")
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: result,
				},
			},
		}, nil
	}
}

// HandleGetTokenDetails handles the getTokenDetails tool call
func HandleGetTokenDetails(blockchainTool *BlockchainTool) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		log.Printf("Handling getTokenDetails tool call with params: %v", request.Params.Arguments)

		arguments := request.Params.Arguments
		chainID, ok := arguments["chainID"].(string)
		if !ok {
			log.Printf("Error: Invalid chainID argument")
			return nil, fmt.Errorf("invalid chainID argument")
		}
		contractAddress, ok := arguments["contractAddress"].(string)
		if !ok {
			log.Printf("Error: Invalid contractAddress argument")
			return nil, fmt.Errorf("invalid contractAddress argument")
		}

		result, err := blockchainTool.GetTokenDetails(chainID, contractAddress)
		if err != nil {
			log.Printf("Error getting token details: %v", err)
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					mcp.TextContent{
						Type: "text",
						Text: fmt.Sprintf("Error getting token details: %v", err),
					},
				},
				IsError: true,
			}, nil
		}

		log.Printf("Successfully retrieved token details")
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: result,
				},
			},
		}, nil
	}
}

// RegisterTools registers the MCP tools for blockchain data querying
func RegisterTools(configPath string) ([]mcp.Tool, map[string]server.ToolHandlerFunc) {
	log.Printf("Registering MCP tools with config: %s", configPath)
	blockchainTool, err := NewBlockchainTool(configPath)
	if err != nil {
		log.Fatalf("Error creating BlockchainTool: %v", err)
	}

	tools := []mcp.Tool{
		mcp.NewTool("getAccountBalance",
			mcp.WithDescription("Get the balance of an account on a specific blockchain"),
			mcp.WithString("chainID",
				mcp.Description("The chain ID (e.g., 1 for Ethereum)"),
				mcp.Required(),
			),
			mcp.WithString("address",
				mcp.Description("The account address"),
				mcp.Required(),
			),
		),
		mcp.NewTool("getTokenBalance",
			mcp.WithDescription("Get the token balance of an account on a specific blockchain"),
			mcp.WithString("chainID",
				mcp.Description("The chain ID (e.g., 1 for Ethereum)"),
				mcp.Required(),
			),
			mcp.WithString("contractAddress",
				mcp.Description("The token contract address"),
				mcp.Required(),
			),
			mcp.WithString("address",
				mcp.Description("The account address"),
				mcp.Required(),
			),
		),
		mcp.NewTool("getTransactionByHash",
			mcp.WithDescription("Get transaction details by hash"),
			mcp.WithString("chainID",
				mcp.Description("The chain ID (e.g., 1 for Ethereum)"),
				mcp.Required(),
			),
			mcp.WithString("txHash",
				mcp.Description("The transaction hash"),
				mcp.Required(),
			),
		),
		mcp.NewTool("getBlockByNumber",
			mcp.WithDescription("Get block information by block number"),
			mcp.WithString("chainID",
				mcp.Description("The chain ID (e.g., 1 for Ethereum)"),
				mcp.Required(),
			),
			mcp.WithString("blockNumber",
				mcp.Description("The block number (or 'latest')"),
				mcp.Required(),
			),
		),
		mcp.NewTool("getContractABI",
			mcp.WithDescription("Get the ABI for a verified contract"),
			mcp.WithString("chainID",
				mcp.Description("The chain ID (e.g., 1 for Ethereum)"),
				mcp.Required(),
			),
			mcp.WithString("contractAddress",
				mcp.Description("The contract address"),
				mcp.Required(),
			),
		),
		mcp.NewTool("getTransactionsByAddress",
			mcp.WithDescription("Get list of transactions by address"),
			mcp.WithString("chainID",
				mcp.Description("The chain ID (e.g., 1 for Ethereum)"),
				mcp.Required(),
			),
			mcp.WithString("address",
				mcp.Description("The account address"),
				mcp.Required(),
			),
			mcp.WithString("startBlock",
				mcp.Description("Starting block number"),
			),
			mcp.WithString("endBlock",
				mcp.Description("Ending block number"),
			),
			mcp.WithString("page",
				mcp.Description("Page number"),
			),
			mcp.WithString("offset",
				mcp.Description("Number of records to return"),
			),
		),
		mcp.NewTool("getInternalTransactionsByAddress",
			mcp.WithDescription("Get list of internal transactions by address"),
			mcp.WithString("chainID",
				mcp.Description("The chain ID (e.g., 1 for Ethereum)"),
				mcp.Required(),
			),
			mcp.WithString("address",
				mcp.Description("The account address"),
				mcp.Required(),
			),
			mcp.WithString("startBlock",
				mcp.Description("Starting block number"),
			),
			mcp.WithString("endBlock",
				mcp.Description("Ending block number"),
			),
			mcp.WithString("page",
				mcp.Description("Page number"),
			),
			mcp.WithString("offset",
				mcp.Description("Number of records to return"),
			),
		),
		mcp.NewTool("getTokenTransfersByAddress",
			mcp.WithDescription("Get list of token transfers by address"),
			mcp.WithString("chainID",
				mcp.Description("The chain ID (e.g., 1 for Ethereum)"),
				mcp.Required(),
			),
			mcp.WithString("address",
				mcp.Description("The account address"),
				mcp.Required(),
			),
			mcp.WithString("contractAddress",
				mcp.Description("The token contract address"),
			),
			mcp.WithString("startBlock",
				mcp.Description("Starting block number"),
			),
			mcp.WithString("endBlock",
				mcp.Description("Ending block number"),
			),
			mcp.WithString("page",
				mcp.Description("Page number"),
			),
			mcp.WithString("offset",
				mcp.Description("Number of records to return"),
			),
		),
		mcp.NewTool("getERC721Transfers",
			mcp.WithDescription("Get list of ERC721 token transfers by address"),
			mcp.WithString("chainID",
				mcp.Description("The chain ID (e.g., 1 for Ethereum)"),
				mcp.Required(),
			),
			mcp.WithString("address",
				mcp.Description("The account address"),
				mcp.Required(),
			),
			mcp.WithString("contractAddress",
				mcp.Description("The token contract address"),
			),
			mcp.WithString("startBlock",
				mcp.Description("Starting block number"),
			),
			mcp.WithString("endBlock",
				mcp.Description("Ending block number"),
			),
			mcp.WithString("page",
				mcp.Description("Page number"),
			),
			mcp.WithString("offset",
				mcp.Description("Number of records to return"),
			),
		),
		mcp.NewTool("getBlockRewards",
			mcp.WithDescription("Get block rewards by block number"),
			mcp.WithString("chainID",
				mcp.Description("The chain ID (e.g., 1 for Ethereum)"),
				mcp.Required(),
			),
			mcp.WithString("blockNumber",
				mcp.Description("The block number"),
				mcp.Required(),
			),
		),
		mcp.NewTool("getContractSourceCode",
			mcp.WithDescription("Get the source code of a verified contract"),
			mcp.WithString("chainID",
				mcp.Description("The chain ID (e.g., 1 for Ethereum)"),
				mcp.Required(),
			),
			mcp.WithString("contractAddress",
				mcp.Description("The contract address"),
				mcp.Required(),
			),
		),
		mcp.NewTool("getGasOracle",
			mcp.WithDescription("Get current gas price oracle output"),
			mcp.WithString("chainID",
				mcp.Description("The chain ID (e.g., 1 for Ethereum)"),
				mcp.Required(),
			),
		),
		mcp.NewTool("executeContractMethod",
			mcp.WithDescription("Execute a read contract function"),
			mcp.WithString("chainID",
				mcp.Description("The chain ID (e.g., 1 for Ethereum)"),
				mcp.Required(),
			),
			mcp.WithString("contractAddress",
				mcp.Description("The contract address"),
				mcp.Required(),
			),
			mcp.WithString("methodABI",
				mcp.Description("The ABI of the contract method"),
				mcp.Required(),
			),
			mcp.WithString("methodParams",
				mcp.Description("Comma-separated parameter values for the method"),
			),
		),
		mcp.NewTool("getTokenDetails",
			mcp.WithDescription("Get comprehensive token information"),
			mcp.WithString("chainID",
				mcp.Description("The chain ID (e.g., 1 for Ethereum)"),
				mcp.Required(),
			),
			mcp.WithString("contractAddress",
				mcp.Description("The token contract address"),
				mcp.Required(),
			),
		),
	}

	handlers := map[string]server.ToolHandlerFunc{
		"getAccountBalance":                HandleGetAccountBalance(blockchainTool),
		"getTokenBalance":                  HandleGetTokenBalance(blockchainTool),
		"getTransactionByHash":             HandleGetTransactionByHash(blockchainTool),
		"getBlockByNumber":                 HandleGetBlockByNumber(blockchainTool),
		"getContractABI":                   HandleGetContractABI(blockchainTool),
		"getTransactionsByAddress":         HandleGetTransactionsByAddress(blockchainTool),
		"getInternalTransactionsByAddress": HandleGetInternalTransactionsByAddress(blockchainTool),
		"getTokenTransfersByAddress":       HandleGetTokenTransfersByAddress(blockchainTool),
		"getERC721Transfers":               HandleGetERC721Transfers(blockchainTool),
		"getBlockRewards":                  HandleGetBlockRewards(blockchainTool),
		"getContractSourceCode":            HandleGetContractSourceCode(blockchainTool),
		"getGasOracle":                     HandleGetGasOracle(blockchainTool),
		"executeContractMethod":            HandleExecuteContractMethod(blockchainTool),
		"getTokenDetails":                  HandleGetTokenDetails(blockchainTool),
	}

	log.Printf("Registered %d tools", len(tools))
	return tools, handlers
}
