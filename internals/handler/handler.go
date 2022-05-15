package handler

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/koba1108/discord-bot/internals/config"
	"log"
)

type EventType string

const (
	EventTypeJoin  = EventType("Join")
	EventTypeLeave = EventType("Leave")
)

func OnVoiceStateUpdate(s *discordgo.Session, m *discordgo.VoiceStateUpdate) {
	eventType := GetEventType(m)
	if eventType == "" {
		return
	}
	user, err := s.User(m.UserID)
	if err != nil {
		log.Println("Error get user: ", err)
		return
	}
	msg := eventType.Msg(user.Username)
	if _, err := s.ChannelMessageSend(config.DiscordTextChannelID(), msg); err != nil {
		log.Println("Error sending message: ", err)
	}
}

func GetEventType(m *discordgo.VoiceStateUpdate) EventType {
	switch {
	case m.BeforeUpdate == nil:
		return EventTypeJoin
	case m.ChannelID != m.BeforeUpdate.ChannelID:
		return EventTypeLeave
	}
	return ""
}

func (et EventType) Msg(username string) string {
	switch et {
	case EventTypeJoin:
		return fmt.Sprintf("@here %sさんが入室しました。", username)
	case EventTypeLeave:
		return fmt.Sprintf("@here %sさんが退室しました。", username)
	}
	return ""
}
