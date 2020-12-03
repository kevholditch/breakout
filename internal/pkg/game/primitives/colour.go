package primitives

import "github.com/go-gl/mathgl/mgl32"

type Colour struct {
	R, G, B, A float32
}

func (c Colour) ToVec4() mgl32.Vec4 {
	return [4]float32{c.R, c.G, c.B, c.A}
}
