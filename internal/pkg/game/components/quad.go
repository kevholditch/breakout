package components

import "github.com/kevholditch/breakout/internal/pkg/ecs"

var IsQuad *Quad

type Quad interface {
	QuadComponent() *QuadComponent
}

type QuadComponent struct{}

func NewQuadComponent() *QuadComponent {
	return &QuadComponent{}
}

func (q *QuadComponent) QuadComponent() *QuadComponent {
	return q
}

func init() {
	ecs.RegisterComponent(&QuadComponent{})
}
