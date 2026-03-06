<p align="center">
  <img src="https://aniliststream.edmit.in/logo.png" alt="AnilistStream Logo" width="120" />
</p>

<h1 align="center">AnilistStream</h1>

<p align="center">
  A self-hostable Stremio addon for HTTP anime streaming with AniList integration for metadata and watch progress syncing.
</p>

<p align="center">
  <img src="https://img.shields.io/github/stars/Saadiq8149/AnilistStream?style=for-the-badge&logo=github" alt="GitHub Stars" />
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

**AnilistStream** is a [Stremio](https://www.stremio.com/) addon for streaming anime from multiple HTTP sources (no torrents, no P2P) while keeping your AniList watch list in sync. Originally ported from a Node.js implementation to Go, it features a modular provider architecture for easily adding or swapping metadata and stream sources, and can be fully self-hosted via Docker or from source.

Originally ported from a Node.js implementation, the Go version is built around a **modular provider architecture**, making it straightforward to add or swap metadata and stream sources. It ships with a built-in static web UI for configuration and addon installation, and is fully self-hostable via Docker or from source.

---
## Features

- 🎬 **HTTP-based streams** — direct video streams, no torrents or P2P required
- 🔄 **Watch progress syncing** — log in with AniList OAuth to automatically track and sync watched episodes
- 🧩 **Modular provider architecture** — easily extend or replace metadata and stream providers
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

Pull and run the latest image from Docker Hub.

Refer to the Configuration section below or copy `.env.example` to create your `.env` file.

```bash
docker pull 12345saadiq/aniliststream:latest && docker rm -f aniliststream 2>/dev/null || true && docker run -d --restart unless-stopped --name aniliststream -p 7000:7000 --env-file /.env 12345saadiq/aniliststream:latest
```

---

## Configuration

Configuration is handled via **environment variables**. You can set these in your `.env` file.

| Variable | Default | Description |
|---|---|---|
| `PORT` | `7000` | Port the server listens on |
| `SERVER_URL` | `http://127.0.0.1:7000` | Url the server runs on |
| `ANILIST_CLIENT_ID` | — | AniList OAuth application client ID |
| `METADATA_PROVIDER` |`ANILIST or ALL_ANIME` | Pick one, Anilist Better | 
| `SOURCE_PROVIDERS` |`ALL_ANIME` | Comma separated providers if multiple available | 



### AniList OAuth Setup

To enable watch progress syncing, you need to register an AniList API application:

1. Go to [https://anilist.co/settings/developer](https://anilist.co/settings/developer)
2. Click **Create new client**
3. Set the **Redirect URL** to your instance's callback URL (`{SERVER_URL}/configure`)
4. Copy the **Client ID** into your environment variables

Users can then log in from the web UI to authorize progress syncing.

---

## Roadmap / TODO

Planned improvements and upcoming features:

- [ ]  Redis caching layer for faster metadata and stream resolution
- [ ]  CI/CD pipeline with GitHub Actions (build, test, Docker publish)
- [ ]  Better background images for catalog and metadata views
- [ ]  Kitsu support
- [ ]  Episode Data (Titles, Dates and Thumbnails)
- [ ]  More Lists (Trending, Planning, Watching)
<img src="https://img.shields.io/github/issues/Saadiq8149/AnilistStream?style=for-the-badge" />

## Architecture

AnilistStream uses a **modular provider architecture** that cleanly separates concerns:
      
**Provider types:**

- **Metadata Providers** — Supply catalog and episode metadata (AniList, AllAnime)
- **Stream Providers** — Return playable HTTP stream URLs for a given episode (AllAnime)

New providers can be implemented by satisfying the provider interface, with no changes required to the core addon logic.

---

## Project Structure

```
AnilistStream/
├── main.go                     # Application entry point
│
├── internal/                   # Core application logic
│   ├── anilist/                # AniList OAuth + API client
│   ├── cache/                  # Caching utilities
│   ├── handlers/               # HTTP handlers and routing
│   │   └── routes.go
│   ├── metadata/               # Metadata provider implementations
│   ├── pages/                  # HTML page handlers (index, configure)
│   ├── streams/                # Stream provider implementations
│   ├── stremio/                # Stremio addon protocol logic
│   ├── types/                  # Shared data structures
│   └── util/                   # Utility helpers
│                  
└── public/                     # Static frontend asset
```

---

## Development

### Adding a New Provider

1. Create a new provider under `metadata/yourprovider/` or `streams/yourprovider`
2. Implement the `MetadataProvider` and/or `StreamProvider` interface defined in `metadata/provider.go` or `streams/provider.go`
3. Register the provider in `provider.go`

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
