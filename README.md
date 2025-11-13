# AniList Stremio Addon

A Stremio addon that integrates AniList with Stremio for seamless anime streaming and watch list synchronization.

## Features

- Search for anime across AniList
- View your AniList watch lists (Watching Now, Planning to Watch)
- Stream anime episodes with automatic watch status updates
- Automatic progress synchronization with AniList
- Multiple streaming sources with quality options
- Subtitle support

## Installation

Install this addon in Stremio by visiting the addon URL:

https://miraitv.stremio.edmit.in

## Setup

1. Visit the addon configuration page
2. Authorize with AniList by clicking the provided link
3. Copy your AniList access token
4. Paste the token into the addon configuration
5. Start streaming and syncing your watch list

## How It Works

### Search
Search for any anime and get results from AniList's extensive database.

### Watch Lists
- View anime you're currently watching
- View anime you're planning to watch
- Automatic synchronization when you start watching

### Streaming
- Multiple streaming sources available
- Auto-detection of episode availability
- Automatic watch status updates:
  - Planning -> Current (when you start watching)
  - Current -> Repeating (when you complete a series)
  - Automatically saves your episode progress

## Configuration

The addon requires an AniList access token for personalized features:

- Access token is used only for your watch list synchronization
- No data is stored on external servers
- Your credentials are secure

To get your access token:
1. Visit the AniList OAuth authorization page (link provided in addon settings)
2. Authorize the application
3. Copy the token from the response
4. Paste it into the addon configuration

## Technical Details

- Built with Node.js and Express
- Uses AniList GraphQL API
- Integrates with Stremio's addon protocol
- Sources anime from multiple providers

## To-Do / Future Features

Planned tasks and enhancements for upcoming development cycles:

- [ ] **feat:** next-season detection — auto-link sequels (e.g., "Attack on Titan S3 → Final Season")
- [ ] **test:** add tests and CI/CD auto-deploy
- [ ] **feat:** adding dubbed anime
- [ ] **fix:**  fixing subtitle issues for auto HLS stream
- [ ] **fix:**  better error displays when user not logged in
- [ ] **feat:** sorting and filtering search results
- [ ] **feat:** adding logging

## License

This project is provided as-is for personal use.

## Support

For issues or questions, please refer to the AniList or Stremio documentation.
