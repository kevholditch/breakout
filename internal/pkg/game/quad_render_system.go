package game

import (
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/kevholditch/breakout/internal/pkg/game/components"
	"github.com/kevholditch/breakout/internal/pkg/render"
	"github.com/liamg/ecs"
)

type quadEntity struct {
	base       *ecs.Entity
	position   *components.PositionedComponent
	dimensions *components.DimensionComponent
	colour     *components.ColouredComponent
}

type QuadRendererSystem struct {
	quads            []quadEntity
	buffer           []float32
	vertexBuffer     *render.VertexBuffer
	vertexArray      *render.VertexArray
	program          *render.Program
	projectionMatrix mgl32.Mat4
	generator        *TriangleIndexBufferGenerator
}

func NewQuadRenderSystem(windowSize WindowSize) *QuadRendererSystem {
	return &QuadRendererSystem{quads: []quadEntity{},
		buffer:           []float32{},
		program:          NewQuadShaderProgramOrPanic(),
		projectionMatrix: mgl32.Ortho(0, windowSize.Width, 0, windowSize.Height, -1.0, 1.0),
		generator:        NewTriangleIndexBufferGenerator(),
	}
}

func (qr *QuadRendererSystem) New(*ecs.World) {
	render.UseDefaultBlending()
	gl.ClearColor(colourDarkNavy.R, colourDarkNavy.G, colourDarkNavy.B, colourDarkNavy.A)
}

func (qr *QuadRendererSystem) Update(float32) {
	gl.Clear(gl.COLOR_BUFFER_BIT)

	var buffer []float32
	for _, quad := range qr.quads {
		buffer = append(buffer, qr.ComputeBuffer(quad.position, quad.dimensions, quad.colour)...)
	}
	render.NewVertexArray().
		AddBuffer(
			render.NewVertexBuffer(buffer), render.NewVertexBufferLayout().
				AddLayoutFloats(2).AddLayoutFloats(4))
	qr.program.Bind()
	qr.program.SetUniformMat4f("u_MVP", qr.projectionMatrix)

	gl.DrawElements(gl.TRIANGLES, render.NewIndexBuffer(qr.generator.Generate(len(qr.quads))).GetCount(), gl.UNSIGNED_INT, gl.PtrOffset(0))

}

func (qr *QuadRendererSystem) ComputeBuffer(position *components.PositionedComponent, dimensions *components.DimensionComponent, colour *components.ColouredComponent) []float32 {
	return []float32{
		position.X, position.Y, colour.Colour.R, colour.Colour.G, colour.Colour.B, colour.Colour.A,
		position.X + dimensions.Width, position.Y, colour.Colour.R, colour.Colour.G, colour.Colour.B, colour.Colour.A,
		position.X + dimensions.Width, position.Y + dimensions.Height, colour.Colour.R, colour.Colour.G, colour.Colour.B, colour.Colour.A,
		position.X, position.Y + dimensions.Height, colour.Colour.R, colour.Colour.G, colour.Colour.B, colour.Colour.A,
	}
}

func (qr *QuadRendererSystem) Add(entity *ecs.Entity) {

	qr.quads = append(qr.quads, quadEntity{
		base:       entity,
		position:   entity.Component(components.IsPositioned).(*components.PositionedComponent),
		dimensions: entity.Component(components.HasDimensions).(*components.DimensionComponent),
		colour:     entity.Component(components.IsColoured).(*components.ColouredComponent),
	})
}

func (qr *QuadRendererSystem) Remove(entity *ecs.Entity) {
	var del = -1
	for index, e := range qr.quads {
		if e.base.ID() == entity.ID() {
			del = index
			break
		}
	}
	if del >= 0 {
		qr.quads = append(qr.quads[:del], qr.quads[del+1:]...)
	}

}

func (qr *QuadRendererSystem) RequiredTypes() []interface{} {
	return []interface{}{
		components.IsColoured,
		components.IsPositioned,
		components.HasDimensions,
		components.IsQuad,
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
