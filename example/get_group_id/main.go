package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type webhookEvent struct {
	Events []struct {
		Source struct {
			Type    string `json:"type"`
			UserID  string `json:"userId"`
			GroupID string `json:"groupId"`
			RoomID  string `json:"roomId"`
		} `json:"source"`
	} `json:"events"`
}

func main() {
	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		var body webhookEvent
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		for _, event := range body.Events {
			switch event.Source.Type {
			case "group":
				fmt.Printf("Group ID: %s\n", event.Source.GroupID)
			case "room":
				fmt.Printf("Room ID:  %s\n", event.Source.RoomID)
			case "user":
				fmt.Printf("User ID:  %s\n", event.Source.UserID)
			}
		}

		w.WriteHeader(http.StatusOK)
	})

	fmt.Println("Webhook server running on :8080")
	fmt.Println("Set webhook URL to: https://<ngrok-url>/webhook")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
