package main

import (
	"github.com/koba1108/discord-bot/internals/config"
	"github.com/koba1108/discord-bot/internals/handler"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	log.Printf("Listening... PORT:%s", os.Getenv("PORT"))
	discord, err := discordgo.New(config.DiscordBotToken())
	if err != nil {
		panic(err)
	}
	discord.AddHandler(handler.OnVoiceStateUpdate)
	if err = discord.Open(); err != nil {
		panic(err)
	}
	defer discord.Close()

	stopBot := make(chan os.Signal, 1)
	signal.Notify(stopBot, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-stopBot
	return
}
