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
	EndedMainLap
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
	StartTimeLastLap           int
	StartTimeLastPenaltyLap    int
	TotalNumberPenaltyLaps     int
	CurrentNumberPenaltyLaps   int
	TotalMs2CompleteEachLaps   []int
	TotalMs2CompletePenaltyLap int
}

var ErrInvalidNumberArguments = errors.New("invalid number arguments")
var ErrInvalidEventID = errors.New("invalid event id")
var ErrInvalidCompetitorID = errors.New("invalid competitor id")
var ErrCompetitorAreExist = errors.New("competitor are exist")
var ErrCompetitorNotRegistered = errors.New("competitor not registered")
var ErrStartTimeNotSetForCompetitor = errors.New("start time not set for competitor")
var ErrCompetitorNotOnStartLine = errors.New("competitor not on start line")
var ErrCompetitorNotOnMainLap = errors.New("competitor not on main lap")
var ErrFiringRangeNotExist = errors.New("firing range not exist")
var ErrFiringRangeNotFree = errors.New("firing range not free")
var ErrCompetitorNotOnFiringRange = errors.New("competitor not on firing range")
var ErrTargetNotExist = errors.New("target not exist")
var ErrTargetAlreadyHit = errors.New("target has already been hit")
var ErrCompetitorNotLeftFiringRange = errors.New("competitor not left firing range")

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
