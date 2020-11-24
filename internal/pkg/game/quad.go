package game

import "github.com/go-gl/mathgl/mgl32"

/*
  X,W  ------- Z,W
  |             |
  X,Y -------  Z,Y

*/

const QuadBufferSize = 24

type Quad struct {
	Position mgl32.Vec4
	Colour   mgl32.Vec4
}

func NewQuad(x, y, w, h, r, g, b, a float32) *Quad {
	return &Quad{
		Position: [4]float32{x, y, x + w, y + h},
		Colour:   [4]float32{r, g, b, a},
	}
}

func NewQuadWithColour(x, y, w, h float32, c Colour) *Quad {
	return NewQuad(x, y, w, h, c.R, c.G, c.B, c.A)
}

func (q *Quad) Width() float32 {
	return q.Position.Z() - q.Position.X()
}

func (q *Quad) ToBuffer() []float32 {
	return []float32{
		q.Position.X(), q.Position.Y(), q.Colour.X(), q.Colour.Y(), q.Colour.Z(), q.Colour.W(),
		q.Position.Z(), q.Position.Y(), q.Colour.X(), q.Colour.Y(), q.Colour.Z(), q.Colour.W(),
		q.Position.Z(), q.Position.W(), q.Colour.X(), q.Colour.Y(), q.Colour.Z(), q.Colour.W(),
		q.Position.X(), q.Position.W(), q.Colour.X(), q.Colour.Y(), q.Colour.Z(), q.Colour.W(),
	}
}
