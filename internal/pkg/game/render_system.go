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
	entities         []renderEntityHolder
	program          *render.Program
	indexBuffer      *render.IndexBuffer
	vertexArray      *render.VertexArray
	vertexBuffer     *render.VertexBuffer
	projectionMatrix mgl32.Mat4
}

type renderEntityHolder struct {
	entity          *ecs.BasicEntity
	renderComponent *RenderComponent
}

func NewRenderSystem(width, height float32) *RenderSystem {
	return &RenderSystem{
		entities: []renderEntityHolder{},
		width:    width,
		height:   height,
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

	r.vertexArray = render.NewVertexArray()
	r.indexBuffer = render.NewIndexBuffer(r.generateIndexBuffer())

	r.projectionMatrix = mgl32.Ortho(0, r.width, 0, r.height, -1.0, 1.0)

	r.vertexBuffer = render.NewVertexBuffer(r.generateVertexBuffer())
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

func (r *RenderSystem) generateIndexBuffer() []int32 {
	result := make([]int32, 6*len(r.entities))

	for i := int32(0); i < int32(len(r.entities)); i++ {
		result = append(result, i*4)
		result = append(result, i*4+1)
		result = append(result, i*4+2)
		result = append(result, i*4)
		result = append(result, i*4+3)
		result = append(result, i*4+2)
	}

	return result
}

func (r *RenderSystem) generateVertexBuffer() []float32 {
	result := make([]float32, render.QuadBufferSize*len(r.entities))

	for _, e := range r.entities {
		result = append(result, e.renderComponent.Quad.ToBuffer()...)
	}

	return result
}

func (r *RenderSystem) Update(float32) {
	r.indexBuffer.Update(r.generateIndexBuffer())
	for _, e := range r.entities {
		r.vertexBuffer.Update(e.renderComponent.Quad.ToBuffer())
	}

	render.Clear()

	r.program.Bind()
	m := mgl32.Ident4().Mul4(mgl32.Translate3D(0, 0, 0))
	mvp := r.projectionMatrix.Mul4(m)
	r.program.SetUniformMat4f("u_MVP", mvp)

	render.Render(r.vertexArray, r.indexBuffer, r.program)
}

func (r *RenderSystem) Add(entity *ecs.BasicEntity, renderComponent *RenderComponent) *RenderSystem {
	r.entities = append(r.entities, renderEntityHolder{
		entity: entity, renderComponent: renderComponent,
	})
	return r
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
