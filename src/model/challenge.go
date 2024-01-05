package model

import (
	"errors"
	"math/rand"
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
	DifferenceGuess time.Duration
	Difference      time.Duration

	SubmitButton *widget.Button
	Window       fyne.Window
}

func NewChallenge(win fyne.Window) *Challenge {
	c := Challenge{
		Window: win,
	}

	t1Hour, t1Min, t2Hour, t2Min := newRandomTimes()

	t1 := time.Date(1, 1, 1, t1Hour, t1Min, 00, 00, time.UTC)
	t2 := time.Date(1, 1, 1, t2Hour, t2Min, 00, 00, time.UTC)

	if t2.Before(t1) {
		diff := t1.Sub(t2)

		t2 = t2.Add(diff)
		t2 = t2.Add(diff)
	}

	c.New(t1, t2)

	return &c
}

func newRandomTimes() (t1Hour, t1Min, t2Hour, t2Min int) {
	t1Hour = rand.Intn(25)    // gets a random integer between 0 and 24
	t2Hour = rand.Intn(25)    // gets a random integer between 0 and 24
	t1Min = rand.Intn(13) * 5 // gets a random integer between 0 and 60, on nearest 5
	t2Min = rand.Intn(13) * 5 // gets a random integer between 0 and 60, on nearest 5

	return
}

func (c *Challenge) New(time1, time2 time.Time) {
	c.Clock1Time = time1
	c.Clock1 = CreateClock(time1)

	c.Clock2Time = time2
	c.Clock2 = CreateClock(time2)

	c.Difference = time2.Sub(time1)

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

	c.DifferenceGuess, err = parseInputDuration(c.DifferenceInput.Text)
	if err != nil {
		return err
	}

	if c.Clock1Time != c.Clock1Guess {
		return ErrInvalidClock1Time
	}

	if c.Clock2Time != c.Clock2Guess {
		return ErrInvalidClock2Time
	}

	if c.DifferenceGuess != c.Difference {
		return ErrInvalidDifference
	}

	return nil
}

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
