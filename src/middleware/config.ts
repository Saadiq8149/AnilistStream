export interface StremioUserConfig {
  anilistAuthToken?: string;
  cfWorkerURL: string;
}

import { createMiddleware } from "hono/factory";

export const configMiddleware = createMiddleware<{
  Variables: {
    config: StremioUserConfig;
  };
}>(async (c, next) => {
  const encoded = c.req.param("config");

  if (!encoded) {
    return c.json({ error: "Missing config" }, 400);
  }

  try {
    const json = atob(encoded.replace(/-/g, "+").replace(/_/g, "/"));
    const config = JSON.parse(json) as StremioUserConfig;

    c.set("config", config);

    await next();
  } catch {
    return c.json({ error: "Invalid config" }, 400);
  }
});
