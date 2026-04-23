package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/ukritfongsomboon/DinoNotify/dinonotify"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// ========================================================================== Line

	line := dinonotify.NewLineMessageAPIProvider(
		os.Getenv("LINE_CHANNEL_ACCESS_TOKEN"),
		os.Getenv("LINE_TO"),
	)

	// Text messages
	line.Info(dinonotify.MessagePayload{
		Title:    "Project : worker-job",
		Subtitle: "Stage   : UAT",
		Message:  "Service started successfully",
	})
	line.Warning(dinonotify.MessagePayload{
		Title:    "Project : worker-job",
		Subtitle: "Stage   : UAT",
		Message:  "Queue delay > 10s",
	})
	line.Error(dinonotify.MessagePayload{
		Title:    "Project : Payment API",
		Subtitle: "Stage   : UAT",
		Message:  "Cannot connect to PostgreSQL on port 5432",
	})
	line.Success(dinonotify.MessagePayload{
		Title:    "Project : worker-job",
		Subtitle: "Stage   : UAT",
		Message:  "Queue delay back to normal",
	})

	// File messages
	line.SendFile(dinonotify.FilePayload{
		Name: "image.jpg",
		URL:  "https://www.w3schools.com/html/pic_trulli.jpg",
	})
	line.SendFile(dinonotify.FilePayload{
		Name: "username.csv",
		URL:  "https://support.staffbase.com/hc/en-us/article_attachments/360009197031/username.csv",
	})
	line.SendFile(dinonotify.FilePayload{
		Name: "pdf-sample.pdf",
		URL:  "https://www.rd.usda.gov/sites/default/files/pdf-sample_0.pdf",
	})
	line.SendFile(dinonotify.FilePayload{
		Name:       "video.mp4",
		URL:        "https://www.w3schools.com/tags/mov_bbb.mp4",
		PreviewURL: "https://www.w3schools.com/html/pic_trulli.jpg",
		Type:       dinonotify.FileTypeVideo,
	})
	line.SendFile(dinonotify.FilePayload{
		Name: "Google",
		URL:  "https://google.com",
	})

	// ========================================================================== Slack

	slack := dinonotify.NewSlackProvider(
		os.Getenv("SLACK_WEBHOOK_URL"),
	)

	// Text messages
	slack.Info(dinonotify.MessagePayload{
		Title:    "Project : worker-job",
		Subtitle: "Stage   : UAT",
		Message:  "Service started successfully",
	})
	slack.Warning(dinonotify.MessagePayload{
		Title:    "Project : worker-job",
		Subtitle: "Stage   : UAT",
		Message:  "Queue delay > 10s",
	})
	slack.Error(dinonotify.MessagePayload{
		Title:    "Project : Payment API",
		Subtitle: "Stage   : UAT",
		Message:  "Cannot connect to PostgreSQL on port 5432",
	})
	slack.Success(dinonotify.MessagePayload{
		Title:    "Project : worker-job",
		Subtitle: "Stage   : UAT",
		Message:  "Queue delay back to normal",
	})

	// File messages
	slack.SendFile(dinonotify.FilePayload{
		Name: "image.jpg",
		URL:  "https://www.w3schools.com/html/pic_trulli.jpg",
	})
	slack.SendFile(dinonotify.FilePayload{
		Name: "username.csv",
		URL:  "https://support.staffbase.com/hc/en-us/article_attachments/360009197031/username.csv",
	})
	slack.SendFile(dinonotify.FilePayload{
		Name: "pdf-sample.pdf",
		URL:  "https://www.rd.usda.gov/sites/default/files/pdf-sample_0.pdf",
	})
	slack.SendFile(dinonotify.FilePayload{
		Name:       "video.mp4",
		URL:        "https://www.w3schools.com/tags/mov_bbb.mp4",
		PreviewURL: "https://www.w3schools.com/html/pic_trulli.jpg",
		Type:       dinonotify.FileTypeVideo,
	})
	slack.SendFile(dinonotify.FilePayload{
		Name: "Google",
		URL:  "https://google.com",
	})

	// ========================================================================== Discord

	discord := dinonotify.NewDiscordProvider(
		os.Getenv("DISCORD_WEBHOOK_URL"),
	)

	// Text messages
	discord.Info(dinonotify.MessagePayload{
		Title:    "Project : worker-job",
		Subtitle: "Stage   : UAT",
		Message:  "Service started successfully",
	})
	discord.Warning(dinonotify.MessagePayload{
		Title:    "Project : worker-job",
		Subtitle: "Stage   : UAT",
		Message:  "Queue delay > 10s",
	})
	discord.Error(dinonotify.MessagePayload{
		Title:    "Project : Payment API",
		Subtitle: "Stage   : UAT",
		Message:  "Cannot connect to PostgreSQL on port 5432",
	})
	discord.Success(dinonotify.MessagePayload{
		Title:    "Project : worker-job",
		Subtitle: "Stage   : UAT",
		Message:  "Queue delay back to normal",
	})

	// File messages
	discord.SendFile(dinonotify.FilePayload{
		Name: "image.jpg",
		URL:  "https://www.w3schools.com/html/pic_trulli.jpg",
	})
	discord.SendFile(dinonotify.FilePayload{
		Name: "username.csv",
		URL:  "https://support.staffbase.com/hc/en-us/article_attachments/360009197031/username.csv",
	})
	discord.SendFile(dinonotify.FilePayload{
		Name: "pdf-sample.pdf",
		URL:  "https://www.rd.usda.gov/sites/default/files/pdf-sample_0.pdf",
	})
	discord.SendFile(dinonotify.FilePayload{
		Name:       "video.mp4",
		URL:        "https://www.w3schools.com/tags/mov_bbb.mp4",
		PreviewURL: "https://www.w3schools.com/html/pic_trulli.jpg",
		Type:       dinonotify.FileTypeVideo,
	})
	discord.SendFile(dinonotify.FilePayload{
		Name: "Google",
		URL:  "https://google.com",
	})
}
