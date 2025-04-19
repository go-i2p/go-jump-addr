# go-jump-addr

[![Go Version](https://img.shields.io/badge/go-1.23.5-blue.svg)](https://golang.org/doc/devel/release.html#go1.23)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

A library and server implementation for I2P "Jump" service functionality. This allows for human-readable domain names to be mapped to I2P destinations and provides search/discovery capabilities for I2P services.

## Table of Contents
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
  - [Running the Server](#running-the-server)
  - [API Endpoints](#api-endpoints)
- [Development](#development)
- [Contributing](#contributing)
- [License](#license)

## Features

- **Human-readable Hostnames**: Map memorable names to I2P destinations
- **Search Functionality**: Search across hostnames, descriptions, and tags
- **Metadata Extraction**: Automatically extracts metadata from I2P destinations
- **Sync Support**: Synchronize hostname data with other jump servers
- **Web Interface**: Clean, responsive web UI for managing entries
- **RESTful API**: Simple HTTP API for programmatic access
- **Tag System**: Organize entries with tags for better discovery
- **Validation**: Built-in validation for I2P addresses and hostnames

## Installation

Requires Go 1.23.5 or later.

```bash
go install github.com/go-i2p/go-jump-addr/jumpd@latest
```

Or clone and build manually:

```bash
git clone https://github.com/go-i2p/go-jump-addr.git
cd go-jump-addr
make build
```

## Usage

### Running the Server

1. Ensure I2P router is running with SAM enabled on port 7656

2. Start the jump server:

```bash
./jumpserver
```

The server will be accessible via I2P at the address shown in the startup logs.

### API Endpoints

- `GET /` - Homepage with hostname listing
- `GET /search` - Search interface
- `GET /add` - Add new hostname form
- `POST /add` - Submit new hostname
- `GET /all-hosts.txt` - Plain text list of all hostnames

### Example: Adding a Hostname

```bash
curl -X POST http://localhost:7654/add \
  -d "hostname=example.i2p" \
  -d "destination=BASE64_DEST..." \
  -d "type=service" \
  -d "name=Example Service" \
  -d "description=An example service" \
  -d "tags=example,demo"
```

## Development

Requirements:
- Go 1.23.5+
- Running I2P router with SAM enabled
- Make (optional, for build script)

Building from source:

```bash
make build
```

Testing:

```bash 
go test ./...
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/improvement`)
3. Make changes with accompanying tests
4. Run tests (`go test ./...`)
5. Commit changes (`git commit -am 'Add improvement'`)
6. Push to branch (`git push origin feature/improvement`)
7. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

Copyright (c) 2025 I2P For Go