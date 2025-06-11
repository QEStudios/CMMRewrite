package commands

import (
	"github.com/QEStudios/CMMRewrite/util"
	"github.com/bwmarrin/discordgo"
)

var PingCommand = discordgo.ApplicationCommand{
	Name:        "ping",
	Description: "Responds with pong",
}

func PingHandler(s *discordgo.Session, i *discordgo.InteractionCreate, opts util.OptionMap) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong!",
		},
	})
}
