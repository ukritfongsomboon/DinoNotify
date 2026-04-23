package dinonotify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const lineMessagingAPIEndpoint = "https://api.line.me/v2/bot/message/push"

type lineMessage struct {
	channelAccessToken string
	to                 string
}

type lineTextMessage struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type lineImageMessage struct {
	Type               string `json:"type"`
	OriginalContentURL string `json:"originalContentUrl"`
	PreviewImageURL    string `json:"previewImageUrl"`
}

type lineVideoMessage struct {
	Type               string `json:"type"`
	OriginalContentURL string `json:"originalContentUrl"`
	PreviewImageURL    string `json:"previewImageUrl"`
}

type linePushRequest struct {
	To       string `json:"to"`
	Messages []any  `json:"messages"`
}

func NewLineMessageAPIProvider(channelAccessToken string, to string) ProviderMessage {
	return lineMessage{
		channelAccessToken: channelAccessToken,
		to:                 to,
	}
}

func (l lineMessage) Info(message MessagePayload) error {
	message.Severity = SeverityInfo
	return l.send(message)
}

func (l lineMessage) Error(message MessagePayload) error {
	message.Severity = SeverityError
	return l.send(message)
}

func (l lineMessage) Success(message MessagePayload) error {
	message.Severity = SeveritySuccess
	return l.send(message)
}

func (l lineMessage) Warning(message MessagePayload) error {
	message.Severity = SeverityWarning
	return l.send(message)
}

func (l lineMessage) SendFile(file FilePayload) error {
	var msg any
	if file.Type == FileTypeVideo {
		if file.PreviewURL == "" {
			return fmt.Errorf("line: video requires PreviewURL")
		}
		msg = lineVideoMessage{Type: "video", OriginalContentURL: file.URL, PreviewImageURL: file.PreviewURL}
	} else if file.Type == FileTypeImage || (file.Type == "" && isImageURL(file.URL)) {
		preview := file.PreviewURL
		if preview == "" {
			preview = file.URL
		}
		msg = lineImageMessage{Type: "image", OriginalContentURL: file.URL, PreviewImageURL: preview}
	} else {
		msg = lineTextMessage{Type: "text", Text: "📎 " + file.Name + "\n" + file.URL}
	}
	return l.push(msg)
}

func isImageURL(url string) bool {
	for _, ext := range []string{".jpg", ".jpeg", ".png", ".gif", ".webp"} {
		if len(url) >= len(ext) && url[len(url)-len(ext):] == ext {
			return true
		}
	}
	return false
}

func (l lineMessage) send(message MessagePayload) error {
	return l.push(lineTextMessage{Type: "text", Text: FormatMessage(message)})
}

func (l lineMessage) push(msg any) error {
	jsonBody, err := json.Marshal(linePushRequest{To: l.to, Messages: []any{msg}})
	if err != nil {
		return fmt.Errorf("line: marshal payload: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, lineMessagingAPIEndpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("line: create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+l.channelAccessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("line: send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("line: unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
