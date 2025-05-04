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

	if info, ok := eventTracker.Competitor2Info[1]; !ok || info == nil || info.Status != Registred {
		t.Error("expected competitor with ID 1 to be registered")
	}

	err = eventTracker.Register(1)
	if err == nil {
		t.Fatal("expected error")
	}

	if !errors.Is(err, ErrCompetitorAreExist) {
		t.Errorf("expected error %v, got %v", ErrCompetitorAreExist, err)
	}
}

func TestEventTracker_SetStartTime(t *testing.T) {
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

	competitorID := 1
	startTime := 0

	if err = eventTracker.Register(competitorID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if err = eventTracker.SetStartTime(competitorID, startTime); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if info, ok := eventTracker.Competitor2Info[competitorID]; !ok || info == nil || info.Status != StartTimeSet {
		t.Errorf("expected competitor with ID %v to be start time set", competitorID)
	}

	if info, ok := eventTracker.Competitor2Info[competitorID]; !ok || info == nil || info.StartTime != startTime {
		t.Errorf("expected %v time", startTime)
	}

	if err = eventTracker.SetStartTime(competitorID, startTime); !errors.Is(err, ErrCompetitorNotRegistered) {
		t.Errorf("expected error %v, got %v", ErrCompetitorNotRegistered, err)
	}
}
