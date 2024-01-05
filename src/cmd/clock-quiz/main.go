package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/winterspite/clock-quiz/src/model"
)

func main() {
	a := app.New()

	w := a.NewWindow("Clock Quiz")

	clock1 := CreateClock(time.Date(2020, 01, 01, 11, 00, 00, 00, time.UTC))

	clock1Input := widget.NewEntry()
	clock1Input.SetPlaceHolder("Enter Time")

	clock2 := CreateClock(time.Date(2020, 01, 01, 12, 30, 00, 00, time.UTC))

	clock2Input := widget.NewEntry()
	clock2Input.SetPlaceHolder("Enter Time")

	diffInput := widget.NewEntry()
	diffInput.SetPlaceHolder("Enter difference")

	submit := widget.NewButton("Check", CheckButton)

	content := container.New(
		layout.NewGridLayout(2),
		clock1, clock2,
		clock1Input, clock2Input,
		diffInput, submit,
	)

	w.SetContent(content)

	w.ShowAndRun()
}

func CreateClock(clockTime time.Time) fyne.CanvasObject {
	clock := &model.ClockLayout{
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

func CheckButton() {

}
