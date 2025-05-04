package event_tracker

import "fmt"

func (eventTracker *EventTracker) EndedMainLap(competitorID int, endTime int) error {
	info, ok := eventTracker.Competitor2Info[competitorID]
	if !ok || info == nil || (info.Status != LeftFiringRange && info.Status != LeftPenaltyLaps) {
		return fmt.Errorf("%w: competitor id = %v", ErrCompetitorNotLeftFiringRangeOrPenaltyLaps, competitorID)
	}

	if info.Mark != NotFinished {
		return fmt.Errorf("%w: competitor id = %v", ErrCompetitorNotStartedAlreadyFinished, competitorID)
	}

	info.TotalMs2CompleteEachLaps = append(info.TotalMs2CompleteEachLaps, endTime-info.StartTimeLastMainLap)
	if len(info.TotalMs2CompleteEachLaps) == eventTracker.Laps {
		info.Mark = Finished
	}

	info.Status = OnStartLine
	return nil
}
