package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/winterspite/clock-quiz/src/model"
)

func main() {
	a := app.New()

	w := a.NewWindow("Clock Quiz")

	w.SetContent(Show(w))

	w.ShowAndRun()
}

func Show(_ fyne.Window) fyne.CanvasObject {
	clock := &model.ClockLayout{}
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
