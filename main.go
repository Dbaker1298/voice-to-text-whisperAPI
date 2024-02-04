package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	client := whisper.NewClient(whisper.WithKey(os.Getenv("OPEN_API_KEY")))

	response, err := client.TranscribeFile("file.m4a")
	if err != nil {
		log.Fatalf("Error transcribing file: %v", err)
	}

	fmt.Printf("Transcription: %s\n", response.Text)
}
