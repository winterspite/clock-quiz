package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/winterspite/clock-quiz/src/model"
)

func main() {
	a := app.New()

	w := a.NewWindow("Clock Quiz")

	quiz := model.NewChallenge(w)

	content := container.New(
		layout.NewGridLayout(2),
		quiz.Clock1, quiz.Clock2,
		quiz.Clock1Input, quiz.Clock2Input,
		quiz.DifferenceInput, quiz.SubmitButton,
	)

	w.SetContent(content)

	w.ShowAndRun()
}
