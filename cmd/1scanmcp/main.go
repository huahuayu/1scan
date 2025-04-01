package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/huahuayu/1scan/internal/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Parse command line flags
	port := flag.String("port", "3000", "Port to run the MCP server on")
	path := flag.String("path", "/mcp", "Path for the MCP server endpoint")
	configPath := flag.String("config", "config.json", "Path to the config file")
	transport := flag.String("transport", "sse", "Transport type (sse or stdio)")
	flag.Parse()

	// Ensure path starts with a slash
	if !strings.HasPrefix(*path, "/") {
		*path = "/" + *path
	}

	// Check if config file exists
	if _, err := os.Stat(*configPath); os.IsNotExist(err) {
		log.Fatalf("Config file not found: %s", *configPath)
	}

	// Create MCP server
	mcpServer := server.NewMCPServer(
		"1scan-blockchain-tools",
		"1.0.0",
		server.WithLogging(),
	)

	// Register tools and handlers
	tools, handlers := mcp.RegisterTools(*configPath)

	// Add tools and handlers to the server
	for _, tool := range tools {
		mcpServer.AddTool(tool, handlers[tool.Name])
	}

	// Start the server based on the transport type
	if *transport == "sse" {
		// Configure the SSE server with proper endpoints
		serverURL := fmt.Sprintf("http://localhost:%s", *port)
		log.Printf("Setting up SSE server with base URL: %s and path: %s", serverURL, *path)

		sseServer := server.NewSSEServer(
			mcpServer,
			server.WithBaseURL(serverURL),
			server.WithBasePath(*path),
		)

		go func() {
			log.Printf("Starting MCP server with SSE transport on %s%s", serverURL, *path)
			log.Printf("SSE endpoint: %s", sseServer.CompleteSseEndpoint())
			log.Printf("Message endpoint: %s", sseServer.CompleteMessageEndpoint())
			log.Printf("MCP server using config file: %s", *configPath)

			addr := fmt.Sprintf(":%s", *port)
			if err := sseServer.Start(addr); err != nil {
				log.Fatalf("Failed to start MCP server: %v", err)
			}
		}()

		// Handle graceful shutdown
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		fmt.Println("\nShutting down MCP server...")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := sseServer.Shutdown(shutdownCtx); err != nil {
			log.Fatalf("Error shutting down server: %v", err)
		}
	} else if *transport == "stdio" {
		log.Printf("Starting MCP server with stdio transport")
		log.Printf("MCP server using config file: %s", *configPath)

		if err := server.ServeStdio(mcpServer); err != nil {
			log.Fatalf("Failed to start MCP server: %v", err)
		}
	} else {
		log.Fatalf("Unsupported transport type: %s (must be 'sse' or 'stdio')", *transport)
	}

	fmt.Println("MCP server has been shut down")
}
