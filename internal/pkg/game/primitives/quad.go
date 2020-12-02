package primitives

import (
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/kevholditch/breakout/internal/pkg/render"
)

type Quad struct {
	Width, Height float32
	Colour        mgl32.Vec4
}

func NewQuad(w, h, r, g, b, a float32) *Quad {
	return &Quad{
		Width:  w,
		Height: h,
		Colour: [4]float32{r, g, b, a},
	}
}

func NewQuadWithColour(w, h float32, c Colour) *Quad {
	return NewQuad(w, h, c.R, c.G, c.B, c.A)
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
