package util

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

type OptionMap = map[string]*discordgo.ApplicationCommandInteractionDataOption

// https://github.com/bwmarrin/discordgo/blob/6e8fa27c7917ea54d8b9ec26f126becae59058d2/examples/echo/main.go#L15
func ParseOptions(options []*discordgo.ApplicationCommandInteractionDataOption) (om OptionMap) {
	om = make(OptionMap)
	for _, opt := range options {
		om[opt.Name] = opt
	}
	return
}

func HandleErrorAndRespond(s *discordgo.Session, i *discordgo.InteractionCreate, msg string, err ...error) bool {
	if len(err) > 0 {
		if err[0] == nil {
			return false
		}
		fullMsg := fmt.Sprintf("%s: %v", msg, err[0])
		s.FollowupMessageCreate(i.Interaction, false, &discordgo.WebhookParams{
			Content: fullMsg,
		})
		log.Println(fullMsg)
		return true
	}

	// No error provided — just send the message
	s.FollowupMessageCreate(i.Interaction, false, &discordgo.WebhookParams{
		Content: msg,
	})
	log.Println(msg)
	return true
}
