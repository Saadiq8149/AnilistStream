import express from "express";
import cors from "cors";
import middleware from "./handlers/middleware.js";
import staticRouter from "./handlers/static.js";
import meta from "./handlers/meta.js";
import poster from "./handlers/poster.js";
import stream from "./handlers/stream.js";
import subtitles from "./handlers/subtitles.js";
import catalog from "./handlers/catalog.js";
import "dotenv/config";

const app = express();

app.use(cors());
app.use(express.static("public"));
app.set("trust proxy", 1);

app.use(middleware);
app.use(staticRouter);
app.use(meta);
app.use(poster);
app.use(stream);
app.use(subtitles);
app.use(catalog);

const PORT = process.env.PORT || 7000;
const HOST = "127.0.0.1";
app.listen(PORT, HOST, () => {
    console.log(`AnilistStream running at http://${HOST}:${PORT}`);
    console.log(`Visit http://${HOST}:${PORT}/configure to set up your token.`);
});
