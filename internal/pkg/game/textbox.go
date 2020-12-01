package game

import "github.com/go-gl/mathgl/mgl32"

type TextBox struct {
	Position mgl32.Vec2
	Text     string
	Colour   Colour
}

func NewTextBox(position mgl32.Vec2, text string, colour Colour) *TextBox {
	return &TextBox{
		Position: position,
		Text:     text,
		Colour:   colour,
	}
}
