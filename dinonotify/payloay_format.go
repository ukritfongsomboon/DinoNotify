package dinonotify

import (
	"fmt"
	"time"
)

const separator = "──────────────"

func FormatMessage(severity Severity, p MessagePayload) string {
	e := severity.emoji()
	header := fmt.Sprintf("%s %s (%s)", e, severity, time.Now().Format("2006-01-02 15:04"))
	return fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s", header, separator, p.Title, p.Subtitle, separator, p.Message)
}
