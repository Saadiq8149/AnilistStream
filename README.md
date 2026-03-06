<p align="center">
  <img src="https://aniliststream.edmit.in/logo.png" alt="AnilistStream Logo" width="120" />
</p>

<h1 align="center">AnilistStream</h1>

<p align="center">
  A self-hostable Stremio addon for HTTP anime streaming with AniList integration for metadata and watch progress syncing.
</p>

<p align="center">
  Works with the addon’s own catalogs and <a href="https://github.com/danamag/stremio-anime-kitsu">AnimeKitsu catalogs</a> :)
</p>

<p align="center">
  <img src="https://img.shields.io/github/stars/Saadiq8149/AnilistStream?style=for-the-badge&logo=github" alt="GitHub Stars" />
  <img src="https://img.shields.io/github/go-mod/go-version/Saadiq8149/AnilistStream?style=for-the-badge&logo=go" alt="Go Version" />
  <img src="https://img.shields.io/badge/Stremio-Addon-blue?style=for-the-badge&logo=stremio" alt="Stremio Addon" />
</p>

<p align="center">
  <a href="https://aniliststream.edmit.in">
    <img src="https://img.shields.io/badge/⚡%20Install%20Now-AnilistStream-6441A5?style=for-the-badge&logo=stremio&logoColor=white" alt="Install AnilistStream" />
  </a>
  <a href="https://hub.docker.com/r/12345saadiq/aniliststream">
    <img src="https://img.shields.io/badge/Docker%20Hub-12345saadiq%2Faniliststream-2496ED?style=for-the-badge&logo=docker&logoColor=white" alt="Docker Hub Repository" />
  </a>
</p>

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

| Variable | Format/Options | Description |
|---|---|---|
| `PORT` | `7000` | Port the server listens on |
| `SERVER_URL` | `http://127.0.0.1:7000` | Url the server runs on |
| `METADATA_PROVIDER` |`ANILIST or ALL_ANIME` | Pick one, Anilist Better | 
| `SOURCE_PROVIDERS` |`ALL_ANIME` | Comma separated providers if multiple available | 
| `ANILIST_CLIENT_ID` | — | AniList OAuth application client ID (Optional) |
| `IDS_MOE_API_KEY` | — | IdsMoe API KEY (Optional) | 



### AniList OAuth Setup

To enable watch progress syncing, you need to register an AniList API application:

1. Go to [https://anilist.co/settings/developer](https://anilist.co/settings/developer)
2. Click **Create new client**
3. Set the **Redirect URL** to your instance's callback URL (`{SERVER_URL}/configure`)
4. Copy the **Client ID** into your environment variables

Users can then log in from the web UI to authorize progress syncing.

### IdsMoe API Key

To get Kitsu Catalog support, you need to obtain an API key from IdsMoe:

1. Go to [https://ids.moe/](https://ids.moe/)
2. Sign up for an account and generate an API key
3. Copy the **API key** into your environment variables

---

## Roadmap / TODO

Planned improvements and upcoming features:

- [ ]  Redis caching layer for faster metadata and stream resolution
- [x]  CI/CD pipeline with GitHub Actions (build, test, Docker publish)
- [x]  Kitsu support
- [ ]  More Lists (Trending, Planning, Watching)
<img src="https://img.shields.io/github/issues/Saadiq8149/AnilistStream?style=for-the-badge" />

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

## Disclaimer

AnilistStream is an independent open-source project and is **not affiliated with, endorsed by, or connected to Stremio, AniList, or any content provider** used as a stream source.

This addon indexes and links to streams hosted by third-party services. The developers of AnilistStream do not host, store, or distribute any video content. Users are responsible for ensuring that their use of this addon complies with the laws and regulations applicable in their jurisdiction.

The public instance is provided on a best-effort basis with no guarantees of uptime, availability, or continued operation.
