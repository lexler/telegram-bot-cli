package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func runAuth() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter bot token: ")
	token, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	token = strings.TrimSpace(token)
	if token == "" {
		fmt.Fprintln(os.Stderr, "Token cannot be empty")
		os.Exit(1)
	}

	cfg, err := loadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		os.Exit(1)
	}

	cfg.Token = token

	err = saveConfig(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error saving config: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Token saved. Run 'telegram-bot-cli status' to verify.")
}
