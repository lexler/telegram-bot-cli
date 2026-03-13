# telegram-bot-cli

A Go CLI for sending messages to Telegram.

## Installation

Download the latest binary from [GitHub Releases](https://github.com/lexler/telegram-bot-cli/releases).

```bash
# macOS Apple Silicon (arm64)
curl -L https://github.com/lexler/telegram-bot-cli/releases/latest/download/telegram-bot-cli-darwin-arm64 -o telegram-bot-cli

# macOS Intel (amd64)
curl -L https://github.com/lexler/telegram-bot-cli/releases/latest/download/telegram-bot-cli-darwin-amd64 -o telegram-bot-cli

chmod +x telegram-bot-cli
mv telegram-bot-cli /usr/local/bin/
```

## Quick start

```bash
telegram-bot-cli auth           # Set your bot token
telegram-bot-cli status         # Verify it works
telegram-bot-cli trace          # Discover your chat ID
telegram-bot-cli send -m "hi"  # Send a message
```

## Getting a bot token

1. Open Telegram, search for `@BotFather`
2. Send `/newbot`
3. Give it a name and username (must end in `bot`)
4. BotFather replies with your token

## Usage

See the [User Guide](docs/guide.md) for full documentation.
