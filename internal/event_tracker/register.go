package event_tracker

import "fmt"

func (eventTracker *EventTracker) Register(competitorID int) error {
	if _, ok := eventTracker.Competitor2Status[competitorID]; ok {
		return fmt.Errorf("%w: competitor id = %v", ErrCompetitorAreExist, competitorID)
	}
	eventTracker.Competitor2Status[competitorID] = Registred
	return nil
}
