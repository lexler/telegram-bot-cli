package main

import (
	"bufio"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

type update struct {
	UpdateID int `json:"update_id"`
	Message  struct {
		Text            string `json:"text"`
		MessageThreadID int64  `json:"message_thread_id"`
		IsTopicMessage  bool   `json:"is_topic_message"`
		Chat            struct {
			ID    int64  `json:"id"`
			Title string `json:"title"`
			Type  string `json:"type"`
		} `json:"chat"`
	} `json:"message"`
}

type updatesResponse struct {
	OK     bool     `json:"ok"`
	Result []update `json:"result"`
	Desc   string   `json:"description"`
}

func runTrace() {
	cfg, err := loadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		os.Exit(1)
	}

	if cfg.Token == "" {
		fmt.Fprintln(os.Stderr, "Not authenticated. Run 'telegram-bot-cli auth' first.")
		os.Exit(1)
	}

	traceCode := generateTraceCode()
	copyToClipboard(traceCode)

	fmt.Println("Paste this code into the chat where you want to send messages:")
	fmt.Println("(copied to clipboard)")
	fmt.Println()
	fmt.Printf("  %s\n", traceCode)
	fmt.Println()
	fmt.Println("Note: The bot must be an admin in the group/channel to see messages.")
	fmt.Println()
	fmt.Println("Waiting for message...")

	chatID, chatName, chatType, threadID := waitForTrace(cfg.Token, traceCode)

	fmt.Println()
	if chatName != "" {
		fmt.Printf("Found: %s (%s)\n", chatName, chatType)
	} else {
		fmt.Printf("Found: %s chat\n", chatType)
	}
	fmt.Printf("Chat ID: %d\n", chatID)
	if threadID != 0 {
		fmt.Printf("Topic thread ID: %d\n", threadID)
	}
	fmt.Println()

	if promptSetDefault(chatID) {
		cfg.DefaultChat = fmt.Sprintf("%d", chatID)
		if threadID != 0 {
			cfg.DefaultThread = fmt.Sprintf("%d", threadID)
		}
		err := saveConfig(cfg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error saving config: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Default chat updated.")
	}
}

func generateTraceCode() string {
	bytes := make([]byte, 4)
	rand.Read(bytes)
	return "/trace_" + hex.EncodeToString(bytes)
}

func copyToClipboard(text string) {
	cmd := exec.Command("pbcopy")
	cmd.Stdin = strings.NewReader(text)
	cmd.Run()
}

func promptSetDefault(chatID int64) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Set %d as default chat? [Y/n] ", chatID)
	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(strings.ToLower(answer))
	return answer == "" || answer == "y" || answer == "yes"
}

func waitForTrace(token, traceCode string) (int64, string, string, int64) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates", token)

	for {
		resp, err := http.Get(url)
		if err != nil {
			time.Sleep(time.Second)
			continue
		}

		var updates updatesResponse
		json.NewDecoder(resp.Body).Decode(&updates)
		resp.Body.Close()

		if !updates.OK {
			time.Sleep(time.Second)
			continue
		}

		for _, u := range updates.Result {
			if u.Message.Text == traceCode {
				return u.Message.Chat.ID, u.Message.Chat.Title, u.Message.Chat.Type, u.Message.MessageThreadID
			}
		}

		time.Sleep(time.Second)
	}
}
