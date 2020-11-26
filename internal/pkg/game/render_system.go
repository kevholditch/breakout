package game

import (
	"github.com/EngoEngine/ecs"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/kevholditch/breakout/internal/pkg/render"
)

type RenderSystem struct {
	width             float32
	height            float32
	entities          []renderEntityHolder
	quadRenderProgram *render.Program
	projectionMatrix  mgl32.Mat4
	generator         *TriangleIndexBufferGenerator
	circleRenderer    *CircleRenderer
}

type renderEntityHolder struct {
	entity          *ecs.BasicEntity
	renderComponent *RenderComponent
}

func NewRenderSystem(windowSize WindowSize) *RenderSystem {
	proj := mgl32.Ortho(0, windowSize.Width, 0, windowSize.Height, -1.0, 1.0)
	return &RenderSystem{
		entities:         []renderEntityHolder{},
		width:            windowSize.Width,
		height:           windowSize.Height,
		projectionMatrix: proj,
		generator:        NewTriangleIndexBufferGenerator(),
		circleRenderer:   NewCircleRenderer(proj),
	}
}

func (r *RenderSystem) New(*ecs.World) {

	render.UseDefaultBlending()
	gl.ClearColor(colourDarkNavy.R, colourDarkNavy.G, colourDarkNavy.B, colourDarkNavy.A)

	r.quadRenderProgram = NewQuadShaderProgramOrPanic()
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

	r.circleRenderer.Render()

}

func (r *RenderSystem) Add(entity *ecs.BasicEntity, renderComponent *RenderComponent) {
	r.entities = append(r.entities, renderEntityHolder{
		entity: entity, renderComponent: renderComponent,
	})
	if renderComponent.Circle != nil {
		r.circleRenderer.Add(entity.ID(), renderComponent.Circle)
	}
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
