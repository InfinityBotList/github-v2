package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type CreateEvent struct {
	Repo         Repository `json:"repository"`
	Sender       User       `json:"sender"`
	Ref          string     `json:"ref"`
	RefType      string     `json:"ref_type"`
	MasterBranch string     `json:"master_branch"`
	PusherType   string     `json:"pusher_type"`
}

func createFn(bytes []byte) (discordgo.MessageSend, error) {
	var gh CreateEvent

	// Unmarshal the JSON into our struct
	err := json.Unmarshal(bytes, &gh)

	if err != nil {
		return discordgo.MessageSend{}, err
	}

	return discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{
			{
				Color: 0x00ff1a,
				URL:   gh.Repo.HTMLURL,
				Author: &discordgo.MessageEmbedAuthor{
					Name:    gh.Sender.Login,
					IconURL: gh.Sender.AvatarURL,
				},
				Title: "New " + gh.RefType + " created on " + gh.Repo.FullName,
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:  "User",
						Value: fmt.Sprintf("[%s](%s)", gh.Sender.Login, gh.Sender.HTMLURL),
					},
					{
						Name:   "Ref",
						Value:  gh.Ref,
						Inline: true,
					},
					{
						Name:   "Ref Type",
						Value:  gh.RefType,
						Inline: true,
					},
					{
						Name:   "Master Branch",
						Value:  gh.MasterBranch,
						Inline: true,
					},
					{
						Name:   "Pusher Type",
						Value:  gh.PusherType,
						Inline: true,
					},
				},
			},
		},
	}, nil
}
