<p align="center">
  <img src="https://aniliststream.edmit.in/logo.png" alt="AnilistStream Logo" width="120" />
</p>

<h1 align="center">AnilistStream</h1>

<p align="center">
  A self-hostable Stremio addon for HTTP anime streaming with AniList integration for metadata and watch progress syncing.
</p>

<p align="center">
  <img src="https://img.shields.io/github/stars/Saadiq8149/AnilistStream?style=for-the-badge&logo=github" alt="GitHub Stars" />
  <img src="https://img.shields.io/github/license/Saadiq8149/AnilistStream?style=for-the-badge" alt="License" />
  <img src="https://img.shields.io/docker/pulls/12345saadiq/aniliststream?style=for-the-badge&logo=docker" alt="Docker Pulls" />
  <img src="https://img.shields.io/github/go-mod/go-version/Saadiq8149/AnilistStream?style=for-the-badge&logo=go" alt="Go Version" />
  <img src="https://img.shields.io/badge/Stremio-Addon-blue?style=for-the-badge&logo=stremio" alt="Stremio Addon" />
</p>

<p align="center">
  <a href="https://aniliststream.edmit.in">
    <img src="https://img.shields.io/badge/⚡%20Install%20Now-AnilistStream-6441A5?style=for-the-badge&logo=stremio&logoColor=white" alt="Install AnilistStream" />
  </a>
</p>

---

## Overview

**AnilistStream** is a [Stremio](https://www.stremio.com/) addon written in Go that provides HTTP-based anime streams — no torrents, no P2P. It integrates tightly with [AniList](https://anilist.co/) for rich metadata and optional watch progress syncing via OAuth login.

Originally ported from a Node.js implementation, the Go version is built around a **modular provider architecture**, making it straightforward to add or swap metadata and stream sources. It ships with a built-in static web UI for configuration and addon installation, and is fully self-hostable via Docker or from source.

---

## Features

- 🎬 **HTTP-based streams** — direct video streams, no torrents or P2P required
- 📋 **AniList integration** — rich metadata including titles, descriptions, cover art, ratings, and airing status
- 🔄 **Watch progress syncing** — log in with AniList OAuth to automatically track and sync watched episodes
- 🧩 **Modular provider architecture** — easily extend or replace metadata and stream providers
- 🌐 **Built-in web UI** — configure your addon and generate an install link directly from the browser
- 🐳 **Docker support** — single-container deployment with a public image on Docker Hub
- ☁️ **Public instance available** — use it immediately without self-hosting
- ⚙️ **Self-hostable** — full control over your own instance

---

## Install (Public Instance)

The easiest way to get started is to use the hosted public instance.

<p align="center">
  <a href="https://aniliststream.edmit.in">
    <img src="https://img.shields.io/badge/Open%20Configuration%20Page-aniliststream.edmit.in-6441A5?style=for-the-badge&logo=stremio&logoColor=white" />
  </a>
</p>

1. Visit **[https://aniliststream.edmit.in](https://aniliststream.edmit.in)**
2. Optionally log in with your **AniList account** to enable watch progress syncing
3. Click **Install Addon** to add it to Stremio

> **Note:** The public instance is community-maintained. For guaranteed uptime or custom configuration, consider self-hosting.

---

## Self Hosting

### Docker

Pull and run the latest image from Docker Hub:

```bash
docker pull 12345saadiq/aniliststream:latest
```

```bash
docker run -d \
  --name aniliststream \
  -p 8080:8080 \
  12345saadiq/aniliststream:latest
```

The web UI and addon endpoint will be available at `http://localhost:8080`.

---

### Docker Compose

Create a `docker-compose.yml` file:

```yaml
version: "3.8"

services:
  aniliststream:
    image: 12345saadiq/aniliststream:latest
    container_name: aniliststream
    ports:
      - "8080:8080"
    environment:
      - PORT=8080
      # Optional: set your AniList OAuth client credentials
      # - ANILIST_CLIENT_ID=your_client_id
      # - ANILIST_CLIENT_SECRET=your_client_secret
      # - ANILIST_REDIRECT_URI=http://localhost:8080/callback
    restart: unless-stopped
```

Then run:

```bash
docker compose up -d
```

---

### Run from Source

**Prerequisites:**
- [Go 1.21+](https://go.dev/dl/)
- Git

**Steps:**

```bash
# Clone the repository
git clone https://github.com/Saadiq8149/AnilistStream.git
cd AnilistStream

# Install dependencies
go mod download

# Run the server
go run .
```

The server will start on `http://localhost:8080` by default.

To build a binary:

```bash
go build -o aniliststream .
./aniliststream
```

---

## Configuration

Configuration is handled via **environment variables**. You can set these in your shell, a `.env` file, or your Docker Compose configuration.

| Variable | Default | Description |
|---|---|---|
| `PORT` | `8080` | Port the server listens on |
| `ANILIST_CLIENT_ID` | — | AniList OAuth application client ID |
| `ANILIST_CLIENT_SECRET` | — | AniList OAuth application client secret |
| `ANILIST_REDIRECT_URI` | — | OAuth redirect URI (must match your AniList app settings) |

### AniList OAuth Setup

To enable watch progress syncing, you need to register an AniList API application:

1. Go to [https://anilist.co/settings/developer](https://anilist.co/settings/developer)
2. Click **Create new client**
3. Set the **Redirect URL** to your instance's callback URL (e.g. `http://localhost:8080/callback`)
4. Copy the **Client ID** and **Client Secret** into your environment variables

Users can then log in from the web UI to authorize progress syncing.

---

## Architecture

AnilistStream uses a **modular provider architecture** that cleanly separates concerns:

```
┌──────────────────────────────────────────────────────┐
│                   Stremio Client                     │
└────────────────────────┬─────────────────────────────┘
                         │  HTTP (addon protocol)
┌────────────────────────▼─────────────────────────────┐
│               AnilistStream Server (Go)               │
│                                                      │
│  ┌─────────────┐  ┌──────────────┐  ┌─────────────┐ │
│  │   Catalog   │  │     Meta     │  │   Stream    │ │
│  │  Handler   │  │   Handler   │  │   Handler  │ │
│  └──────┬──────┘  └──────┬───────┘  └──────┬────┘ │
│         │                │                  │       │
│  ┌──────▼──────────────────▼────────────────▼────┐  │
│  │              Provider Interface               │  │
│  └──────┬──────────────────┬────────────────┬────┘  │
│         │                  │                │       │
│  ┌──────▼──────┐  ┌────────▼───────┐  ┌────▼─────┐ │
│  │   AniList   │  │   AllAnime     │  │AllAnime  │ │
│  │  (metadata) │  │  (metadata)    │  │(streams) │ │
│  └─────────────┘  └────────────────┘  └──────────┘ │
└──────────────────────────────────────────────────────┘
```

**Provider types:**

- **Metadata Providers** — Supply catalog and episode metadata (AniList, AllAnime)
- **Stream Providers** — Return playable HTTP stream URLs for a given episode (AllAnime)

New providers can be implemented by satisfying the provider interface, with no changes required to the core addon logic.

---

## Project Structure

```
AnilistStream/
├── main.go                 # Entry point, server setup
├── go.mod / go.sum         # Go module files
├── handlers/               # Stremio addon endpoint handlers
│   ├── catalog.go          # /catalog endpoint
│   ├── meta.go             # /meta endpoint
│   └── stream.go           # /stream endpoint
├── providers/              # Provider interface + implementations
│   ├── provider.go         # Shared provider interface
│   ├── anilist/            # AniList metadata provider
│   └── allanime/           # AllAnime metadata + stream provider
├── anilist/                # AniList OAuth + API client
├── web/                    # Static web UI (install + configuration)
│   └── static/
└── config/                 # Configuration loading
```

---

## Development

### Running Tests

```bash
go test ./...
```

### Building for Multiple Platforms

```bash
# Linux (amd64)
GOOS=linux GOARCH=amd64 go build -o aniliststream-linux-amd64 .

# macOS (arm64)
GOOS=darwin GOARCH=arm64 go build -o aniliststream-darwin-arm64 .

# Windows (amd64)
GOOS=windows GOARCH=amd64 go build -o aniliststream-windows-amd64.exe .
```

### Building the Docker Image

```bash
docker build -t aniliststream .
```

### Adding a New Provider

1. Create a new package under `providers/yourprovider/`
2. Implement the `MetadataProvider` and/or `StreamProvider` interface defined in `providers/provider.go`
3. Register the provider in `main.go`

---

## Contributing

Contributions are welcome! Please follow these steps:

1. **Fork** the repository
2. **Create** a feature branch: `git checkout -b feat/your-feature`
3. **Commit** your changes: `git commit -m "feat: add your feature"`
4. **Push** to your fork: `git push origin feat/your-feature`
5. **Open** a Pull Request against `main`

Please open an issue first for major changes or new providers so the approach can be discussed before implementation.

---

## Disclaimer

AnilistStream is an independent open-source project and is **not affiliated with, endorsed by, or connected to Stremio, AniList, or any content provider** used as a stream source.

This addon indexes and links to streams hosted by third-party services. The developers of AnilistStream do not host, store, or distribute any video content. Users are responsible for ensuring that their use of this addon complies with the laws and regulations applicable in their jurisdiction.

The public instance is provided on a best-effort basis with no guarantees of uptime, availability, or continued operation.
