const { addonBuilder } = require("stremio-addon-sdk");
const { searchAnime, getAnimeDetails } = require("./anilist");
const { getAnimeByAnilistId, getEpisodeUrls } = require("./anicli");

const manifest = {
  id: "community.AnilistStream",
  version: "0.0.1",
  catalogs: [
    {
      type: "series",
      id: "anilist",
      name: "Anime",
      extra: [
        {
          name: "search",
          isRequired: true,
        },
      ],
    },
  ],
  resources: [
    "catalog",
    {
      name: "meta",
      types: ["series"],
      idPrefixes: ["ani_"],
    },
    "stream",
  ],
  types: ["series", "movie"],
  name: "AnilistStream",
  description: "Streaming anime and Anilist sync",
  idPrefixes: ["ani_"],
  behaviorHints: {
    configurable: true,
    configurationRequired: true,
  },
  config: [
    {
      key: "anilist_access_token",
      type: "text",
      title: "Anilist Access Token",
      required: true,
    },
  ],
};
const builder = new addonBuilder(manifest);

builder.defineCatalogHandler(async ({ type, id, extra }) => {
  var anime = [];
  if (extra != null) {
    const searchQuery = extra.search;
    anime = await searchAnime(searchQuery);
  }

  return { metas: anime };
});

builder.defineMetaHandler(async ({ type, id }) => {
  return { meta: await getAnimeDetails(id) };
});

builder.defineStreamHandler(async ({ type, id }) => {
  if (!id.startsWith("ani_")) {
    return { streams: [] };
  }

  console.log("request for streams: " + type + " " + id);
  const animeId = id.split("_")[1];
  const title = id.split("_")[2].replace("?", "").replace("!", "");
  const episodeNumber = id.split("_")[3] || 1;
  const privateId = await getAnimeByAnilistId(animeId, title);
  const sources = await getEpisodeUrls(privateId.id, episodeNumber);

  const streams = sources.map((source) => ({
    url: source.url,
    name: `AnilistStream`,
    description: `Source: ${source.source} - Quality: ${source.quality}`,
    subtitles: source.subtitles
      ? [{ id: "eng", lang: "English", url: source.subtitles }]
      : [],
    behaviorHints: {
      notWebReady: true,
      proxyHeaders: {
        request: {
          Referer: source.referrer,
          "User-Agent": source["user-agent"],
        },
      },
    },
  }));

  if (streams.length > 0) {
    return { streams: streams };
  }

  return { streams: [] };
});

module.exports = builder.getInterface();
