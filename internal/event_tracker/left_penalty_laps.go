package event_tracker

import "fmt"

func (eventTracker *EventTracker) LeftPenaltyLaps(competitorID int, endTime int) error {
	info, ok := eventTracker.Competitor2Info[competitorID]
	if !ok || info == nil || info.Status != EnteredPenaltyLaps {
		return fmt.Errorf("%w: competitor id = %v", ErrCompetitorNotEnteredPenaltyLaps, competitorID)
	}

	info.TotalNumberPenaltyLaps += info.CurrentNumberPenaltyLaps
	info.CurrentNumberPenaltyLaps = 0
	info.TotalMs2CompletePenaltyLap += endTime - info.StartTimeLastPenaltyLap
	info.Status = LeftPenaltyLaps

	return nil
}
