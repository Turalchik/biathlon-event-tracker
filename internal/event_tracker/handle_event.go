package event_tracker

import (
	"fmt"
	"strconv"
	"strings"
)

func (eventTracker *EventTracker) HandleEvent(eventString string) error {
	sa := strings.Split(eventString, " ")
	if len(sa) < 3 {
		return ErrInvalidNumberArguments
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
	case 2:
		if len(sa) != 4 {
			return ErrInvalidNumberArguments
		}
		if startTime, err := TimeToMilliseconds(sa[3]); err == nil {
			err = eventTracker.SetStartTime(competitorID, startTime)
		}
	case 3:

	}

	return err
}
