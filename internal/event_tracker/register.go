package event_tracker

import "fmt"

func (eventTracker *EventTracker) Register(competitorID int) error {
	if _, ok := eventTracker.Competitor2Info[competitorID]; ok {
		return fmt.Errorf("%w: competitor id = %v", ErrCompetitorAreExist, competitorID)
	}
	eventTracker.Competitor2Info[competitorID] = &Info{
		Status:                     Registered,
		Mark:                       NotStarted,
		StartTime:                  0,
		TotalMs2CompleteEachLaps:   make([]int, 0, eventTracker.Laps),
		TotalMs2CompletePenaltyLap: 0,
		NumberHittingTarget:        0,
	}
	return nil
}
