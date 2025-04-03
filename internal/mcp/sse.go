package mcp

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mark3labs/mcp-go/server"
)

// CustomSSEServer extends the SSEServer with heartbeat functionality
type CustomSSEServer struct {
	*server.SSEServer
	heartbeatInterval time.Duration
	srv               *http.Server
}

// NewCustomSSEServer creates a new CustomSSEServer
func NewCustomSSEServer(mcpServer *server.MCPServer, opts ...server.SSEOption) *CustomSSEServer {
	return &CustomSSEServer{
		SSEServer:         server.NewSSEServer(mcpServer, opts...),
		heartbeatInterval: 25 * time.Second, // Default heartbeat interval
	}
}

// WithHeartbeatInterval sets the heartbeat interval
func (s *CustomSSEServer) WithHeartbeatInterval(interval time.Duration) *CustomSSEServer {
	s.heartbeatInterval = interval
	return s
}

// HandleSSE overrides the original handleSSE to add heartbeat functionality
func (s *CustomSSEServer) handleSSE(w http.ResponseWriter, r *http.Request) {
	// Set up headers as in the original
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	// Access the original SSEServer method using a different approach
	// by making a standard call to the original ServeHTTP but redirecting it
	// to a custom handler that we control
	done := make(chan struct{})
	defer close(done)

	// Create a heartbeat ticker
	ticker := time.NewTicker(s.heartbeatInterval)
	defer ticker.Stop()

	// Keep the connection alive with comment events
	go func() {
		for {
			select {
			case <-ticker.C:
				// Send a comment as heartbeat
				// Comments are ignored by EventSource but keep the connection alive
				fmt.Fprintf(w, ": heartbeat %v\n\n", time.Now())
				flusher.Flush()
			case <-done:
				return
			case <-r.Context().Done():
				return
			}
		}
	}()

	// Call the original SSEServer's ServeHTTP
	s.SSEServer.ServeHTTP(w, r)
}

// ServeHTTP implements the http.Handler interface
func (s *CustomSSEServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	// For SSE endpoint, use our custom handler
	ssePath := s.CompleteSsePath()
	if ssePath != "" && path == ssePath {
		s.handleSSE(w, r)
		return
	}

	// For all other endpoints, delegate to the original
	s.SSEServer.ServeHTTP(w, r)
}

// Start begins serving SSE connections on the specified address
func (s *CustomSSEServer) Start(addr string) error {
	// Create a new http.Server
	s.srv = &http.Server{
		Addr:    addr,
		Handler: s,
	}

	return s.srv.ListenAndServe()
}
