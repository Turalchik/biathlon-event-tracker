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

	var msg string
	switch eventID {
	case 1:
		err = eventTracker.Register(competitorID)
		msg = fmt.Sprintf("%s The competitor(%v) registered\n", sa[0], competitorID)
	case 2:
		if len(sa) != 4 {
			return fmt.Errorf("%w in %s", ErrInvalidTimeFormat, eventString)
		}
		var startTime int
		startTime, err = TimeToMilliseconds(sa[3])
		if err != nil {
			return fmt.Errorf("%w in %s", ErrInvalidTimeFormat, eventString)
		}
		err = eventTracker.SetStartTime(competitorID, startTime)
		msg = fmt.Sprintf("%s The start time for the competitor(%v) was set by a draw to %s\n", sa[0], competitorID, strings.Trim(sa[3], "[] "))
	case 3:
		err = eventTracker.OnStartLine(competitorID)
		msg = fmt.Sprintf("%s The competitor(%v) is on the start line\n", sa[0], competitorID)
	case 4:
		var startTime int
		startTime, err = TimeToMilliseconds(sa[0])
		if err != nil {
			return fmt.Errorf("%w in %s", ErrInvalidTimeFormat, eventString)
		}
		err = eventTracker.StartMoving(competitorID, startTime)
		msg = fmt.Sprintf("%s The competitor(%v) has started\n", sa[0], competitorID)
	case 5:
		if len(sa) != 4 {
			return ErrInvalidNumberArguments
		}
		var firingRange int
		firingRange, err = strconv.Atoi(sa[3])
		if err != nil {
			return fmt.Errorf("%w in %s", ErrInvalidFiringRangeFormat, eventString)
		}
		err = eventTracker.OnFiringRange(competitorID, firingRange)
		msg = fmt.Sprintf("%s The competitor(%v) is on the firing range(%v)\n", sa[0], competitorID, firingRange)
	case 6:
		if len(sa) != 4 {
			return ErrInvalidNumberArguments
		}
		var target int
		target, err = strconv.Atoi(sa[3])
		if err != nil {
			return fmt.Errorf("%w in %s", ErrInvalidFiringRangeFormat, eventString)
		}
		err = eventTracker.HitTarget(competitorID, target)
		msg = fmt.Sprintf("%s The target(%v) by competitor(%v)\n", sa[0], target, competitorID)
	case 7:
		err = eventTracker.LeftFiringRange(competitorID)
		msg = fmt.Sprintf("%s The competitor(%v) left the firing range\n", sa[0], competitorID)
	case 8:
		var startTime int
		startTime, err = TimeToMilliseconds(sa[0])
		if err != nil {
			return fmt.Errorf("%w in %s", ErrInvalidTimeFormat, eventString)
		}
		err = eventTracker.EnteredPenaltyLaps(competitorID, startTime)
		msg = fmt.Sprintf("%s The competitor(%v) entered the penalty laps\n", sa[0], competitorID)
	case 9:
		var endTime int
		endTime, err = TimeToMilliseconds(sa[0])
		if err != nil {
			return fmt.Errorf("%w in %s", ErrInvalidTimeFormat, eventString)
		}
		err = eventTracker.LeftPenaltyLaps(competitorID, endTime)
		msg = fmt.Sprintf("%s The competitor(%v) left the penalty laps\n", sa[0], competitorID)
	case 10:
		var endTime int
		endTime, err = TimeToMilliseconds(sa[0])
		if err != nil {
			return fmt.Errorf("%w in %s", ErrInvalidTimeFormat, eventString)
		}
		err = eventTracker.EndedMainLap(competitorID, endTime)
		msg = fmt.Sprintf("%s The competitor(%v) ended the main lap\n", sa[0], competitorID)
	case 11:
		if len(sa) != 4 {
			return ErrInvalidNumberArguments
		}
		err = eventTracker.CantContinue(competitorID)
		msg = fmt.Sprintf("%s The competitor(%v) can't continue: %s\n", sa[0], competitorID, sa[3])
	}

	if err == nil {
		fmt.Print(msg)
	}
	return err
}
