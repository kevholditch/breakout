package game

import (
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/kevholditch/breakout/internal/pkg/render"
)

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

func NewQuadShaderProgramOrPanic() *render.Program {
	program, err := NewQuadShaderProgram()
	if err != nil {
		panic(err)
	}
	return program
}

func NewQuadShaderProgram() (*render.Program, error) {
	vertex := `#version 410 core

layout(location = 0) in vec4 position;
layout(location = 1) in vec4 color;

uniform mat4 u_MVP;

out vec4 v_Color;

void main()
{
	gl_Position = u_MVP * position;
	v_Color = color;
}`
	vs, err := render.NewShaderFromString(vertex, gl.VERTEX_SHADER)
	if err != nil {
		return nil, err
	}

	fragment := `#version 410 core

layout(location = 0) out vec4 o_Color;

in vec4 v_Color;

void main()
{
	o_Color = v_Color;
}`
	fs, err := render.NewShaderFromString(fragment, gl.FRAGMENT_SHADER)
	if err != nil {
		return nil, err
	}

	program, err := render.NewProgram(vs, fs)
	if err != nil {
		return nil, err
	}
	return program, nil
}
