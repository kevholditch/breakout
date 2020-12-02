package components

import (
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/kevholditch/breakout/internal/pkg/render"
	"math"
)

type Circle struct {
	Position mgl32.Vec2
	Colour   mgl32.Vec4
	Radius   float32
	buffer   []float32
}

func NewCircle(x, y, radius, r, g, b, a float32) *Circle {
	triangleAmount := float32(60)
	twicePi := float32(2.0) * math.Pi

	var buffer []float32
	for i := float32(0); i <= triangleAmount; i++ {
		x1 := float32(0.5) + (radius * float32(math.Cos(float64(i*twicePi/triangleAmount))))
		y1 := float32(0.5) + (radius * float32(math.Sin(float64(i*twicePi/triangleAmount))))
		buffer = append(buffer, x1, y1)
	}

	return &Circle{
		Position: [2]float32{x, y},
		Colour:   [4]float32{r, g, b, a},
		Radius:   radius,
		buffer:   buffer,
	}
}

func NewCircleWithColour(x, y, radius float32, c Colour) *Circle {
	return NewCircle(x, y, radius, c.R, c.G, c.B, c.A)
}

func (c *Circle) ToBuffer() []float32 {
	return c.buffer
}

func (c *Circle) LeftMost() float32 {
	return c.Position.X() - c.Radius
}

func (c *Circle) RightMost() float32 {
	return c.Position.X() + c.Radius
}

func (c *Circle) UpperMost() float32 {
	return c.Position.Y() + c.Radius
}

func (c *Circle) LowerMost() float32 {
	return c.Position.Y() - c.Radius
}

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
