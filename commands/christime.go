package commands

import (
	"fmt"
	"time"

	"github.com/QEStudios/CMMRewrite/util"
	"github.com/bwmarrin/discordgo"
)

var ChrisTimeCommand = discordgo.ApplicationCommand{
	Name:        "christime",
	Description: "See what time it is for chris",
}

func ChrisTimeHandler(s *discordgo.Session, i *discordgo.InteractionCreate, opts util.OptionMap) {
	loc, err := time.LoadLocation("US/Pacific")
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
