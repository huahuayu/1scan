package handler

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/huahuayu/1scan/config"
	"github.com/huahuayu/1scan/internal/loadbalancer"
	"github.com/huahuayu/1scan/internal/logger"
)

type Handler struct {
	config       *config.Manager
	loadBalancer *loadbalancer.LoadBalancer
}

func NewHandler(config *config.Manager, lb *loadbalancer.LoadBalancer) *Handler {
	return &Handler{
		config:       config,
		loadBalancer: lb,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger.Debug("Received request: %s %s", r.Method, r.URL.String())

	// Extract chainID from path
	pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(pathParts) < 1 {
		logger.Warn("Missing chainID in request path")
		http.Error(w, "Missing chainID", http.StatusBadRequest)
		return
	}

	chainID, err := strconv.Atoi(pathParts[0])
	if err != nil {
		logger.Warn("Invalid chainID: %s", pathParts[0])
		http.Error(w, "Invalid chainID", http.StatusBadRequest)
		return
	}

	// Get chain config
	chainConfig, exists := h.config.GetChainConfig(chainID)
	if !exists {
		logger.Warn("Unsupported chain ID: %d", chainID)
		http.Error(w, "Chain not supported", http.StatusNotFound)
		return
	}

	// Check for user-provided API key
	userAPIKey := r.URL.Query().Get("apikey")

	var apiKey string
	if userAPIKey != "" {
		apiKey = userAPIKey
		logger.Debug("Using user-provided API key: %s", apiKey)
	} else {
		apiKey, err = h.loadBalancer.GetNextAPIKey(chainID)
		if err != nil {
			logger.Error("No available API keys for chain %d", chainID)
			http.Error(w, "No available API keys", http.StatusServiceUnavailable)
			return
		}
		logger.Debug("Using load-balanced API key: %s", apiKey)
	}

	// Forward the request
	targetURL := fmt.Sprintf("https://%s/api", chainConfig.Endpoint)
	proxyReq, err := h.createProxyRequest(r, targetURL, apiKey)
	if err != nil {
		logger.Error("Failed to create proxy request: %v", err)
		http.Error(w, "Error creating proxy request", http.StatusInternalServerError)
		return
	}

	logger.Debug("Forwarding request to: %s", proxyReq.URL.String())

	client := &http.Client{}
	resp, err := client.Do(proxyReq)
	if err != nil {
		logger.Error("Failed to forward request: %v", err)
		http.Error(w, "Error forwarding request", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	// Copy response headers
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	// Copy response status code
	w.WriteHeader(resp.StatusCode)

	// Copy and log response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Failed to read response body: %v", err)
		return
	}

	logger.Debug("Response status: %d, body: %s", resp.StatusCode, string(body))
	w.Write(body)
}

func (h *Handler) createProxyRequest(r *http.Request, targetURL, apiKey string) (*http.Request, error) {
	// Create new URL with existing query parameters
	queryParams := r.URL.Query()
	queryParams.Set("apikey", apiKey)

	proxyURL := fmt.Sprintf("%s?%s", targetURL, queryParams.Encode())

	// Create new request
	proxyReq, err := http.NewRequest(r.Method, proxyURL, r.Body)
	if err != nil {
		return nil, err
	}

	// Copy headers
	for key, values := range r.Header {
		for _, value := range values {
			proxyReq.Header.Add(key, value)
		}
	}

	return proxyReq, nil
}
