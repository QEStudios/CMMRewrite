package commands

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/QEStudios/CMMRewrite/util"
	"github.com/bwmarrin/discordgo"
)

var GetUserCommand = discordgo.ApplicationCommand{
	Name:        "getuser",
	Description: "Get the UserId for a roblox username.",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "username",
			Description: "The username to lookup",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    true,
		},
	},
}

type getUserRequest struct {
	Usernames          []string `json:"usernames"`
	ExcludeBannedUsers bool     `json:"excludeBannedUsers"`
}

type getUserResponse struct {
	Data []userData `json:"data"`
}

type userData struct {
	RequestedUsername string `json:"requestedUsername"`
	HasVerifiedBadge  bool   `json:"hasVerifiedBadge"`
	ID                int    `json:"id"`
	Name              string `json:"name"`
	DisplayName       string `json:"displayName"`
}

func GetUserHandler(s *discordgo.Session, i *discordgo.InteractionCreate, opts util.OptionMap) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})
	if err != nil {
		log.Printf("Failed to defer interaction: %v", err)
		return
	}

	reqBody := getUserRequest{
		Usernames:          []string{opts["username"].StringValue()},
		ExcludeBannedUsers: true,
	}

	jsonReq, err := json.Marshal(reqBody)
	if util.HandleErrorAndRespond(s, i, "Error marshaling JSON", err) {
		return
	}

	resp, err := http.Post("https://users.roblox.com/v1/usernames/users", "application/json", bytes.NewBuffer(jsonReq))
	if util.HandleErrorAndRespond(s, i, "Error sending request", err) {
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if util.HandleErrorAndRespond(s, i, "Error reading response body", err) {
		return
	}

	var resBody getUserResponse
	err = json.Unmarshal(body, &resBody)
	if util.HandleErrorAndRespond(s, i, "Error unmarshaling JSON", err) {
		return
	}

	if len(resBody.Data) == 0 {
		util.HandleErrorAndRespond(s, i, "Roblox returned invalid JSON")
		return
	}

	_, err = s.FollowupMessageCreate(i.Interaction, false, &discordgo.WebhookParams{
		Content: strconv.Itoa(resBody.Data[0].ID),
	})
	if err != nil {
		log.Printf("Failed to send follow-up message: %v", err)
		return
	}
}
