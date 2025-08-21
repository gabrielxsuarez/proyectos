# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a unified WebSocket client application for Chat Argentina (chatarg.com) written in Go. The project connects to WebSocket servers and handles real-time messaging with enhanced nick-based functionality for different modes of operation.

## Build & Run Commands

```bash
# Build the main application
go build main.go struct_envios.go struct_recibos.go

# Run the unified WebSocket client
go run main.go struct_envios.go struct_recibos.go

# Install dependencies
go mod download

# Update dependencies
go mod tidy
```

## Architecture

The unified codebase consists of:

- **main.go**: Main unified WebSocket client with enhanced nick processing
- **struct_envios.go**: Data structures for outgoing WebSocket messages
- **struct_recibos.go**: Data structures for incoming WebSocket messages
- **config.txt**: Configuration file containing server and room settings
- **puertos.txt**: List of ports for WebSocket testing functionality
- **ip.txt**: List of IP addresses for random selection

Key technical details:
- Uses gorilla/websocket library for WebSocket connections
- Implements keepalive mechanism with periodic ping messages (30s intervals)
- Handles graceful shutdown on interrupt signals
- Simulates browser headers (Origin, User-Agent) for compatibility
- Integrated port testing functionality for diagnostics

## Enhanced Nick Functionality

The application now supports multiple modes based on nick input:

### Normal Mode
- **Input**: `Pablo`
- **Behavior**: Connects to the server and room specified in config.txt

### Room Override Mode
- **Input**: `#Peru Pablo`
- **Behavior**: Connects as "Pablo" to the "#Peru" room, ignoring the room setting in config.txt

### Port Testing Mode
- **Input**: `test all port`
- **Behavior**: Tests all ports from puertos.txt against the server specified in config.txt

### Port Testing with Custom Server
- **Input**: `test all port wss.dalechatea.me`
- **Behavior**: Tests all ports from puertos.txt against the specified server

## Configuration Files

### config.txt Format
```
servidor: ws://wss.dalechatea.me:1245/
sala: argentina
```

### puertos.txt Format
```
21
25
53
80
443
1239
1242
1245
...
```

## Important Context

- The production server (`ws://wss.dalechatea.me:1245/`) is the default in config.txt
- Alternative WebSocket servers available:
  - `wss://wss.dalechatea.me:1242/` - WebSocket server with SSL
  - `ws://wss.dalechatea.me:1245/` - WebSocket server without SSL
- Backend server on port 1239 has additional security restrictions (403 Forbidden)
- Connection requires specific headers mimicking browser behavior
- The application handles both text and binary WebSocket messages
- All servers return a sessionid message upon successful connection
- Port testing includes both WSS and WS protocols for comprehensive coverage