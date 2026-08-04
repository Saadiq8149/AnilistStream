import { Hono } from "hono";
import { serveStatic } from "hono/bun";
import { cors } from "hono/cors";

const app = new Hono();

app.use("*", cors());
app.use("/logo.png", serveStatic({ path: "./public/logo.png" }));
app.use(
  "/assets/*",
  serveStatic({
    root: "./frontend/dist",
  }),
);
app.use(
  "/:config/manifest.json",
  serveStatic({ path: "./public/manifest.json" }),
);

app.get("/", serveStatic({ path: "./frontend/dist/index.html" }));
app.get("/configure", serveStatic({ path: "./frontend/dist/index.html" }));
app.get(
  "/:config/configure",
  serveStatic({ path: "./frontend/dist/index.html" }),
);

export default app;
