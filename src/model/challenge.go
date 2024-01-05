package model

import (
	"errors"
	"regexp"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var (
	ErrInvalidClock1Time = errors.New("clock 1 time is wrong")
	ErrInvalidClock2Time = errors.New("clock 2 time is wrong")
	ErrInvalidDifference = errors.New("difference is wrong")
)

type Challenge struct {
	widget.BaseWidget
	Clock1      fyne.CanvasObject
	Clock1Time  time.Time
	Clock1Input *widget.Entry
	Clock1Guess time.Time

	Clock2      fyne.CanvasObject
	Clock2Time  time.Time
	Clock2Input *widget.Entry
	Clock2Guess time.Time

	DifferenceInput *widget.Entry
	DifferenceGuess time.Time

	SubmitButton *widget.Button
	Window       fyne.Window
}

func NewChallenge(win fyne.Window) *Challenge {
	c := Challenge{
		Window: win,
	}
	c.New()

	return &c
}

func (c *Challenge) New() {
	c.Clock1 = CreateClock(time.Date(2020, 01, 01, 11, 00, 00, 00, time.UTC))
	c.Clock2 = CreateClock(time.Date(2020, 01, 01, 12, 30, 00, 00, time.UTC))

	c.Clock1Input = widget.NewEntry()
	c.Clock1Input.SetPlaceHolder("Enter Time")

	c.Clock2Input = widget.NewEntry()
	c.Clock2Input.SetPlaceHolder("Enter Time")

	c.DifferenceInput = widget.NewEntry()
	c.DifferenceInput.SetPlaceHolder("Enter difference")

	c.SubmitButton = widget.NewButton("Check", c.Check)
}

func CreateClock(clockTime time.Time) fyne.CanvasObject {
	clock := &ClockLayout{
		Time: clockTime,
	}

	content := clock.Render()
	listener := make(chan fyne.Settings)

	fyne.CurrentApp().Settings().AddChangeListener(listener)

	go func() {
		for {
			settings := <-listener
			clock.ApplyTheme(settings)
		}
	}()

	return content
}

func (c *Challenge) Check() {
	err := c.InternalCheck()
	if err != nil {
		dialog.ShowError(err, c.Window)
	} else {
		dialog.ShowInformation("Correct!", "Great Job!", c.Window)
	}
}

func (c *Challenge) InternalCheck() error {
	var err error

	c.Clock1Guess, err = parseInputTime(c.Clock1Input.Text)
	if err != nil {
		return err
	}

	c.Clock2Guess, err = parseInputTime(c.Clock2Input.Text)
	if err != nil {
		return err
	}

	c.DifferenceGuess, err = parseInputTime(c.DifferenceInput.Text)
	if err != nil {
		return err
	}

	// TODO: Compare the times here.

	return nil
}

var (
	timeRegex      = regexp.MustCompile(`^(?P<Hour>[0-9]+):(?P<Minute>[0-9]+)`)
	ErrInvalidTime = errors.New("invalid time format, try hh:mm")
)

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
