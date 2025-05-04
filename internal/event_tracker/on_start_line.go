package event_tracker

import "fmt"

func (eventTracker *EventTracker) OnStartLine(competitorID int) error {
	info, ok := eventTracker.Competitor2Info[competitorID]
	if !ok || info == nil || info.Status != StartTimeSet {
		return fmt.Errorf("%w: competitor id = %v", ErrStartTimeNotSetForCompetitor, competitorID)
	}

	info.Status = OnStartLine
	return nil
}
