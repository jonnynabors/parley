# Parley

A modern desktop HTTP/GraphQL client built with Go and Wails.

## Features

- Send HTTP requests (GET, POST, PUT, DELETE, PATCH)
- Execute GraphQL queries and mutations
- Save and load request collections
- Display formatted responses with timing information
- Support for custom headers and environment variables

## Development

### Prerequisites

- Go 1.21 or higher
- Node.js 16 or higher
- Wails CLI

### Setup

```bash
# Install dependencies
go mod download

# Run in development mode
wails dev

# Build for production
wails build
```

## Project Structure

```
parley/
├── main.go           # Application entry point
├── app.go            # Application logic
├── internal/
│   ├── http/         # HTTP request handling
│   ├── graphql/      # GraphQL client
│   ├── storage/      # Request/collection persistence
│   └── models/       # Data models
├── frontend/         # Web UI (React/TypeScript)
└── pkg/              # Public packages
```

## License

MIT
