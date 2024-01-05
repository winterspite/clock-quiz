package model

import (
	"fmt"

	"fyne.io/fyne/v2/widget"
)

type Scoreboard struct {
	Correct   int
	Incorrect int
	Invalid   int
	Label     *widget.Label
}

func (s *Scoreboard) UpdateScore() {
	s.Label.Text = fmt.Sprintf("Correct: %02d, Incorrect: %02d, Invalid: %02d",
		s.Correct, s.Incorrect, s.Invalid)

	s.Label.Refresh()
}

func NewScoreboard() *Scoreboard {
	s := Scoreboard{
		Label: widget.NewLabel(""),
	}

	s.UpdateScore()

	return &s
}
