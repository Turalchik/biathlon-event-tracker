package event_tracker

import "fmt"

func (eventTracker *EventTracker) competitorInReport(competitorID int, info *Info) string {
	var outStr string
	outStr += fmt.Sprintf("[%s] %v ", Mark2String[info.Mark], competitorID)

	outStr += "["
	for _, ms2CompleteLap := range info.TotalMs2CompleteEachLaps {
		outStr += fmt.Sprintf("{%s, %.3f}, ", MillisecondsToTime(ms2CompleteLap), float64(eventTracker.LapLen)/float64(ms2CompleteLap/1000))
	}

	for i := len(info.TotalMs2CompleteEachLaps); i < eventTracker.Laps; i++ {
		outStr += "{,}, "
	}
	outStr = outStr[:len(outStr)-2]
	outStr += "] "

	var avgSpeedOnPenaltyLaps float64
	if info.TotalNumberPenaltyLaps > 0 {
		avgSpeedOnPenaltyLaps = float64(info.TotalNumberPenaltyLaps*eventTracker.PenaltyLen) / float64(info.TotalMs2CompletePenaltyLap/1000)
	}
	outStr += fmt.Sprintf("{%s, %.3f} ", MillisecondsToTime(info.TotalMs2CompletePenaltyLap), avgSpeedOnPenaltyLaps)
	outStr += fmt.Sprintf("%v/%v\n", info.NumberHitTarget, info.TotalNumberShots)

	return outStr
}
