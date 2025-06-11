package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/QEStudios/CMMRewrite/commands"
	"github.com/QEStudios/CMMRewrite/util"
	"github.com/bwmarrin/discordgo"
	"github.com/lpernett/godotenv"
)

var discord *discordgo.Session
var App string
var Guild string

func init() {
	var err error
	envFile, _ := godotenv.Read(".env")
	var Token = envFile["BOT_TOKEN"]
	App = envFile["APP_ID"]
	Guild = envFile["GUILD_ID"]

	discord, err = discordgo.New("Bot " + Token)
	if err != nil {
		log.Fatalf("Could not login: %v", err)
	}
}

var (
	commandDefs = []*discordgo.ApplicationCommand{
		&commands.PingCommand,
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate, opts util.OptionMap){
		"ping": commands.PingHandler,
	}
)

func main() {
	discord.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as %s", r.User.String())
		s.UpdateGameStatus(0, "Circuit Maker 2")
	})

	discord.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionApplicationCommand {
			return
		}

		data := i.ApplicationCommandData()
		commandHandlers[data.Name](s, i, util.ParseOptions(data.Options))
	})

	_, err := discord.ApplicationCommandBulkOverwrite(App, Guild, commandDefs)
	if err != nil {
		log.Fatalf("Could not register commands: %v", err)
	}

	err = discord.Open()
	if err != nil {
		log.Fatalf("Could not open session: %v", err)
	}

	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, os.Interrupt)
	<-sigch

	err = discord.Close()
	if err != nil {
		log.Printf("Could not gracefully close session: %v", err)
	}
}
