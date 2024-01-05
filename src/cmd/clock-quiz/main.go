package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/winterspite/clock-quiz/src/model"
)

func main() {
	a := app.New()

	w := a.NewWindow("Clock Quiz")

	quiz := model.NewQuiz(w)

	quiz.NewChallenge()

	w.ShowAndRun()
}
