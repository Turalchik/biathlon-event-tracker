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
		err = eventTracker.OnStartLine(competitorID)
	case 4:
		if startTime, err := TimeToMilliseconds(sa[0]); err == nil {
			err = eventTracker.StartMoving(competitorID, startTime)
		}
	case 5:
		if len(sa) != 4 {
			return ErrInvalidNumberArguments
		}
		if firingRange, err := strconv.Atoi(sa[3]); err == nil {
			err = eventTracker.OnFiringRange(competitorID, firingRange)
		}
	case 6:
		if len(sa) != 4 {
			return ErrInvalidNumberArguments
		}
		if target, err := strconv.Atoi(sa[3]); err == nil {
			err = eventTracker.HitTarget(competitorID, target)
		}
	case 7:
		err = eventTracker.LeftFiringRange(competitorID)
	case 8:
		if startTime, err := TimeToMilliseconds(sa[0]); err == nil {
			err = eventTracker.EnteredPenaltyLaps(competitorID, startTime)
		}
	case 9:
		if endTime, err := TimeToMilliseconds(sa[0]); err == nil {
			err = eventTracker.LeftPenaltyLaps(competitorID, endTime)
		}
	case 10:
		if endTime, err := TimeToMilliseconds(sa[0]); err == nil {
			err = eventTracker.EndedMainLap(competitorID, endTime)
		}
	}

	return err
}
