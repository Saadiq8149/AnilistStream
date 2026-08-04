import { Hono } from "hono";
import { serveStatic } from "hono/bun";
import { cors } from "hono/cors";

const app = new Hono();

app.use("*", cors());
app.use("/*", serveStatic({ root: "./public" }));

export default app;
