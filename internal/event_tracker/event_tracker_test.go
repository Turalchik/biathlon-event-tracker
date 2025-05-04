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

	if info, ok := eventTracker.Competitor2Info[1]; !ok || info == nil || info.Status != Registered {
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

func TestEventTracker_OnStartLine(t *testing.T) {
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

	if err = eventTracker.OnStartLine(competitorID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if info, ok := eventTracker.Competitor2Info[competitorID]; !ok || info == nil || info.Status != OnStartLine {
		t.Errorf("expected competitor with ID %v to be on start line", competitorID)
	}

	if err = eventTracker.OnStartLine(competitorID); !errors.Is(err, ErrStartTimeNotSetForCompetitor) {
		t.Errorf("expected error %v, got %v", ErrCompetitorNotRegistered, err)
	}
}

func TestEventTracker_StartMoving(t *testing.T) {
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

	if err = eventTracker.OnStartLine(competitorID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if err = eventTracker.StartMoving(competitorID, eventTracker.Competitor2Info[competitorID].StartTime+eventTracker.StartDelta+1); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if info, ok := eventTracker.Competitor2Info[competitorID]; !ok || info == nil || info.Mark != NotStarted {
		t.Errorf("expected competitor with ID %v to be with not started mark", competitorID)
	}

	if err = eventTracker.StartMoving(competitorID, eventTracker.Competitor2Info[competitorID].StartTime+eventTracker.StartDelta); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if info, ok := eventTracker.Competitor2Info[competitorID]; !ok || info == nil || info.Status != OnMainLap || info.Mark != NotFinished {
		t.Errorf("expected competitor with ID %v to be started", competitorID)
	}

	if err = eventTracker.StartMoving(competitorID, 0); !errors.Is(err, ErrCompetitorNotOnStartLine) {
		t.Errorf("expected error %v, got %v", ErrCompetitorNotRegistered, err)
	}
}

func TestEventTracker_OnFiringRange(t *testing.T) {
	cfg := &Config{
		Laps:        0,
		LapLen:      0,
		PenaltyLen:  0,
		FiringLines: 2,
		Start:       "[10:00:00.000]",
		StartDelta:  "[10:00:00.000]",
	}

	eventTracker, err := NewEventTracker(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	competitorID := 1
	startTime := 0
	firingRange := 1

	if err = eventTracker.Register(competitorID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if err = eventTracker.SetStartTime(competitorID, startTime); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if err = eventTracker.OnStartLine(competitorID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if err = eventTracker.StartMoving(competitorID, eventTracker.Competitor2Info[competitorID].StartTime); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if err = eventTracker.OnFiringRange(competitorID, eventTracker.FiringLines+1); !errors.Is(err, ErrFiringRangeNotExist) {
		t.Errorf("expected error %v, got %v", ErrFiringRangeNotExist, err)
	}

	if err = eventTracker.OnFiringRange(competitorID, firingRange); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if info, ok := eventTracker.Competitor2Info[competitorID]; !ok || info == nil || info.Status != OnFiringRange {
		t.Errorf("expected competitor with ID %v to be on firing range", competitorID)
	}

	if err = eventTracker.OnFiringRange(competitorID, firingRange); !errors.Is(err, ErrCompetitorNotOnMainLap) {
		t.Errorf("expected error %v, got %v", ErrCompetitorNotOnMainLap, err)
	}
}

func TestEventTracker_HitTarget(t *testing.T) {
	cfg := &Config{
		Laps:        0,
		LapLen:      0,
		PenaltyLen:  0,
		FiringLines: 2,
		Start:       "[10:00:00.000]",
		StartDelta:  "[10:00:00.000]",
	}

	eventTracker, err := NewEventTracker(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	competitorID := 1
	startTime := 0
	firingRange := 1
	target := 1

	if err = eventTracker.Register(competitorID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if err = eventTracker.SetStartTime(competitorID, startTime); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if err = eventTracker.OnStartLine(competitorID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if err = eventTracker.StartMoving(competitorID, eventTracker.Competitor2Info[competitorID].StartTime); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if err = eventTracker.OnFiringRange(competitorID, firingRange); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if err = eventTracker.HitTarget(competitorID, target+5); !errors.Is(err, ErrTargetNotExist) {
		t.Errorf("expected error %v, got %v", ErrTargetNotExist, err)
	}

	if err = eventTracker.HitTarget(competitorID, target); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if info, ok := eventTracker.Competitor2Info[competitorID]; !ok || info == nil || !eventTracker.StateFiringRange[info.FiringRange][target] {
		t.Errorf("expected target with ID %v to be hitted", target)
	}

	if err = eventTracker.HitTarget(competitorID, target); !errors.Is(err, ErrTargetAlreadyHitted) {
		t.Errorf("expected error %v, got %v", ErrTargetAlreadyHitted, err)
	}
}
