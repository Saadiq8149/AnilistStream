export default function Header() {
  return (
    <header className="border-[var(--border-primary)]">
      <div className="mx-auto flex h-16 max-w-5xl items-center justify-between px-6">
        <a href="/" className="flex items-center gap-3">
          <img src="/logo.png" className="h-7 w-7" alt="" />
          <span className="text-sm font-semibold tracking-wide">
            AnilistStream
          </span>
        </a>

        <nav className="flex items-center">
          <a
            href="https://github.com/Saadiq8149/AnilistStream"
            target="_blank"
            rel="noreferrer"
            className="border-[var(--border-primary)] px-5 py-5 text-sm hover:bg-white/5"
          >
            GitHub
          </a>

          <a
            href="/configure"
            className="border-[var(--border-primary)] bg-[var(--primary)] px-5 py-5 text-sm hover:bg-[var(--primary-dark)]"
          >
            Install
          </a>
        </nav>
      </div>
    </header>
  );
}
