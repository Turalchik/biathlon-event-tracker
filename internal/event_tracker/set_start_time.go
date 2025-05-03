package event_tracker

import "fmt"

func (eventTracker *EventTracker) SetStartTime(competitorID int, startTime int) error {
	if eventStatus, ok := eventTracker.Competitor2Status[competitorID]; !ok || eventStatus != Registred {
		return fmt.Errorf("%w: competitor id = %v", ErrCompetitorNotRegistered, competitorID)
	}

	eventTracker.Competitor2StartTime[competitorID] = startTime
	eventTracker.Competitor2Status[competitorID] = StartTimeSet

	return nil
}
