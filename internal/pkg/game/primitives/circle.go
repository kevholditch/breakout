package primitives

import (
	"github.com/go-gl/gl/all-core/gl"
	"github.com/kevholditch/breakout/internal/pkg/render"
)

func NewCircleShaderProgramOrPanic() *render.Program {
	program, err := NewCircleShaderProgram()
	if err != nil {
		panic(err)
	}
	return program
}

func NewCircleShaderProgram() (*render.Program, error) {
	vertex := `#version 410 core

layout(location = 0) in vec4 position;

uniform mat4 u_MVP;

void main()
{
	gl_Position = u_MVP * position;
}`
	vs, err := render.NewShaderFromString(vertex, gl.VERTEX_SHADER)
	if err != nil {
		return nil, err
	}

	fragment := `#version 410 core

layout(location = 0) out vec4 o_Color;

uniform vec4 u_Colour;

void main()
{
	o_Color = u_Colour;
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
