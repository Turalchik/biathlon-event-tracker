package event_tracker

import (
	"errors"
	"testing"
)

func TestEventTracker_Register(t *testing.T) {
	cfg := &Config{
		Laps:        0,
		LapLen:      0,
		PenaltyLen:  0,
		FiringLines: 0,
		Start:       "[10:00:00.000]",
		StartDelta:  "[10:00:00.000]",
	}

	eventTracker, err := NewEventTracker(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	err = eventTracker.Register(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if status, ok := eventTracker.Competitor2Status[1]; !ok || status != Registred {
		t.Errorf("expected competitor with ID 1 to be registered, got status: %v", status)
	}

	err = eventTracker.Register(1)
	if err == nil {
		t.Fatal("expected error")
	}

	if !errors.Is(err, ErrCompetitorAreExist) {
		t.Errorf("expected error %v, got %v", ErrCompetitorAreExist, err)
	}
}
