# Project Notes

## How to release

```bash
git tag v0.0.X
git push origin v0.0.X
```

The release workflow builds Mac Apple Silicon and Mac Intel binaries and attaches them to a GitHub release automatically.

## Project structure

```
telegram-bot-cli/
├── src/                   # Go source code (module root)
│   ├── main.go            # entry point + command routing
│   ├── config.go          # ~/.config/telegram-bot-cli/config.toml
│   ├── auth.go            # set bot token
│   ├── status.go          # verify token (getMe API)
│   ├── send.go            # send message to chat
│   ├── trace.go           # discover chat ID interactively
│   └── main_test.go       # placeholder test
├── build.sh               # builds to build/telegram-bot-cli
├── test.sh                # runs go test
├── docs/
│   ├── guide.md           # user guide
│   └── project_notes.md   # this file
└── .github/workflows/
    ├── build.yml           # CI: test + build on push/PR
    └── release.yml         # release: build binaries on v* tag
```

## Origin

Extracted from `lexler/exercise-bot` repo (`telegram_cli/` directory). The extraction process is documented in `exercise-bot/docs/extract_subproject.process.md`.

## Key decisions

- Binary name: `telegram-bot-cli` (config dir: `~/.config/telegram-bot-cli/`)
- Go source lives in `src/` (not repo root)
- macOS only (Apple Silicon + Intel)
- Distribution via GitHub releases (binary download), no package managers
- Single dependency: `github.com/pelletier/go-toml/v2`
