package event_tracker

type EventTracker struct {
	Laps        int
	LapLen      int
	PenaltyLen  int
	FiringLines int
	Start       int
	StartDelta  int
}

type Config struct {
	Laps        int    `json:"laps"`
	LapLen      int    `json:"lapLen"`
	PenaltyLen  int    `json:"penaltyLen"`
	FiringLines int    `json:"firingLines"`
	Start       string `json:"start"`
	StartDelta  string `json:"startDelta"`
}

func NewEventTracker(cfg *Config) (*EventTracker, error) {
	start, err := TimeToMilliseconds(cfg.Start)
	if err != nil {
		return nil, err
	}

	startDelta, err := TimeToMilliseconds(cfg.StartDelta)
	if err != nil {
		return nil, err
	}

	return &EventTracker{
		Laps:        cfg.Laps,
		LapLen:      cfg.LapLen,
		PenaltyLen:  cfg.PenaltyLen,
		FiringLines: 0,
		Start:       start,
		StartDelta:  startDelta,
	}, nil
}
