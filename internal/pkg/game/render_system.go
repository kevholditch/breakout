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
	projectionMatrix mgl32.Mat4
	circleRenderer   *CircleRenderer
	quadRenderer     *QuadRenderer
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
		circleRenderer:   NewCircleRenderer(proj),
		quadRenderer:     NewQuadRenderer(proj),
	}
}

func (r *RenderSystem) New(*ecs.World) {
	render.UseDefaultBlending()
	gl.ClearColor(colourDarkNavy.R, colourDarkNavy.G, colourDarkNavy.B, colourDarkNavy.A)
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

	gl.Clear(gl.COLOR_BUFFER_BIT)

	r.quadRenderer.Render()
	r.circleRenderer.Render()

}

func (r *RenderSystem) Add(entity *ecs.BasicEntity, renderComponent *RenderComponent) {
	r.entities = append(r.entities, renderEntityHolder{
		entity: entity, renderComponent: renderComponent,
	})
	if renderComponent.Circle != nil {
		r.circleRenderer.Add(entity.ID(), renderComponent.Circle)
	}
	if renderComponent.Quad != nil {
		r.quadRenderer.Add(entity.ID(), renderComponent.Quad)
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
	r.circleRenderer.Remove(basic.ID())
	r.quadRenderer.Remove(basic.ID())
}
