package config

import (
	"fmt"
	"os"
)

func DiscordClientID() string {
	if os.Getenv("DISCORD_CLIENT_ID") == "" {
		return ""
	}
	return os.Getenv("DISCORD_CLIENT_ID")
}

func DiscordBotToken() string {
	if os.Getenv("DISCORD_BOT_TOKEN") == "" {
		return ""
	}
	return fmt.Sprintf("Bot %s", os.Getenv("DISCORD_BOT_TOKEN"))
}

func DiscordTextChannelID() string {
	return os.Getenv("DISCORD_TEXT_CHANNEL_ID")
}
