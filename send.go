package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func runSend(args []string) {
	var chatID string
	var threadID string
	var message string

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--chat":
			if i+1 < len(args) {
				chatID = args[i+1]
				i++
			}
		case "--thread":
			if i+1 < len(args) {
				threadID = args[i+1]
				i++
			}
		case "--message", "-m":
			if i+1 < len(args) {
				message = args[i+1]
				i++
			}
		}
	}

	cfg, err := loadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		os.Exit(1)
	}

	if cfg.Token == "" {
		fmt.Fprintln(os.Stderr, "Not authenticated. Run 'telegram-bot-cli auth' first.")
		os.Exit(1)
	}

	if chatID == "" {
		chatID = cfg.DefaultChat
	}
	if chatID == "" {
		fmt.Fprintln(os.Stderr, "No chat ID specified. Use --chat or set default-chat in config.")
		os.Exit(1)
	}

	if threadID == "" {
		threadID = cfg.DefaultThread
	}

	if message == "" {
		fmt.Fprintln(os.Stderr, "No message specified. Use --message or -m.")
		os.Exit(1)
	}

	params := url.Values{
		"chat_id": {chatID},
		"text":    {message},
	}
	if threadID != "" {
		params.Set("message_thread_id", threadID)
	}

	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", cfg.Token)
	resp, err := http.PostForm(apiURL, params)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error sending message: %v\n", err)
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

	fmt.Println("Message sent.")
}
