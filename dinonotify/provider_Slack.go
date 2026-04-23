package dinonotify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type slackMessage struct {
	webhookURL string
}

type slackAttachment struct {
	Color string `json:"color"`
	Text  string `json:"text"`
}

type slackImageBlock struct {
	Type     string `json:"type"`
	ImageURL string `json:"image_url"`
	AltText  string `json:"alt_text"`
}

type slackSectionBlock struct {
	Type string          `json:"type"`
	Text slackTextObject `json:"text"`
}

type slackTextObject struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type slackPayload struct {
	Attachments []slackAttachment `json:"attachments,omitempty"`
	Blocks      []any             `json:"blocks,omitempty"`
}

func NewSlackProvider(webhookURL string) ProviderMessage {
	return slackMessage{webhookURL: webhookURL}
}

func (s slackMessage) Info(message MessagePayload) error {
	return s.send(SeverityInfo, message)
}

func (s slackMessage) Error(message MessagePayload) error {
	return s.send(SeverityError, message)
}

func (s slackMessage) Success(message MessagePayload) error {
	return s.send(SeveritySuccess, message)
}

func (s slackMessage) Warning(message MessagePayload) error {
	return s.send(SeverityWarning, message)
}

func (s slackMessage) SendFile(file FilePayload) error {
	var block any
	if file.Type == FileTypeImage || (file.Type == "" && isImageURL(file.URL)) {
		block = slackImageBlock{Type: "image", ImageURL: file.URL, AltText: file.Name}
	} else {
		block = slackSectionBlock{
			Type: "section",
			Text: slackTextObject{Type: "mrkdwn", Text: "📎 *" + file.Name + "*\n" + file.URL},
		}
	}
	return s.push(slackPayload{Blocks: []any{block}})
}

func (s slackMessage) send(severity Severity, message MessagePayload) error {
	return s.push(slackPayload{
		Attachments: []slackAttachment{
			{Color: severityColor(severity), Text: FormatMessage(severity, message)},
		},
	})
}

func (s slackMessage) push(payload slackPayload) error {
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("slack: marshal payload: %w", err)
	}

	resp, err := http.Post(s.webhookURL, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("slack: send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("slack: unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func severityColor(s Severity) string {
	switch s {
	case SeverityCritical:
		return "#9C27B0"
	case SeverityError:
		return "#F44336"
	case SeverityWarning:
		return "#FF9800"
	case SeverityInfo:
		return "#2196F3"
	case SeveritySuccess:
		return "#4CAF50"
	default:
		return "#9E9E9E"
	}
}
