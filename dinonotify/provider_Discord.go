package dinonotify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type discordMessage struct {
	webhookURL string
}

type discordEmbed struct {
	Description string             `json:"description,omitempty"`
	Color       int                `json:"color,omitempty"`
	Image       *discordEmbedImage `json:"image,omitempty"`
}

type discordEmbedImage struct {
	URL string `json:"url"`
}

type discordPayload struct {
	Content string         `json:"content,omitempty"`
	Embeds  []discordEmbed `json:"embeds,omitempty"`
}

func NewDiscordProvider(webhookURL string) ProviderMessage {
	return discordMessage{webhookURL: webhookURL}
}

func (d discordMessage) Info(message MessagePayload) error {
	message.Severity = SeverityInfo
	return d.send(message)
}

func (d discordMessage) Error(message MessagePayload) error {
	message.Severity = SeverityError
	return d.send(message)
}

func (d discordMessage) Success(message MessagePayload) error {
	message.Severity = SeveritySuccess
	return d.send(message)
}

func (d discordMessage) Warning(message MessagePayload) error {
	message.Severity = SeverityWarning
	return d.send(message)
}

func (d discordMessage) SendFile(file FilePayload) error {
	var embed discordEmbed
	if file.Type == FileTypeImage || (file.Type == "" && isImageURL(file.URL)) {
		embed = discordEmbed{Image: &discordEmbedImage{URL: file.URL}}
	} else {
		embed = discordEmbed{Description: "📎 **" + file.Name + "**\n" + file.URL}
	}
	return d.push(discordPayload{Embeds: []discordEmbed{embed}})
}

func (d discordMessage) send(message MessagePayload) error {
	return d.push(discordPayload{
		Embeds: []discordEmbed{
			{Description: FormatMessage(message), Color: severityColorDiscord(message.Severity)},
		},
	})
}

func (d discordMessage) push(payload discordPayload) error {
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("discord: marshal payload: %w", err)
	}

	resp, err := http.Post(d.webhookURL, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("discord: send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("discord: unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func severityColorDiscord(s Severity) int {
	switch s {
	case SeverityCritical:
		return 0x9C27B0
	case SeverityError:
		return 0xF44336
	case SeverityWarning:
		return 0xFF9800
	case SeverityInfo:
		return 0x2196F3
	case SeveritySuccess:
		return 0x4CAF50
	default:
		return 0x9E9E9E
	}
}
