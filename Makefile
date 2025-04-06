.PHONY: build build-1scan build-1scan-mcp run run-1scan run-1scan-mcp install stop stop-1scan stop-1scan-mcp clean help

# Default target
.DEFAULT_GOAL := help

# Variables
BINDIR := bin
GOFLAGS := -ldflags="-s -w"

help:
	@echo "Available commands:"
	@echo "  make build           - Build both 1scan and 1scan-mcp binaries"
	@echo "  make build-1scan     - Build only 1scan binary"
	@echo "  make build-1scan-mcp  - Build only 1scan-mcp binary"
	@echo "  make run-1scan       - Start 1scan in debug mode"
	@echo "  make run-1scan-mcp    - Start 1scan-mcp in debug mode"
	@echo "  make install         - Install both binaries to /usr/local/bin"
	@echo "  make stop-1scan      - Stop running 1scan process"
	@echo "  make stop-1scan-mcp   - Stop running 1scan-mcp process"
	@echo "  make stop            - Stop all running processes"
	@echo "  make clean           - Remove all built binaries"
	@echo "  make help            - Show this help message"

$(BINDIR):
	@mkdir -p $(BINDIR)

build: build-1scan build-1scan-mcp

build-1scan: $(BINDIR)
	@echo "Building 1scan..."
	@go build $(GOFLAGS) -o $(BINDIR)/1scan cmd/1scan/main.go
	@echo "1scan binary built at $(BINDIR)/1scan"

build-1scan-mcp: $(BINDIR)
	@echo "Building 1scan-mcp..."
	@go build $(GOFLAGS) -o $(BINDIR)/1scan-mcp cmd/1scan-mcp/main.go
	@echo "1scan-mcp binary built at $(BINDIR)/1scan-mcp"

run-1scan: clean-port
	@echo "Starting 1scan..."
	@if [ -f "$(BINDIR)/1scan" ]; then \
		$(BINDIR)/1scan --log-level=debug; \
	else \
		echo "Building 1scan first..."; \
		$(MAKE) build-1scan; \
		$(BINDIR)/1scan --log-level=debug; \
	fi

run-1scan-mcp:
	@echo "Starting 1scan-mcp..."
	@if [ -f "$(BINDIR)/1scan-mcp" ]; then \
		$(BINDIR)/1scan-mcp; \
	else \
		echo "Building 1scan-mcp first..."; \
		$(MAKE) build-1scan-mcp; \
		$(BINDIR)/1scan-mcp; \
	fi

run: run-1scan

stop-1scan:
	@echo "Stopping 1scan..."
	@if pgrep -f "1scan " > /dev/null || pgrep -f "bin/1scan" > /dev/null; then \
		pkill -f "1scan " || pkill -f "bin/1scan" && echo "1scan stopped"; \
	else \
		echo "No running 1scan process found"; \
	fi
	@$(MAKE) clean-port

stop-1scan-mcp:
	@echo "Stopping 1scan-mcp..."
	@if pgrep -f "1scan-mcp " > /dev/null || pgrep -f "bin/1scan-mcp" > /dev/null; then \
		pkill -f "1scan-mcp " || pkill -f "bin/1scan-mcp" && echo "1scan-mcp stopped"; \
	else \
		echo "No running 1scan-mcp process found"; \
	fi

stop: stop-1scan stop-1scan-mcp

install: build
	@echo "Installing 1scan and 1scan-mcp to /usr/local/bin..."
	@cp $(BINDIR)/1scan /usr/local/bin/
	@cp $(BINDIR)/1scan-mcp /usr/local/bin/
	@echo "Installation complete"

clean:
	@echo "Cleaning up binaries..."
	@rm -rf $(BINDIR)
	@echo "Cleanup complete" 