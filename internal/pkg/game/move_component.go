package game

import "github.com/kevholditch/breakout/internal/pkg/render"

type MoveComponent struct {
	Quad  *render.Quad
	Speed float32
}

func NewMoveComponent(q *render.Quad, speed float32) *MoveComponent {
	return &MoveComponent{Quad: q, Speed: speed}
}
