import { useEffect, useState } from "preact/hooks";
import Header from "../../components/Header";

const CLIENT_ID = import.meta.env.VITE_ANILIST_CLIENT_ID;

export default function Configure() {
  const [cfWorkerURL, setCfWorkerURL] = useState("");
  const [anilistToken, setAnilistToken] = useState<string | null>(null);

  useEffect(() => {
    const params = new URLSearchParams(window.location.hash.slice(1));
    const token = params.get("access_token");

    if (token) {
      sessionStorage.setItem("anilist_token", token);
      setAnilistToken(token);

      history.replaceState(null, "", window.location.pathname);
      return;
    }

    const saved = sessionStorage.getItem("anilist_token");
    if (saved) setAnilistToken(saved);
  }, []);

  function loginToAniList() {
    const url =
      `https://anilist.co/api/v2/oauth/authorize` +
      `?client_id=${CLIENT_ID}` +
      `&response_type=token`;

    window.location.href = url;
  }

  function disconnectAniList() {
    sessionStorage.removeItem("anilist_token");
    setAnilistToken(null);
  }

  function encodeConfig(config: object): string {
    return btoa(JSON.stringify(config))
      .replace(/\+/g, "-")
      .replace(/\//g, "_")
      .replace(/=+$/, "");
  }

  async function submit(e: Event) {
    e.preventDefault();

    const config = {
      cfWorkerURL,
      anilistAuthToken: anilistToken ?? undefined,
    };

    const encoded = encodeConfig(config);
    const manifest = `${window.location.origin}/${encoded}/manifest.json`;

    const isLocalhost =
      window.location.hostname === "localhost" ||
      window.location.hostname === "127.0.0.1";

    if (isLocalhost) {
      try {
        await navigator.clipboard.writeText(manifest);
        alert("Manifest URL copied to clipboard.");
      } catch {
        alert(manifest);
      }
      return;
    }

    window.location.href = `stremio://${window.location.host}/${encoded}/manifest.json`;
  }

  return (
    <div className="min-h-screen bg-[var(--bg)] text-[var(--text-main)]">
      <Header />

      <main className="mx-auto max-w-4xl px-6 py-20">
        <div className="border border-[var(--border-primary)]">
          {/* Header */}

          <div className="border-b border-[var(--border-primary)] px-8 py-8">
            <p className="text-xs uppercase tracking-[0.25em] text-[var(--text-muted)]">
              Configuration
            </p>

            <h1 className="mt-4 text-4xl font-bold">Configure AnilistStream</h1>

            <p className="mt-4 max-w-2xl leading-7 text-[var(--text-muted)]">
              Connect your Cloudflare Worker and optionally log in with AniList
              to enable automatic watch progress syncing.
            </p>
          </div>

          <form onSubmit={submit}>
            {/* Worker */}

            <section className="border-b border-[var(--border-primary)] p-8">
              <label className="block text-xs uppercase tracking-[0.2em] text-[var(--text-muted)]">
                Cloudflare Worker URL
              </label>

              <input
                required
                type="url"
                value={cfWorkerURL}
                onInput={(e) =>
                  setCfWorkerURL((e.currentTarget as HTMLInputElement).value)
                }
                placeholder="https://example.workers.dev"
                className="mt-4 w-full border border-[var(--border-primary)] bg-transparent px-4 py-3 outline-none transition-colors focus:border-[var(--primary)]"
              />

              <p className="mt-3 text-sm leading-6 text-[var(--text-muted)]">
                The URL of the Cloudflare Worker you deployed.
              </p>
            </section>

            {/* AniList */}

            <section className="border-b border-[var(--border-primary)] p-8">
              <div className="flex items-center justify-between">
                <div>
                  <p className="text-xs uppercase tracking-[0.2em] text-[var(--text-muted)]">
                    AniList
                  </p>

                  <h2 className="mt-2 text-xl font-semibold">
                    Account Connection
                  </h2>

                  <p className="mt-2 text-sm leading-6 text-[var(--text-muted)]">
                    Optional. Required only if you want automatic watch progress
                    syncing.
                  </p>
                </div>

                {anilistToken ? (
                  <span className="border border-green-600 px-3 py-1 text-xs uppercase tracking-wider text-green-400">
                    Connected
                  </span>
                ) : (
                  <span className="border border-[var(--border-primary)] px-3 py-1 text-xs uppercase tracking-wider text-[var(--text-muted)]">
                    Not Connected
                  </span>
                )}
              </div>

              {anilistToken ? (
                <div className="mt-8 border border-[var(--border-primary)]">
                  <div className="border-b border-[var(--border-primary)] px-6 py-3 text-xs uppercase tracking-widest text-[var(--text-muted)]">
                    Access Token
                  </div>

                  <div className="flex flex-col gap-6 p-6 md:flex-row md:items-center md:justify-between">
                    <code className="break-all font-mono text-sm text-[var(--text-muted)]">
                      {anilistToken.slice(0, 12)}
                      ••••••••••••••••
                      {anilistToken.slice(-8)}
                    </code>

                    <button
                      type="button"
                      onClick={disconnectAniList}
                      className="border border-red-700 px-5 py-3 transition-colors hover:bg-red-900/20"
                    >
                      Disconnect
                    </button>
                  </div>
                </div>
              ) : (
                <button
                  type="button"
                  onClick={loginToAniList}
                  className="mt-8 border border-[var(--primary)] bg-[var(--primary)] px-6 py-3 transition-colors hover:bg-[var(--primary-dark)]"
                >
                  Login with AniList
                </button>
              )}
            </section>

            {/* Install */}

            <section className="flex flex-col gap-6 p-8 md:flex-row md:items-center md:justify-between">
              <div>
                <p className="text-xs uppercase tracking-[0.2em] text-[var(--text-muted)]">
                  Finish
                </p>

                <h2 className="mt-2 text-2xl font-semibold">Generate Addon</h2>

                <p className="mt-2 max-w-xl text-sm leading-6 text-[var(--text-muted)]">
                  Generate your personalized Stremio addon manifest and install
                  it into Stremio.
                </p>
              </div>

              <button
                type="submit"
                className="border border-[var(--primary)] bg-[var(--primary)] px-8 py-3 transition-colors hover:bg-[var(--primary-dark)]"
              >
                Install Addon
              </button>
            </section>
          </form>
        </div>
      </main>
    </div>
  );
}
