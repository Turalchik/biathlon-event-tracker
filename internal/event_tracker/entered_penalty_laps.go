package event_tracker

import "fmt"

func (eventTracker *EventTracker) EnteredPenaltyLaps(competitorID int, startTime int) error {
	info, ok := eventTracker.Competitor2Info[competitorID]
	if !ok || info == nil || info.Status != LeftFiringRange {
		return fmt.Errorf("%w: competitor id = %v", ErrCompetitorNotLeftFiringRange, competitorID)
	}

	info.StartTimeLastPenaltyLap = startTime
	info.Status = EnteredPenaltyLaps
	return nil
}
