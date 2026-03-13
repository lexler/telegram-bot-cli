package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "auth":
		runAuth()
	case "status":
		runStatus()
	case "send":
		runSend(args)
	case "trace":
		runTrace()
	case "config":
		runConfig(args)
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: telegram-bot-cli <command> [options]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  auth                          Set bot token")
	fmt.Println("  status                        Verify bot token works")
	fmt.Println("  send [options]                Send a message")
	fmt.Println("    --chat <id>                 Chat ID (uses default if not set)")
	fmt.Println("    --message, -m <text>        Message text")
	fmt.Println("  trace                         Discover chat ID by pasting a code")
	fmt.Println("  config set <key> <value>      Set a config value")
	fmt.Println("  config get <key>              Get a config value")
}
