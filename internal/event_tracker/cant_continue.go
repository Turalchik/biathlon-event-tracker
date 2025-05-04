package event_tracker

import "fmt"

func (eventTracker *EventTracker) CantContinue(competitorID int) error {
	info, ok := eventTracker.Competitor2Info[competitorID]
	if !ok || info == nil {
		return fmt.Errorf("%w: competitor id = %v", ErrCompetitorNotRegistered, competitorID)
	}

	if info.Status == CantContinue {
		return fmt.Errorf("%w: competitor id = %v", ErrCompetitorAlreadyCantContinue, competitorID)
	}

	if info.Mark == Finished {
		return fmt.Errorf("%w: competitor id = %v", ErrCompetitorNotStartedOrAlreadyFinished, competitorID)
	}

	info.Mark = NotFinished
	info.Status = CantContinue
	return nil
}
