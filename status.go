package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type telegramResponse struct {
	OK     bool            `json:"ok"`
	Result json.RawMessage `json:"result"`
	Desc   string          `json:"description"`
}

type botInfo struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
}

func runStatus() {
	cfg, err := loadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		os.Exit(1)
	}

	if cfg.Token == "" {
		fmt.Fprintln(os.Stderr, "Not authenticated. Run 'telegram-bot-cli auth' first.")
		os.Exit(1)
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/getMe", cfg.Token)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to Telegram: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	var tgResp telegramResponse
	err = json.NewDecoder(resp.Body).Decode(&tgResp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing response: %v\n", err)
		os.Exit(1)
	}

	if !tgResp.OK {
		fmt.Fprintf(os.Stderr, "Telegram API error: %s\n", tgResp.Desc)
		os.Exit(1)
	}

	var bot botInfo
	err = json.Unmarshal(tgResp.Result, &bot)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing bot info: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Authenticated as @%s (%s)\n", bot.Username, bot.FirstName)
}
