package game

import (
	"github.com/EngoEngine/ecs"
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/kevholditch/breakout/internal/pkg/render"
)

type Renderer interface {
	Render()
	Remove(id uint64)
}

type RenderSystem struct {
	width          float32
	height         float32
	entities       []renderEntityHolder
	circleRenderer *CircleRenderer
	quadRenderer   *QuadRenderer
	renderers      []Renderer
}

type renderEntityHolder struct {
	entity          *ecs.BasicEntity
	renderComponent *RenderComponent
}

func NewRenderSystem(windowSize WindowSize) *RenderSystem {
	proj := mgl32.Ortho(0, windowSize.Width, 0, windowSize.Height, -1.0, 1.0)
	circleRenderer := NewCircleRenderer(proj)
	quadRenderer := NewQuadRenderer(proj)

	return &RenderSystem{
		entities:       []renderEntityHolder{},
		width:          windowSize.Width,
		height:         windowSize.Height,
		circleRenderer: circleRenderer,
		quadRenderer:   quadRenderer,
		renderers:      []Renderer{circleRenderer, quadRenderer},
	}
}

func (r *RenderSystem) New(*ecs.World) {
	render.UseDefaultBlending()
	gl.ClearColor(colourDarkNavy.R, colourDarkNavy.G, colourDarkNavy.B, colourDarkNavy.A)
}

func (r *RenderSystem) Update(float32) {
	gl.Clear(gl.COLOR_BUFFER_BIT)

	for _, renderer := range r.renderers {
		renderer.Render()
	}
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
	for _, renderer := range r.renderers {
		renderer.Remove(basic.ID())
	}
}
