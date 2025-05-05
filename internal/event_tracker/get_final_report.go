package event_tracker

func (eventTracker *EventTracker) GetFinalReport() string {
	var finalReport string
	for competitorID, info := range eventTracker.Competitor2Info {
		finalReport += eventTracker.competitorInReport(competitorID, info)
	}
	return finalReport
}
