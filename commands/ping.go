package commands

import "github.com/bwmarrin/discordgo"

var PingDef = discordgo.ApplicationCommand{
	Name:        "ping",
	Description: "Responds with pong",
}

func HandlePing(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong!",
		},
	})
}
