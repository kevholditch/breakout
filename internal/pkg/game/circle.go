package game

import "github.com/go-gl/mathgl/mgl32"

type Circle struct {
	Position mgl32.Vec2
	Colour   mgl32.Vec4
}

func NewCircle(x, y, r, g, b, a float32) *Circle {
	return &Circle{
		Position: [2]float32{x, y},
		Colour:   [4]float32{r, g, b, a},
	}
}

func NewCircleWithColour(x, y float32, c Colour) *Circle {
	return NewCircle(x, y, c.R, c.G, c.B, c.A)
}

func (c *Circle) ToBuffer() [2]float32 {
	return c.Position
}
