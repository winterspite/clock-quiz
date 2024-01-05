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

// parseInputTime parses the input string and returns a Golang duration object or error.
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

// parseInputTime parses the input string and returns a Golang time object or error.
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

// fixupClockTime helps to ensure that 13 and 1 are treated the same to make it easier on guesser.
func fixupClockTime(guess, actual time.Time) time.Time {
	if guess == actual {
		return guess
	} else if guess.Add(time.Hour*12) == actual {
		return actual
	}

	return guess
}
