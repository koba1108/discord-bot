package handler

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/koba1108/discord-bot/internals/config"
	"log"
)

func OnVoiceStateUpdate(s *discordgo.Session, m *discordgo.VoiceStateUpdate) {
	user, err := s.User(m.UserID)
	if err != nil {
		log.Println("Error get user: ", err)
		return
	}
	msg := fmt.Sprintf("@here %sさんが", user.Username)
	if m.BeforeUpdate == nil {
		msg += "入室しました。"
	} else {
		msg += "退室しました。"
	}
	if _, err := s.ChannelMessageSend(config.DiscordTextChannelID(), msg); err != nil {
		log.Println("Error sending message: ", err)
	}
}
