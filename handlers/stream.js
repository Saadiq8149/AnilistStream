import {
    getAnimeStreams,
    updateUserWatchStatusOnAnilist,
} from "../core/addon.js";
import { getAnilistId } from "../core/anilist.js";
import express from "express";

const app = express();
export default app;

app.get("/:anilistToken/stream/:type/:id.json", async (req, res) => {
    try {
        const { anilistToken, id } = req.params;

        let animeId, title, episode, _;

        if (id.startsWith("ani_")) {
            [_, animeId, title, episode] = id.split("_");
        } else if (id.startsWith("kitsu")) {
            [_, animeId, episode] = id.split(":");
        } else {
            return res.json({ streams: [] });
        }

        if (!title) {
            // converting kitsu id to anilist id
            [animeId, title] = await getAnilistId(animeId);
        }
        const streams = await getAnimeStreams(animeId, title, episode);

        // Update user's watch status on Anilist
        if (anilistToken) {
            updateUserWatchStatusOnAnilist(
                anilistToken,
                animeId,
                episode,
                streams,
            );
        }

        res.json({ streams });
    } catch (err) {
        console.log("Stream error:", err);
        res.json({ streams: [] });
    }
});

app.get("/stream/:type/:id.json", async (req, res) => {
    try {
        const { id } = req.params;
        let animeId, title, episode, _;

        if (id.startsWith("ani_")) {
            [_, animeId, title, episode] = id.split("_");
        } else if (id.startsWith("kitsu")) {
            [_, animeId, episode] = id.split(":");
        } else {
            return res.json({ streams: [] });
        }

        if (!title) {
            // converting kitsu id to anilist id
            [animeId, title] = await getAnilistId(animeId);
        }

        const streams = await getAnimeStreams(animeId, title, episode);

        res.json({ streams });
    } catch (err) {
        console.log("Stream error:", err);
        res.json({ streams: [] });
    }
});
