package event_tracker

type EventTracker struct {
	Laps        int
	LapLen      int
	PenaltyLen  int
	FiringLines int
	Start       int
	StartDelta  int

	IsBusyFiringRange []bool
	StateFiringRange  []IsHittingTarget
	Competitor2Info   map[int]*Info
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
		FiringLines: cfg.FiringLines,
		Start:       start,
		StartDelta:  startDelta,

		IsBusyFiringRange: make([]bool, cfg.FiringLines+1),
		StateFiringRange:  make([]IsHittingTarget, cfg.FiringLines+1),
		Competitor2Info:   make(map[int]*Info),
	}, nil
}
