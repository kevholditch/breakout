package game

import (
	"github.com/EngoEngine/ecs"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/kevholditch/breakout/internal/pkg/render"
)

type RenderSystem struct {
	width            float32
	height           float32
	entities         map[uint64]holder
	program          *render.Program
	indexBuffer      *render.IndexBuffer
	vertexArray      *render.VertexArray
	vertexBuffer     *render.VertexBuffer
	projectionMatrix mgl32.Mat4
}

type holder struct {
	entity    ecs.BasicEntity
	component *RenderComponent
}

func NewRenderSystem(width, height float32) *RenderSystem {
	return &RenderSystem{
		entities: map[uint64]holder{},
		width:    width,
		height:   height,
	}
}

func (r *RenderSystem) New(world *ecs.World) {
	err := r.initialise()
	if err != nil {
		panic(err)
	}
}

func (r *RenderSystem) initialise() error {
	var rc *RenderComponent
	for _, v := range r.entities {
		rc = v.component
		break
	}

	render.UseDefaultBlending()

	indices := []int32{
		0, 1, 2,
		0, 3, 2,
	}

	r.vertexArray = render.NewVertexArray()
	r.indexBuffer = render.NewIndexBuffer(indices)

	r.projectionMatrix = mgl32.Ortho(0, r.width, 0, r.height, -1.0, 1.0)

	r.vertexBuffer = render.NewVertexBuffer(rc.Quad.ToBuffer())
	r.vertexArray.AddBuffer(r.vertexBuffer, render.NewVertexBufferLayout().AddLayoutFloats(2).AddLayoutFloats(4))

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

func (r *RenderSystem) Update(dt float32) {
	for _, e := range r.entities {
		r.vertexBuffer.Update(e.component.Quad.ToBuffer())
	}

	render.Clear()

	r.program.Bind()
	m := mgl32.Ident4().Mul4(mgl32.Translate3D(0, 0, 0))
	mvp := r.projectionMatrix.Mul4(m)
	r.program.SetUniformMat4f("u_MVP", mvp)

	render.Render(r.vertexArray, r.indexBuffer, r.program)
}

func (r *RenderSystem) Add(entity ecs.BasicEntity, renderComponent *RenderComponent) *RenderSystem {
	r.entities[entity.ID()] = holder{
		entity:    entity,
		component: renderComponent,
	}
	return r
}

func (r *RenderSystem) Remove(_ ecs.BasicEntity) {}
