package game

import "github.com/kevholditch/breakout/internal/pkg/render"

type MoveComponent struct {
	Quad  *render.Quad
	Speed float32
}

func NewMoveComponent(q *render.Quad) *MoveComponent {
	return &MoveComponent{Quad: q, Speed: 0.0}
}
