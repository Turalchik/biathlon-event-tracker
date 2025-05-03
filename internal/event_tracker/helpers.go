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
	Registred EventStatus = iota + 1
	StartTimeSet
	OnStartLine
	Started
	OnFiringRange
	TargetHit
	LeftFiringRange
	EnteredPenaltyLaps
	LeftPenaltyLaps
	EndedMainLap
	CantContinue
)

var ErrNotEnoughArguments = errors.New("not enough arguments")
var ErrInvalidEventID = errors.New("invalid event id")
var ErrInvalidCompetitorID = errors.New("invalid competitor id")
var ErrCompetitorAreExist = errors.New("competitor are exist")

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
