package model

import "fyne.io/fyne/v2"

type Quiz struct {
	Challenge  *Challenge
	Scoreboard *Scoreboard
}

func NewQuiz(w fyne.Window) Quiz {
	q := Quiz{
		Challenge:  NewChallenge(w),
		Scoreboard: NewScoreboard(),
	}

	q.Challenge.Quiz = &q

	return q
}
