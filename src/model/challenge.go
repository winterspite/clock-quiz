package model

import (
	"crypto/rand"
	"errors"
	"log"
	"math/big"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var (
	ErrInvalidClock1Time = errors.New("clock 1 time is wrong")
	ErrInvalidClock2Time = errors.New("clock 2 time is wrong")
	ErrInvalidDifference = errors.New("difference is wrong")
)

type Challenge struct {
	Quiz *Quiz
	widget.BaseWidget
	Clock1            fyne.CanvasObject
	Clock1Time        time.Time
	Clock1Input       *widget.Entry
	Clock1InputString string
	Clock1Guess       time.Time

	Clock2            fyne.CanvasObject
	Clock2Time        time.Time
	Clock2Input       *widget.Entry
	Clock2InputString string
	Clock2Guess       time.Time

	DifferenceInput       *widget.Entry
	DifferenceInputString string
	DifferenceGuess       time.Duration
	Difference            time.Duration

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

	t1, t2 = fixupTimes(t1, t2)

	c.New(t1, t2)

	return &c
}

// fixupTimes ensures that our times make for good guesses
func fixupTimes(t1, t2 time.Time) (time.Time, time.Time) {
	if t2.Before(t1) {
		t1, t2 = t2, t1
	}

	return t1, t2
}

func newRandomTimes() (int, int, int, int) {
	t1Hour, _ := rand.Int(rand.Reader, big.NewInt(25)) // gets a random integer between 0 and 24
	t2Hour, _ := rand.Int(rand.Reader, big.NewInt(25)) // gets a random integer between 0 and 24
	t1Min, _ := rand.Int(rand.Reader, big.NewInt(13))  // gets a random integer between 0 and 60, on nearest 5
	t2Min, _ := rand.Int(rand.Reader, big.NewInt(13))  // gets a random integer between 0 and 60, on nearest 5

	t1Min = t1Min.Mul(t1Min, big.NewInt(5))
	t2Min = t2Min.Mul(t2Min, big.NewInt(5))

	return int(t1Hour.Int64()), int(t1Min.Int64()),
		int(t2Hour.Int64()), int(t2Min.Int64())
}

// New creates a new challenge and sets all the internal variables.
func (c *Challenge) New(time1, time2 time.Time) {
	c.Clock1Time = time1
	c.Clock1 = CreateClock(time1)

	c.Clock2Time = time2
	c.Clock2 = CreateClock(time2)

	c.Difference = time2.Sub(time1)

	c.Clock1Input = widget.NewEntry()
	c.Clock1Input.SetPlaceHolder("Enter Time (hh:mm)")

	c.Clock2Input = widget.NewEntry()
	c.Clock2Input.SetPlaceHolder("Enter Time (hh:mm)")

	c.DifferenceInput = widget.NewEntry()
	c.DifferenceInput.SetPlaceHolder("Enter difference (hh:mm)")

	c.SubmitButton = widget.NewButton("Check Answer", c.Check)

	log.Printf("T1: %02d:%02d, T2: %02d:%02d, Diff: %v",
		time1.Hour(), time1.Minute(),
		time2.Hour(), time2.Minute(),
		time2.Sub(time1))
}

// CreateClock creates a fyne canvas object representing a clock set to the time passed in.
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

// SyncInputFields syncs the Fyne widget fields into in-struct text fields for better testing.
func (c *Challenge) SyncInputFields() {
	c.Clock1InputString = c.Clock1Input.Text
	c.Clock2InputString = c.Clock2Input.Text
	c.DifferenceInputString = c.DifferenceInput.Text
}

// Check is the button-press function for checking a challenge.
func (c *Challenge) Check() {
	log.Printf("Check: t1: %s, t2: %s, diff: %s",
		c.Clock1Input.Text,
		c.Clock2Input.Text,
		c.DifferenceInput.Text,
	)

	c.SyncInputFields()

	score, err := c.InternalCheck()

	c.UpdateScore(score)

	content := c.getDialogContents(score, err)

	if err != nil {
		dialog.ShowCustom("Incorrect", "OK", content, c.Window)
	} else {
		dialog.ShowCustom("Correct!", "Great Job!", content, c.Window)

		c.Quiz.NewChallenge()
	}
}

func (c *Challenge) getDialogContents(score Score, err error) *fyne.Container {
	c1Label := widget.NewLabel("Clock 1 Guess")
	c1Text := widget.NewLabel(c.Clock1InputString)
	c2Label := widget.NewLabel("Clock 2 Guess")
	c2Text := widget.NewLabel(c.Clock2InputString)
	diffLabel := widget.NewLabel("Difference Guess")
	diffText := widget.NewLabel(c.DifferenceInputString)
	resultLabel := widget.NewLabel("Result")

	result := string(score)
	if err != nil {
		result += " " + err.Error()
	}

	resultString := widget.NewLabel(result)

	dContainer := container.New(layout.NewFormLayout(),
		c1Label, c1Text,
		c2Label, c2Text,
		diffLabel, diffText,
		resultLabel, resultString,
	)

	return dContainer
}

type Score string

const (
	ScoreCorrect   Score = "Correct"
	ScoreIncorrect Score = "Incorrect"
	ScoreInvalid   Score = "Invalid"
)

// UpdateScore updates the associated score for our challenge.
func (c *Challenge) UpdateScore(score Score) {
	switch score {
	case ScoreCorrect:
		c.Quiz.Scoreboard.Correct++
	case ScoreIncorrect:
		c.Quiz.Scoreboard.Incorrect++
	case ScoreInvalid:
		c.Quiz.Scoreboard.Invalid++
	}

	c.Quiz.Scoreboard.UpdateScore()
}

// InternalCheck does the main score comparison loop.
func (c *Challenge) InternalCheck() (Score, error) {
	var err error

	c.Clock1Guess, err = parseInputTime(c.Clock1InputString)
	if err != nil {
		return ScoreInvalid, err
	}

	c.Clock2Guess, err = parseInputTime(c.Clock2InputString)
	if err != nil {
		return ScoreInvalid, err
	}

	c.DifferenceGuess, err = parseInputDuration(c.DifferenceInputString)
	if err != nil {
		return ScoreInvalid, err
	}

	c.Clock1Guess = fixupClockTime(c.Clock1Guess, c.Clock1Time)
	c.Clock2Guess = fixupClockTime(c.Clock2Guess, c.Clock2Time)

	if c.Clock1Time != c.Clock1Guess {
		return ScoreIncorrect, ErrInvalidClock1Time
	}

	if c.Clock2Time != c.Clock2Guess {
		return ScoreIncorrect, ErrInvalidClock2Time
	}

	if c.DifferenceGuess != c.Difference {
		if c.Difference > (time.Hour * 12) { // 24hr check
			if c.DifferenceGuess == (c.Difference - (time.Hour * 12)) {
				return ScoreCorrect, nil
			}
		}

		return ScoreIncorrect, ErrInvalidDifference
	}

	return ScoreCorrect, nil
}
