package event_tracker

import "fmt"

func (eventTracker *EventTracker) OnFiringRange(competitorID int, firingRange int) error {
	info, ok := eventTracker.Competitor2Info[competitorID]
	if !ok || info == nil || info.Status != OnMainLap {
		return fmt.Errorf("%w: competitor id = %v", ErrCompetitorNotOnMainLap, competitorID)
	}

	if firingRange < 1 || firingRange > eventTracker.FiringLines {
		return fmt.Errorf("%w: firing range id = %v", ErrFiringRangeNotExist, firingRange)
	}

	if eventTracker.IsBusyFiringRange[firingRange] {
		return fmt.Errorf("%w: firing range id = %v", ErrFiringRangeNotFree, firingRange)
	}

	info.Status = OnFiringRange
	info.FiringRange = firingRange
	info.TotalNumberShots += 5
	eventTracker.IsBusyFiringRange[firingRange] = true
	return nil
}
