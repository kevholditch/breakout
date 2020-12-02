package components

import (
	"github.com/kevholditch/breakout/internal/pkg/ecs"
)

var IsRenderable *Renderable

type Renderable interface {
	RenderableComponent() *RenderComponent
}

type RenderComponent struct {
	Quad    *Quad
	Circle  *Circle
	TextBox *TextBox
}

func NewRenderComponent(q *Quad) *RenderComponent {
	return &RenderComponent{Quad: q}
}

func (r *RenderComponent) RenderableComponent() *RenderComponent {
	return r
}

func init() {
	ecs.RegisterComponent(&RenderComponent{})
}
