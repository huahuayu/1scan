.PHONY: build build-1scan build-1scanmcp run run-1scan run-1scanmcp install stop stop-1scan stop-1scanmcp clean clean-port help

# Default target
.DEFAULT_GOAL := help

# Variables
BINDIR := bin
GOFLAGS := -ldflags="-s -w"

help:
	@echo "Available commands:"
	@echo "  make build           - Build both 1scan and 1scanmcp binaries"
	@echo "  make build-1scan     - Build only 1scan binary"
	@echo "  make build-1scanmcp  - Build only 1scanmcp binary"
	@echo "  make run-1scan       - Start 1scan in debug mode"
	@echo "  make run-1scanmcp    - Start 1scanmcp in debug mode"
	@echo "  make install         - Install both binaries to /usr/local/bin"
	@echo "  make stop-1scan      - Stop running 1scan process"
	@echo "  make stop-1scanmcp   - Stop running 1scanmcp process"
	@echo "  make stop            - Stop all running processes"
	@echo "  make clean           - Remove all built binaries"
	@echo "  make clean-port      - Free up port 8080 if occupied"
	@echo "  make help            - Show this help message"

$(BINDIR):
	@mkdir -p $(BINDIR)

build: build-1scan build-1scanmcp

build-1scan: $(BINDIR)
	@echo "Building 1scan..."
	@go build $(GOFLAGS) -o $(BINDIR)/1scan cmd/1scan/main.go
	@echo "1scan binary built at $(BINDIR)/1scan"

build-1scanmcp: $(BINDIR)
	@echo "Building 1scanmcp..."
	@go build $(GOFLAGS) -o $(BINDIR)/1scanmcp cmd/1scanmcp/main.go
	@echo "1scanmcp binary built at $(BINDIR)/1scanmcp"

clean-port:
	@echo "Checking port 8080..."
	@if lsof -i :8080 >/dev/null 2>&1; then \
		echo "Port 8080 is in use. Attempting to free it..."; \
		lsof -ti :8080 | xargs kill -9; \
		echo "Port 8080 has been freed."; \
	else \
		echo "Port 8080 is available."; \
	fi

run-1scan: clean-port
	@echo "Starting 1scan..."
	@if [ -f "$(BINDIR)/1scan" ]; then \
		$(BINDIR)/1scan --log-level=debug; \
	else \
		echo "Building 1scan first..."; \
		$(MAKE) build-1scan; \
		$(BINDIR)/1scan --log-level=debug; \
	fi

run-1scanmcp:
	@echo "Starting 1scanmcp..."
	@if [ -f "$(BINDIR)/1scanmcp" ]; then \
		$(BINDIR)/1scanmcp; \
	else \
		echo "Building 1scanmcp first..."; \
		$(MAKE) build-1scanmcp; \
		$(BINDIR)/1scanmcp; \
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

stop-1scanmcp:
	@echo "Stopping 1scanmcp..."
	@if pgrep -f "1scanmcp " > /dev/null || pgrep -f "bin/1scanmcp" > /dev/null; then \
		pkill -f "1scanmcp " || pkill -f "bin/1scanmcp" && echo "1scanmcp stopped"; \
	else \
		echo "No running 1scanmcp process found"; \
	fi

stop: stop-1scan stop-1scanmcp

install: build
	@echo "Installing 1scan and 1scanmcp to /usr/local/bin..."
	@cp $(BINDIR)/1scan /usr/local/bin/
	@cp $(BINDIR)/1scanmcp /usr/local/bin/
	@echo "Installation complete"

clean:
	@echo "Cleaning up binaries..."
	@rm -rf $(BINDIR)
	@echo "Cleanup complete" 