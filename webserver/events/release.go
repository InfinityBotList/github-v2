package events

import (
	"github.com/bwmarrin/discordgo"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type ReleaseEvent struct {
	Action  string     `json:"action"`
	Repo    Repository `json:"repository"`
	Sender  User       `json:"sender"`
	Release struct {
		HTMLUrl string `json:"html_url"`
		Body    string `json:"body"`
		TagName string `json:"tag_name"`
	} `json:"release"`
}

func releaseFn(bytes []byte) (discordgo.MessageSend, error) {
	var gh ReleaseEvent

	// Unmarshal the JSON into our struct
	err := json.Unmarshal(bytes, &gh)

	if err != nil {
		return discordgo.MessageSend{}, err
	}

	var color int
	var title string = cases.Title(language.English).String(gh.Action) + " new release on " + gh.Repo.FullName
	if gh.Action == "created" || gh.Action == "published" || gh.Action == "edited" || gh.Action == "prereleased" || gh.Action == "released" {
		color = 0x00ff1a
	} else {
		color = 0xff0000
	}

	var body string = gh.Release.Body
	if len(gh.Release.Body) > 996 {
		body = gh.Release.Body[:996] + "..."
	}

	if body == "" {
		body = "No description available"
	}

	return discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{
			{
				Color: color,
				URL:   gh.Repo.URL,
				Title: title,
				Author: &discordgo.MessageEmbedAuthor{
					Name:    gh.Sender.Login,
					IconURL: gh.Sender.AvatarURL,
				},
				Description: body,
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:   "User",
						Value:  "[" + gh.Sender.Login + "]" + "(" + gh.Sender.HTMLURL + ")",
						Inline: true,
					},
					{
						Name:   "Release",
						Value:  "[" + gh.Release.TagName + "]" + "(" + gh.Release.HTMLUrl + ")",
						Inline: true,
					},
				},
			},
		},
	}, nil
}
