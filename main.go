package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const githubEventsURL = "https://api.github.com/users/%s/events"

func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage: ./ghactivity <username>")
	}

	username := os.Args[1]
	resp, err := http.Get(fmt.Sprintf(githubEventsURL, username))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode > 299 {
		log.Fatalf("github api returned status %d", resp.StatusCode)
	}

	var events []json.RawMessage
	err = json.Unmarshal(body, &events)
	if err != nil {
		log.Fatal(err)
	}

	for i, event := range events {
		pretty, err := json.MarshalIndent(event, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Event %d:\n%s\n\n", i+1, pretty)
	}
}
