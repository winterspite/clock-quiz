package model

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Quiz struct {
	Challenge          *Challenge
	Scoreboard         *Scoreboard
	Window             fyne.Window
	NewChallengeButton *widget.Button
}

func NewQuiz(w fyne.Window) Quiz {
	q := Quiz{
		Window:     w,
		Challenge:  NewChallenge(w),
		Scoreboard: NewScoreboard(),
	}

	q.NewChallengeButton = widget.NewButton("New Challenge", q.NewChallenge)

	q.Challenge.Quiz = &q

	return q
}

// NewChallenge sets everything up for the next challenge.
func (q *Quiz) NewChallenge() {
	q.Challenge = NewChallenge(q.Window)

	q.Challenge.DifferenceInput.Text = "Enter difference"
	q.Challenge.DifferenceInput.Refresh()

	q.Challenge.Clock1Input.Text = "Enter Time"
	q.Challenge.Clock1Input.Refresh()
	q.Challenge.Clock1.Refresh()
	q.Challenge.Clock2Input.Text = "Enter Time"
	q.Challenge.Clock2Input.Refresh()
	q.Challenge.Clock2.Refresh()

	q.Challenge.Quiz = q
}
