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

	quiz := model.NewQuiz(w)

	content := container.New(
		layout.NewFormLayout(),
		quiz.Challenge.Clock1, quiz.Challenge.Clock2,
		quiz.Challenge.Clock1Input, quiz.Challenge.Clock2Input,
		quiz.Challenge.DifferenceInput, quiz.Challenge.SubmitButton,
		quiz.Scoreboard.Label, quiz.Scoreboard.Label,
	)

	w.SetContent(content)

	w.ShowAndRun()
}
