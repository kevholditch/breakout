package game

import (
	"github.com/EngoEngine/ecs"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/kevholditch/breakout/internal/pkg/render"
)

type RenderSystem struct {
	width               float32
	height              float32
	entities            []renderEntityHolder
	quadRenderProgram   *render.Program
	circleRenderProgram *render.Program
	projectionMatrix    mgl32.Mat4
	generator           *TriangleIndexBufferGenerator
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

	render.UseDefaultBlending()
	gl.ClearColor(colourDarkNavy.R, colourDarkNavy.G, colourDarkNavy.B, colourDarkNavy.A)

	r.quadRenderProgram = NewQuadShaderProgramOrPanic()
	r.circleRenderProgram = NewCircleShaderProgramOrPanic()
}

func (r *RenderSystem) quadVertexBuffer() []float32 {
	var result []float32

	for _, e := range r.entities {
		if e.renderComponent.Quad != nil {
			result = append(result, e.renderComponent.Quad.ToBuffer()...)
		}
	}

	return result
}

func (r *RenderSystem) Update(float32) {

	// quad buffers
	render.NewVertexArray().
		AddBuffer(
			render.NewVertexBuffer(r.quadVertexBuffer()), render.NewVertexBufferLayout().
				AddLayoutFloats(2).AddLayoutFloats(4))

	gl.Clear(gl.COLOR_BUFFER_BIT)

	r.quadRenderProgram.Bind()
	r.quadRenderProgram.SetUniformMat4f("u_MVP", r.projectionMatrix)
	gl.DrawElements(gl.TRIANGLES, render.NewIndexBuffer(r.generator.Generate(len(r.entities))).GetCount(), gl.UNSIGNED_INT, gl.PtrOffset(0))

	// render circles - crap code right now
	r.circleRenderProgram.Bind()

	for _, e := range r.entities {
		if e.renderComponent.Circle == nil {
			continue
		}
		mvp := r.projectionMatrix.Mul4(mgl32.Ident4().Mul4(mgl32.Translate3D(e.renderComponent.Circle.Position.X(), e.renderComponent.Circle.Position.Y(), 0)))
		r.circleRenderProgram.SetUniformMat4f("u_MVP", mvp)
		r.circleRenderProgram.SetUniformVec4("u_Colour", e.renderComponent.Circle.Colour)

		gl.DrawArrays(gl.TRIANGLE_FAN, 0, render.NewVertexArray().
			AddBuffer(render.NewVertexBuffer(e.renderComponent.Circle.ToBuffer()), render.NewVertexBufferLayout().AddLayoutFloats(2)).GetBufferCount())
	}

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
