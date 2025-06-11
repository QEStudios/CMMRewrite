package util

import "github.com/bwmarrin/discordgo"

type OptionMap = map[string]*discordgo.ApplicationCommandInteractionDataOption

// https://github.com/bwmarrin/discordgo/blob/6e8fa27c7917ea54d8b9ec26f126becae59058d2/examples/echo/main.go#L15
func ParseOptions(options []*discordgo.ApplicationCommandInteractionDataOption) (om OptionMap) {
	om = make(OptionMap)
	for _, opt := range options {
		om[opt.Name] = opt
	}
	return
}
