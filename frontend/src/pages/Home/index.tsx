import Header from "../../components/Header";

export default function Home() {
  return (
    <div className="min-h-screen bg-[var(--bg)] text-[var(--text-main)]">
      <Header />

      <main className="mx-auto max-w-5xl px-6 py-24">
        <p className="text-xs uppercase tracking-[0.3em] text-[var(--text-muted)]">
          Self-hosted • Stremio • AniList
        </p>

        <h1 className="mt-8 text-6xl font-bold leading-none tracking-tight md:text-7xl">
          Stream anime.
          <br />
          Sync AniList.
        </h1>

        <p className="mt-8 max-w-2xl text-lg leading-8 text-[var(--text-muted)]">
          AnilistStream is a self-hostable Stremio addon that streams anime over
          HTTP and automatically syncs your watch progress with AniList. Deploy
          your own Cloudflare Worker and keep your library up to date across
          multiple providers.
        </p>

        <div className="mt-12 flex flex-wrap gap-4">
          <a
            href="/configure"
            className="border border-[var(--primary)] bg-[var(--primary)] px-6 py-3 transition-colors hover:bg-[var(--primary-dark)]"
          >
            Install Addon
          </a>

          <a
            href="https://github.com/Saadiq8149/AnilistStream"
            target="_blank"
            rel="noreferrer"
            className="border border-[var(--border-primary)] px-6 py-3 transition-colors hover:bg-white/5"
          >
            View on GitHub
          </a>
        </div>

        {/*<section className="mt-24 border border-[var(--border-primary)]">
          <div className="border-b border-[var(--border-primary)] px-6 py-3 text-xs uppercase tracking-[0.2em] text-[var(--text-muted)]">
            Features
          </div>

          <div className="grid md:grid-cols-3">
            <div className="border-b border-[var(--border-primary)] p-6 md:border-b-0 md:border-r">
              <h3 className="mb-3 font-medium">HTTP Streaming</h3>

              <p className="text-sm leading-7 text-[var(--text-muted)]">
                Watch anime through HTTP providers directly inside Stremio
                without relying on torrents.
              </p>
            </div>

            <div className="border-b border-[var(--border-primary)] p-6 md:border-b-0 md:border-r">
              <h3 className="mb-3 font-medium">AniList Sync</h3>

              <p className="text-sm leading-7 text-[var(--text-muted)]">
                Automatically update your AniList watch progress while you watch
                episodes.
              </p>
            </div>

            <div className="p-6">
              <h3 className="mb-3 font-medium">Self Hosted</h3>

              <p className="text-sm leading-7 text-[var(--text-muted)]">
                Deploy your own Cloudflare Worker and keep full control over
                your installation.
              </p>
            </div>
          </div>
        </section>*/}

        {/*<section className="mt-20 border border-[var(--border-primary)]">
          <div className="border-b border-[var(--border-primary)] px-6 py-3 text-xs uppercase tracking-[0.2em] text-[var(--text-muted)]">
            How it works
          </div>

          <div className="grid md:grid-cols-3">
            <div className="border-b border-[var(--border-primary)] p-6 md:border-b-0 md:border-r">
              <div className="mb-4 text-3xl font-light text-[var(--primary)]">
                01
              </div>

              <h3 className="mb-2 font-medium">Deploy</h3>

              <p className="text-sm leading-7 text-[var(--text-muted)]">
                Deploy the Cloudflare Worker using your own account.
              </p>
            </div>

            <div className="border-b border-[var(--border-primary)] p-6 md:border-b-0 md:border-r">
              <div className="mb-4 text-3xl font-light text-[var(--primary)]">
                02
              </div>

              <h3 className="mb-2 font-medium">Configure</h3>

              <p className="text-sm leading-7 text-[var(--text-muted)]">
                Connect your worker and optionally log in with AniList.
              </p>
            </div>

            <div className="p-6">
              <div className="mb-4 text-3xl font-light text-[var(--primary)]">
                03
              </div>

              <h3 className="mb-2 font-medium">Watch</h3>

              <p className="text-sm leading-7 text-[var(--text-muted)]">
                Install the addon in Stremio and enjoy automatic progress
                syncing.
              </p>
            </div>
          </div>
        </section>*/}

        <section className="mt-20 flex flex-col items-start justify-between gap-8 border border-[var(--border-primary)] p-8 md:flex-row md:items-center">
          <div>
            <p className="text-xs uppercase tracking-[0.2em] text-[var(--text-muted)]">
              Ready?
            </p>

            <h2 className="mt-3 text-3xl font-bold">Install AnilistStream</h2>

            <p className="mt-3 max-w-xl text-[var(--text-muted)]">
              Configure your Cloudflare Worker, connect AniList, and install the
              addon into Stremio in just a few minutes.
            </p>
          </div>

          <a
            href="/configure"
            className="border border-[var(--primary)] bg-[var(--primary)] px-8 py-3 transition-colors hover:bg-[var(--primary-dark)]"
          >
            Get Started
          </a>
        </section>
      </main>
    </div>
  );
}
