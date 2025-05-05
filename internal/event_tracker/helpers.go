package event_tracker

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
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
	ErrCantOpenFile                              = errors.New("can't open file")
	ErrReadingFile                               = errors.New("error reading the file")
	ErrCantDecodeJSON                            = errors.New("can't decode json file")
	ErrInvalidTimeFormat                         = errors.New("invalid time format")
	ErrInvalidFiringRangeFormat                  = errors.New("invalid firing range format")
)

var Mark2String = map[Mark]string{
	NotStarted:  "NotStarted",
	NotFinished: "NotFinished",
	Finished:    "Finished",
}

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

	seconds, err := strconv.Atoi(secParts[0])
	if err != nil || seconds < 0 || seconds >= 60 {
		return 0, fmt.Errorf("invalid seconds")
	}

	var milliseconds int
	if len(secParts) == 2 {
		milliseconds, err = strconv.Atoi(secParts[1])
		if err != nil || milliseconds < 0 || milliseconds >= 1000 {
			return 0, fmt.Errorf("invalid milliseconds")
		}
	}

	totalMs := (hours*3600+minutes*60+seconds)*1000 + milliseconds
	return totalMs, nil
}

func MillisecondsToTime(ms int) string {

	totalSeconds := ms / 1000
	milliseconds := ms % 1000

	hours := totalSeconds / 3600
	remainingSeconds := totalSeconds % 3600
	minutes := remainingSeconds / 60
	seconds := remainingSeconds % 60

	return fmt.Sprintf("%02d:%02d:%02d.%03d",
		hours,
		minutes,
		seconds,
		milliseconds,
	)
}

func ParseConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("%w with error %s", ErrCantOpenFile, err)
	}
	defer file.Close()

	var config Config
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("%w with error %s", ErrCantDecodeJSON, err)
	}

	return &config, nil
}
