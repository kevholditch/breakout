package components

import (
	"github.com/kevholditch/breakout/internal/pkg/ecs"
	"github.com/kevholditch/breakout/internal/pkg/game/primitives"
)

var IsRenderable *Renderable

type Renderable interface {
	RenderableComponent() *RenderComponent
}

type RenderComponent struct {
	Quad    *primitives.Quad
	Circle  *primitives.Circle
	TextBox *primitives.TextBox
}

func NewRenderComponent(q *primitives.Quad) *RenderComponent {
	return &RenderComponent{Quad: q}
}

func (r *RenderComponent) RenderableComponent() *RenderComponent {
	return r
}

func init() {
	ecs.RegisterComponent(&RenderComponent{})
}
