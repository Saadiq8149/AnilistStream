const { getAnimeByAnilistId, getEpisodeUrls } = require("./anicli");
const { getUserWatchStatus, updateUserWatchList } = require("./anilist");

async function getAnimeStreams(animeId, title, episodeNumber) {
  const privateId = await getAnimeByAnilistId(animeId, title);
  if (!privateId) return [];

  const sources = await getEpisodeUrls(privateId.id, episodeNumber);
  var streams = [];

  for (const source of sources.sub) {
    streams.push({
      url: source.url,
      name: `AnilistStream Sub ${source.type}`,
      description: `Source: ${source.quality}`,
      subtitles: source.subtitles
        ? [{ id: "eng", lang: "English", url: source.subtitles }]
        : [],
      behaviorHints: {
        notWebReady: true,
        proxyHeaders: {
          request: {
            Referer: source.referer,
            "User-Agent": source["user-agent"],
          },
        },
      },
    });
  }

  for (const source of sources.dub) {
    streams.push({
      url: source.url,
      name: "AnilistStream Dub",
      description: `Quality: ${source.quality}`,
      subtitles: source.subtitles
        ? [{ id: "eng", lang: "English", url: source.subtitles }]
        : [],
      behaviorHints: {
        notWebReady: true,
        proxyHeaders: {
          request: {
            Referer: source.referer,
            "User-Agent": source["user-agent"],
          },
        },
      },
    });
  }

  return streams;
}

async function updateUserWatchStatusOnAnilist(
  anilistToken,
  animeId,
  episodeNumber,
  streams
) {
  if (anilistToken && streams.length > 0) {
    const userWatchStatus = await getUserWatchStatus(anilistToken, animeId);
    if (userWatchStatus) {
      switch (userWatchStatus) {
        case "PLANNING":
          await updateUserWatchList(
            anilistToken,
            animeId,
            episodeNumber,
            "CURRENT"
          );
          break;
        case "COMPLETED":
          await updateUserWatchList(
            anilistToken,
            animeId,
            episodeNumber,
            "REPEATING"
          );
          break;
        case "REPEATING":
          await updateUserWatchList(
            anilistToken,
            animeId,
            episodeNumber,
            "REPEATING"
          );
          break;
        default:
          await updateUserWatchList(
            anilistToken,
            animeId,
            episodeNumber,
            "CURRENT"
          );
      }
    }
  }
}

module.exports = {
  getAnimeStreams,
  updateUserWatchStatusOnAnilist,
};
