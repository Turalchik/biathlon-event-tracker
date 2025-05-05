package main

import (
	"github.com/Turalchik/biathlon-event-tracker/internal/event_tracker"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	godotenv.Load()

	configFilename := os.Getenv("CONFIG_FILENAME")
	if configFilename == "" {
		log.Fatal("CONFIG_FILENAME environment variable is not set")
	}

	eventsFilename := os.Getenv("EVENTS_FILENAME")
	if eventsFilename == "" {
		log.Fatal("EVENTS_FILENAME environment variable is not set")
	}

	config, err := event_tracker.ParseConfig(configFilename)
	if err != nil {
		log.Fatalf("Error parsing config: %v\n", err)
	}

	eventTracker, err := event_tracker.NewEventTracker(config)
	if err != nil {
		log.Fatal("Can't initialize the event tracker")
	}

	if err = eventTracker.HandleEventsFromFile(eventsFilename); err != nil {
		log.Printf("Can't handle events from file with error %s\n", err)
	}
}
