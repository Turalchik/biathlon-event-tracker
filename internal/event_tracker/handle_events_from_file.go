package event_tracker

import (
	"bufio"
	"fmt"
	"os"
)

func (eventTracker *EventTracker) HandleEventsFromFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("%w with error %s", ErrCantOpenFile, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		eventString := scanner.Text()
		if err = eventTracker.HandleEvent(eventString); err != nil {
			return fmt.Errorf("%w with error %s", ErrReadingFile, err.Error())
		}
	}

	return nil
}
