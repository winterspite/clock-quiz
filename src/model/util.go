package model

import (
	"errors"
	"regexp"
	"strconv"
	"time"
)

var (
	timeRegex      = regexp.MustCompile(`^(?P<Hour>[0-9]{1,2}):(?P<Minute>[0-9]{2})`)
	ErrInvalidTime = errors.New("invalid time format, try hh:mm")
)

func parseInputDuration(input string) (time.Duration, error) {
	if !timeRegex.MatchString(input) {
		return 0, ErrInvalidTime
	}

	values := timeRegex.FindStringSubmatch(input)

	if len(values) != 3 {
		return 0, ErrInvalidTime
	}

	hourInt, err := strconv.Atoi(values[1])
	if err != nil {
		return 0, err
	}

	minInt, err := strconv.Atoi(values[2])
	if err != nil {
		return 0, err
	}

	d := (time.Hour * time.Duration(hourInt)) +
		(time.Minute * time.Duration(minInt))

	return d, nil
}

func parseInputTime(input string) (time.Time, error) {
	t := time.Time{}

	if !timeRegex.MatchString(input) {
		return t, ErrInvalidTime
	}

	values := timeRegex.FindStringSubmatch(input)

	if len(values) != 3 {
		return t, ErrInvalidTime
	}

	hourInt, err := strconv.Atoi(values[1])
	if err != nil {
		return t, err
	}

	t = t.Add(time.Hour * time.Duration(hourInt))

	minInt, err := strconv.Atoi(values[2])
	if err != nil {
		return t, err
	}

	t = t.Add(time.Minute * time.Duration(minInt))

	return t, nil
}
