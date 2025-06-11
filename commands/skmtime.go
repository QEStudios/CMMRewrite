package commands

import (
	"fmt"
	"time"

	"github.com/QEStudios/CMMRewrite/util"
	"github.com/bwmarrin/discordgo"
)

var SkmTimeCommand = discordgo.ApplicationCommand{
	Name:        "skmtime",
	Description: "See what time it is for skm",
}

func SkmTimeHandler(s *discordgo.Session, i *discordgo.InteractionCreate, opts util.OptionMap) {
	loc, err := time.LoadLocation("Australia/Melbourne")
	if util.HandleErrorAndRespond(s, i, "Error loading timezone", err) {
		return
	}

	melbourneTime := time.Now().In(loc)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Current time for skm: %s", melbourneTime.Format("03:04 PM")),
		},
	})
}
