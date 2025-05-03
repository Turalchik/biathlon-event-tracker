package event_tracker

import (
	"fmt"
	"strconv"
	"strings"
)

func (eventTracker *EventTracker) HandleEvent(eventString string) error {
	sa := strings.Split(eventString, " ")
	if len(sa) < 3 {
		return ErrNotEnoughArguments
	}

	eventID, err := strconv.Atoi(sa[1])
	if err != nil || eventID > 11 || eventID < 1 {
		return fmt.Errorf("%w: event id = %v", ErrInvalidEventID, eventID)
	}

	competitorID, err := strconv.Atoi(sa[2])
	if err != nil || competitorID < 1 {
		return fmt.Errorf("%w: competitor id = %v", ErrInvalidCompetitorID, competitorID)
	}

	switch eventID {
	case 1:
		err = eventTracker.Register(competitorID)
	}

	return err
}
