# User Guide

## Commands

```
telegram-bot-cli auth              Set bot token
telegram-bot-cli status            Verify bot token works (calls getMe)
telegram-bot-cli send [options]    Send a message
  --chat <id>                      Chat ID (uses default if not set)
  --thread <id>                    Topic thread ID (uses default if not set)
  --message, -m <text>             Message text
telegram-bot-cli trace             Discover chat ID by pasting a code
telegram-bot-cli config set <key> <value>
telegram-bot-cli config get <key>
```

## Setup

### Getting a bot token

1. Open Telegram, search for `@BotFather`
2. Send `/newbot`
3. Give it a name (e.g., "Exercise Tracker")
4. Give it a username (must end in `bot`, e.g., `archie_exercise_bot`)
5. BotFather replies with your token like `123456789:ABCdefGHIjklMNOpqrsTUVwxyz`

### Authenticating

```bash
telegram-bot-cli auth
# Enter bot token: <paste your token>
# Token saved. Run 'telegram-bot-cli status' to verify.

telegram-bot-cli status
# Authenticated as @your_bot (Your Bot)
```

### Finding your chat ID

```bash
telegram-bot-cli trace
# Paste this code into the chat where you want to send messages:
# (copied to clipboard)
#
#   /trace_a1b2c3d4
#
# Waiting for message...
```

Paste the code into your Telegram chat. The bot must be an admin in groups/channels to see messages. Once found, it offers to set it as your default chat.

## Config

Location: `~/.config/telegram-bot-cli/config.toml`

```toml
token = "bot123456:ABC..."
default_chat = "12345678"
default_thread = "42"
```

Available config keys: `default-chat`, `token`

## Sending messages

```bash
telegram-bot-cli send -m "Hello from the CLI"

telegram-bot-cli send --chat 12345678 -m "To a specific chat"

telegram-bot-cli send --thread 42 -m "To a specific topic"
```

## Scope

- Send text messages only (no files/images/video)
- No receive/download functionality
- Config-first approach (set token once, reuse)
