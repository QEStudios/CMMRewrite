package commands

import (
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/QEStudios/CMMRewrite/util"
	"github.com/bwmarrin/discordgo"
)

var UwuifyCommand = discordgo.ApplicationCommand{
	Name:        "uwuify",
	Description: "Convert a message into uwu-speak",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "message",
			Description: "The text to uwuify",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    true,
		},
	},
}

func uwuify(s string) string {
	faceChance := 0.25
	face := []string{" (・`ω´・) ", " ;;w;; ", " owo ", " UwU ", " >w< ", " ^w^ "}

	uwu := s

	if len(s) > 0 {
		fr := "rlRL"
		to := "wwWW"
		for i := range fr {
			uwu = strings.ReplaceAll(uwu, string(fr[i]), string(to[i]))
		}
		re := regexp.MustCompile(`(?i)(n)([aeiou])`)
		uwu = re.ReplaceAllStringFunc(uwu, func(match string) string {
			if len(match) >= 2 {
				return string(match[0]) + "y" + string(match[1])
			}
			return match
		})
		uwu = strings.ReplaceAll(uwu, "ove", "uv")

		rand.Seed(time.Now().UnixNano())

		re = regexp.MustCompile(`(\.+)`)
		uwu = re.ReplaceAllStringFunc(uwu, func(match string) string {
			if rand.Float64() < faceChance {
				return face[rand.Intn(len(face))]
			}
			return match
		})
	}

	return uwu
}

func UwuifyHandler(s *discordgo.Session, i *discordgo.InteractionCreate, opts util.OptionMap) {
	uwu := uwuify(opts["message"].StringValue())
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: uwu,
		},
	})
}
