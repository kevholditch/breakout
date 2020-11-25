package game

import (
	"github.com/EngoEngine/ecs"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/kevholditch/breakout/internal/pkg/render"
	"math"
)

type RenderSystem struct {
	width            float32
	height           float32
	entities         []renderEntityHolder
	program          *render.Program
	projectionMatrix mgl32.Mat4
	generator        *TriangleIndexBufferGenerator
}

type renderEntityHolder struct {
	entity          *ecs.BasicEntity
	renderComponent *RenderComponent
}

func NewRenderSystem(windowSize WindowSize) *RenderSystem {
	return &RenderSystem{
		entities:         []renderEntityHolder{},
		width:            windowSize.Width,
		height:           windowSize.Height,
		projectionMatrix: mgl32.Ortho(0, windowSize.Width, 0, windowSize.Height, -1.0, 1.0),
		generator:        NewTriangleIndexBufferGenerator(),
	}
}

func (r *RenderSystem) New(*ecs.World) {
	err := r.initialise()
	if err != nil {
		panic(err)
	}
}

func (r *RenderSystem) initialise() error {

	render.UseDefaultBlending()
	gl.ClearColor(colourDarkNavy.R, colourDarkNavy.G, colourDarkNavy.B, colourDarkNavy.A)

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
		return err
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
		return err
	}

	program, err := render.NewProgram(vs, fs)
	if err != nil {
		return err
	}

	r.program = program

	return nil
}

func (r *RenderSystem) generateVertexBuffer() []float32 {
	var result []float32

	for _, e := range r.entities {
		result = append(result, e.renderComponent.Quad.ToBuffer()...)
	}

	return result
}

func (r *RenderSystem) Update(float32) {

	// quad buffers
	render.NewVertexArray().
		AddBuffer(
			render.NewVertexBuffer(r.generateVertexBuffer()), render.NewVertexBufferLayout().
				AddLayoutFloats(2).AddLayoutFloats(4))

	gl.Clear(gl.COLOR_BUFFER_BIT)

	r.program.Bind()
	r.program.SetUniformMat4f("u_MVP", r.projectionMatrix)
	gl.DrawElements(gl.TRIANGLES, render.NewIndexBuffer(r.generator.Generate(len(r.entities))).GetCount(), gl.UNSIGNED_INT, gl.PtrOffset(0))

	triangleAmount := float32(60)
	twicePi := float32(2.0) * math.Pi

	var positions []float32
	x := float32(0)
	y := float32(0)
	radius := float32(20)
	for i := float32(0); i <= triangleAmount; i++ {
		x1 := x + (radius * float32(math.Cos(float64(i*twicePi/triangleAmount))))
		y1 := y + (radius * float32(math.Sin(float64(i*twicePi/triangleAmount))))
		positions = append(positions, x1, y1)
	}

	// circle buffers
	mvp := r.projectionMatrix.Mul4(mgl32.Ident4().Mul4(mgl32.Translate3D(200, 200, 0)))
	r.program.SetUniformMat4f("u_MVP", mvp)

	gl.DrawArrays(gl.TRIANGLE_FAN, 0, render.NewVertexArray().
		AddBuffer(render.NewVertexBuffer(positions), render.NewVertexBufferLayout().AddLayoutFloats(2)).GetBufferCount())
}

func (r *RenderSystem) Add(entity *ecs.BasicEntity, renderComponent *RenderComponent) {
	r.entities = append(r.entities, renderEntityHolder{
		entity: entity, renderComponent: renderComponent,
	})
}

func (r *RenderSystem) Remove(basic ecs.BasicEntity) {
	var del = -1
	for index, e := range r.entities {
		if e.entity.ID() == basic.ID() {
			del = index
			break
		}
	}
	if del >= 0 {
		r.entities = append(r.entities[:del], r.entities[del+1:]...)
	}
}
