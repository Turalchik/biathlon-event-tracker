package event_tracker

import "fmt"

func (eventTracker *EventTracker) HitTarget(competitorID int, target int) error {
	info, ok := eventTracker.Competitor2Info[competitorID]
	if !ok || info == nil || info.Status != OnFiringRange {
		return fmt.Errorf("%w: competitor id = %v", ErrCompetitorNotOnFiringRange, competitorID)
	}

	if target < 1 || target > 5 {
		return fmt.Errorf("%w: target id = %v", ErrTargetNotExist, target)
	}

	if eventTracker.StateFiringRange[info.FiringRange][target] {
		return fmt.Errorf("%w: target id = %v", ErrTargetAlreadyHitted, target)
	}

	info.NumberHitTarget++
	eventTracker.StateFiringRange[info.FiringRange][target] = true
	return nil
}
