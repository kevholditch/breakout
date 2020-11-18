package render

import "github.com/go-gl/mathgl/mgl32"

/*
  X,W  ------- Z,W
  |             |
  X,Y -------  Z,Y

*/

type Quad struct {
	Position mgl32.Vec4
	Colour   mgl32.Vec4
}

func NewQuad(position, colour mgl32.Vec4) *Quad {
	return &Quad{
		Position: position,
		Colour:   colour,
	}
}
