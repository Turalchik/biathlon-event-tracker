package event_tracker

import "fmt"

func (eventTracker *EventTracker) StartMoving(competitorID int, startTime int) error {
	info, ok := eventTracker.Competitor2Info[competitorID]
	if !ok || info == nil || info.Status != OnStartLine {
		return fmt.Errorf("%w: competitor id = %v", ErrCompetitorNotOnStartLine, competitorID)
	}

	if startTime >= info.StartTime && startTime <= info.StartTime+eventTracker.StartDelta {
		info.Mark = NotFinished
		info.Status = OnMainLap
		info.StartTimeLastMainLap = startTime
	}

	return nil
}
