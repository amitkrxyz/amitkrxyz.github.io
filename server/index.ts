import { Context, Hono } from "hono";
import { env } from "hono/adapter";

type Bindings = {
  SECRET_KEY: string
  ASSETS: any
}

const app = new Hono<{ Bindings: Bindings }>()

app.get("/", async (c) => {
  const userAgent = c.req.header("User-Agent");
  if (userAgent?.startsWith("curl")) {
	  return  c.text(await fetchStaticAsset(c, c.req.url + "/index.txt")) 
  }
  return c.html(await fetchStaticAsset(c, c.req.url));
});

async function fetchStaticAsset(c:Context<{Bindings: Bindings}>, url: string) {
  const res = await c.env.ASSETS.fetch(url)
  const text = await res.text()
  return text
}

export default app;
