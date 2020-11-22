package game

import (
	"github.com/kevholditch/breakout/internal/pkg/render"
)

type RenderComponent struct {
	Quad *render.Quad
}

func NewRenderComponent(q *render.Quad) *RenderComponent {
	return &RenderComponent{Quad: q}
}
