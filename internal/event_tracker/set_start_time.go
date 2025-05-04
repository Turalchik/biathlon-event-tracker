package event_tracker

import "fmt"

func (eventTracker *EventTracker) SetStartTime(competitorID int, startTime int) error {
	info, ok := eventTracker.Competitor2Info[competitorID]
	if !ok || info.Status != Registred {
		return fmt.Errorf("%w: competitor id = %v", ErrCompetitorNotRegistered, competitorID)
	}
	info.StartTime = startTime
	info.Status = StartTimeSet

	return nil
}
