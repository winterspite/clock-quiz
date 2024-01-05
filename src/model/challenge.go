package model

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type Challenge struct {
	widget.BaseWidget
	Clock1      fyne.CanvasObject
	Clock1Time  time.Time
	Clock1Input *widget.Entry
	Clock1Guess string

	Clock2      fyne.CanvasObject
	Clock2Time  time.Time
	Clock2Input *widget.Entry
	Clock2Guess string

	DifferenceInput *widget.Entry
	DifferenceGuess string

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
	dialog.ShowInformation("Checking", "Checking your answers", c.Window)
}
