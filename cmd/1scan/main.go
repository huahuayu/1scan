package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/huahuayu/1scan/config"
	"github.com/huahuayu/1scan/internal/handler"
	"github.com/huahuayu/1scan/internal/loadbalancer"
	"github.com/huahuayu/1scan/internal/logger"
)

func main() {
	// Command line flags
	configFile := flag.String("config", "config.json", "path to config file")
	logLevel := flag.String("log-level", "info", "logging level (debug, info, warn, error)")
	port := flag.String("port", "8080", "port to listen on")
	flag.Parse()

	// Set log level
	switch *logLevel {
	case "debug":
		logger.SetLevel(logger.DebugLevel)
	case "info":
		logger.SetLevel(logger.InfoLevel)
	case "warn":
		logger.SetLevel(logger.WarnLevel)
	case "error":
		logger.SetLevel(logger.ErrorLevel)
	default:
		log.Fatalf("Invalid log level: %s", *logLevel)
	}

	// Initialize config
	cfg := config.NewManager()
	if err := cfg.LoadConfig(*configFile); err != nil {
		logger.Error("Failed to load config: %v", err)
		return
	}
	logger.Info("Config loaded successfully")

	// Initialize load balancer
	lb := loadbalancer.NewLoadBalancer()

	// Update API keys for each chain
	for chainID, chainConfig := range cfg.GetAllChains() {
		keys := make([]string, 0, len(chainConfig.Keys))
		for k := range chainConfig.Keys {
			keys = append(keys, k)
		}
		lb.UpdateAPIKeys(chainID, keys)
		logger.Debug("Loaded API keys for chain %d: %v", chainID, keys)
	}

	// Create handler
	h := handler.NewHandler(cfg, lb)

	// Start server
	addr := ":" + *port
	logger.Info("Starting server on %s", addr)
	if err := http.ListenAndServe(addr, h); err != nil {
		logger.Error("Server failed: %v", err)
	}
}
