package model

import (
	"math"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

type ClockLayout struct {
	Hour      *canvas.Line
	Minute    *canvas.Line
	Pips      [12]*canvas.Line
	HourDot   *canvas.Circle
	SecondDot *canvas.Circle
	Face      *canvas.Circle
	Canvas    fyne.CanvasObject
}

func (c *ClockLayout) Rotate(hand *canvas.Line, middle fyne.Position, facePosition float64, offset float32, length float32) {
	rotation := math.Pi * 2 / 60 * facePosition
	x2 := length * float32(math.Sin(rotation))
	y2 := -length * float32(math.Cos(rotation))

	offX := float32(0)
	offY := float32(0)

	if offset > 0 {
		offX += offset * float32(math.Sin(rotation))
		offY += -offset * float32(math.Cos(rotation))
	}

	hand.Position1 = fyne.NewPos(middle.X+offX, middle.Y+offY)
	hand.Position2 = fyne.NewPos(middle.X+offX+x2, middle.Y+offY+y2)
}

func (c *ClockLayout) Layout(_ []fyne.CanvasObject, size fyne.Size) {
	diameter := fyne.Min(size.Width, size.Height)
	radius := diameter / 2
	dotRadius := radius / 12

	stroke := diameter / 40
	midStroke := diameter / 90
	smallStroke := diameter / 200

	size = fyne.NewSize(diameter, diameter)
	middle := fyne.NewPos(size.Width/2, size.Height/2)
	topLeft := fyne.NewPos(middle.X-radius, middle.Y-radius)

	c.Face.Resize(size)
	c.Face.Move(topLeft)

	c.Hour.StrokeWidth = stroke
	c.Rotate(c.Hour, middle, float64((time.Now().Hour()%12)*5)+(float64(time.Now().Minute())/12), dotRadius, radius/2)

	c.Minute.StrokeWidth = midStroke
	c.Rotate(c.Minute, middle, float64(time.Now().Minute()), dotRadius, radius*0.9)

	c.HourDot.StrokeWidth = stroke
	c.HourDot.Resize(fyne.NewSize(dotRadius*2, dotRadius*2))
	c.HourDot.Move(fyne.NewPos(middle.X-dotRadius, middle.Y-dotRadius))

	for idx, p := range c.Pips {
		c.Rotate(p, middle, float64((idx)*5), radius/8*7, radius/8)
		p.StrokeWidth = smallStroke
	}
}

func (c *ClockLayout) MinSize(_ []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(200, 200)
}

func (c *ClockLayout) Render() *fyne.Container {
	c.HourDot = &canvas.Circle{StrokeColor: theme.ForegroundColor(), StrokeWidth: 5}
	c.Face = &canvas.Circle{StrokeColor: theme.DisabledColor(), StrokeWidth: 1}

	c.Hour = &canvas.Line{StrokeColor: theme.ForegroundColor(), StrokeWidth: 5}
	c.Minute = &canvas.Line{StrokeColor: theme.ForegroundColor(), StrokeWidth: 3}

	container := container.NewWithoutLayout(c.HourDot)

	for idx := range c.Pips {
		pip := &canvas.Line{StrokeColor: theme.DisabledColor(), StrokeWidth: 1}
		container.Add(pip)
		c.Pips[idx] = pip
	}

	container.Objects = append(container.Objects, c.Face, c.Hour, c.Minute)
	container.Layout = c

	c.Canvas = container

	return container
}

func (c *ClockLayout) ApplyTheme(_ fyne.Settings) {
	c.HourDot.StrokeColor = theme.ForegroundColor()
	c.Face.StrokeColor = theme.DisabledColor()

	c.Hour.StrokeColor = theme.ForegroundColor()
	c.Minute.StrokeColor = theme.ForegroundColor()

	for _, pip := range c.Pips {
		pip.StrokeColor = theme.DisabledColor()
	}
}
