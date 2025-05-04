package event_tracker

import "fmt"

func (eventTracker *EventTracker) LeftFiringRange(competitorID int) error {
	info, ok := eventTracker.Competitor2Info[competitorID]
	if !ok || info == nil || info.Status != OnFiringRange {
		return fmt.Errorf("%w: competitor id = %v", ErrCompetitorNotOnFiringRange, competitorID)
	}

	for _, v := range eventTracker.StateFiringRange[info.FiringRange] {
		if !v {
			info.CurrentNumberPenaltyLaps++
		}
	}
	info.CurrentNumberPenaltyLaps--

	eventTracker.StateFiringRange[info.FiringRange] = IsHittingTarget([6]bool{})
	info.Status = LeftFiringRange
	return nil
}
