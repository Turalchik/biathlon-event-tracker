package event_tracker

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type EventStatus int
type IsHittingTarget [6]bool

const (
	Registered EventStatus = iota + 1
	StartTimeSet
	OnStartLine
	OnMainLap
	OnFiringRange
	LeftFiringRange
	EnteredPenaltyLaps
	LeftPenaltyLaps
	CantContinue
)

type Mark int

const (
	NotStarted Mark = iota + 1
	NotFinished
	Finished
)

type Info struct {
	Status      EventStatus
	Mark        Mark
	StartTime   int
	FiringRange int

	NumberHitTarget            int
	TotalNumberShots           int
	StartTimeLastMainLap       int
	StartTimeLastPenaltyLap    int
	TotalNumberPenaltyLaps     int
	CurrentNumberPenaltyLaps   int
	TotalMs2CompleteEachLaps   []int
	TotalMs2CompletePenaltyLap int
}

var (
	ErrInvalidNumberArguments                    = errors.New("invalid number arguments")
	ErrInvalidEventID                            = errors.New("invalid event id")
	ErrInvalidCompetitorID                       = errors.New("invalid competitor id")
	ErrCompetitorAreExist                        = errors.New("competitor are exist")
	ErrCompetitorNotRegistered                   = errors.New("competitor not registered")
	ErrStartTimeNotSetForCompetitor              = errors.New("start time not set for competitor")
	ErrCompetitorNotOnStartLine                  = errors.New("competitor not on start line")
	ErrCompetitorNotOnMainLap                    = errors.New("competitor not on main lap")
	ErrFiringRangeNotExist                       = errors.New("firing range not exist")
	ErrFiringRangeNotFree                        = errors.New("firing range not free")
	ErrCompetitorNotOnFiringRange                = errors.New("competitor not on firing range")
	ErrTargetNotExist                            = errors.New("target not exist")
	ErrTargetAlreadyHit                          = errors.New("target has already been hit")
	ErrCompetitorNotLeftFiringRange              = errors.New("competitor not left firing range")
	ErrCompetitorNotEnteredPenaltyLaps           = errors.New("competitor not entered penalty laps")
	ErrCompetitorNotLeftFiringRangeOrPenaltyLaps = errors.New("competitor not left firing range or penalty lap")
	ErrCompetitorNotStartedOrAlreadyFinished     = errors.New("competitor already finished")
	ErrCompetitorAlreadyCantContinue             = errors.New("competitor already can't continue")
)

func TimeToMilliseconds(timeStr string) (int, error) {
	cleaned := strings.Trim(timeStr, "[] ")

	parts := strings.Split(cleaned, ":")
	if len(parts) != 3 {
		return 0, fmt.Errorf("invalid time format")
	}

	hours, err := strconv.Atoi(parts[0])
	if err != nil || hours < 0 {
		return 0, fmt.Errorf("invalid hours")
	}

	minutes, err := strconv.Atoi(parts[1])
	if err != nil || minutes < 0 || minutes >= 60 {
		return 0, fmt.Errorf("invalid minutes")
	}

	secParts := strings.Split(parts[2], ".")
	if len(secParts) != 2 || len(secParts[1]) != 3 {
		return 0, fmt.Errorf("invalid seconds format")
	}

	seconds, err := strconv.Atoi(secParts[0])
	if err != nil || seconds < 0 || seconds >= 60 {
		return 0, fmt.Errorf("invalid seconds")
	}

	milliseconds, err := strconv.Atoi(secParts[1])
	if err != nil || milliseconds < 0 || milliseconds >= 1000 {
		return 0, fmt.Errorf("invalid milliseconds")
	}

	totalMs := (hours*3600+minutes*60+seconds)*1000 + milliseconds
	return totalMs, nil
}
