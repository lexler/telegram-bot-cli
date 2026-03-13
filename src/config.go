package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Token         string `toml:"token"`
	DefaultChat   string `toml:"default_chat"`
	DefaultThread string `toml:"default_thread,omitempty"`
}

func configDir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "telegram-bot-cli")
}

func configPath() string {
	return filepath.Join(configDir(), "config.toml")
}

func loadConfig() (*Config, error) {
	data, err := os.ReadFile(configPath())
	if err != nil {
		if os.IsNotExist(err) {
			return &Config{}, nil
		}
		return nil, err
	}

	var cfg Config
	err = toml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func saveConfig(cfg *Config) error {
	err := os.MkdirAll(configDir(), 0700)
	if err != nil {
		return err
	}

	data, err := toml.Marshal(cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(configPath(), data, 0600)
}

func runConfig(args []string) {
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "Usage: telegram-bot-cli config <set|get> <key> [value]")
		os.Exit(1)
	}

	subcommand := args[0]

	switch subcommand {
	case "set":
		if len(args) < 3 {
			fmt.Fprintln(os.Stderr, "Usage: telegram-bot-cli config set <key> <value>")
			os.Exit(1)
		}
		configSet(args[1], args[2])
	case "get":
		if len(args) < 2 {
			fmt.Fprintln(os.Stderr, "Usage: telegram-bot-cli config get <key>")
			os.Exit(1)
		}
		configGet(args[1])
	default:
		fmt.Fprintf(os.Stderr, "Unknown config subcommand: %s\n", subcommand)
		os.Exit(1)
	}
}

func configSet(key, value string) {
	cfg, err := loadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		os.Exit(1)
	}

	switch key {
	case "default-chat":
		cfg.DefaultChat = value
	default:
		fmt.Fprintf(os.Stderr, "Unknown config key: %s\n", key)
		fmt.Fprintln(os.Stderr, "Available keys: default-chat")
		os.Exit(1)
	}

	err = saveConfig(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error saving config: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Set %s = %s\n", key, value)
}

func configGet(key string) {
	cfg, err := loadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		os.Exit(1)
	}

	switch key {
	case "default-chat":
		if cfg.DefaultChat == "" {
			fmt.Println("(not set)")
		} else {
			fmt.Println(cfg.DefaultChat)
		}
	case "token":
		if cfg.Token == "" {
			fmt.Println("(not set)")
		} else {
			fmt.Println("(set)")
		}
	default:
		fmt.Fprintf(os.Stderr, "Unknown config key: %s\n", key)
		fmt.Fprintln(os.Stderr, "Available keys: default-chat, token")
		os.Exit(1)
	}
}
