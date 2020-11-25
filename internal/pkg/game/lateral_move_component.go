package game

type LateralMoveComponent struct {
	Quad  *Quad
	Speed float32
}

func NewLateralMoveComponent(q *Quad, speed float32) *LateralMoveComponent {
	return &LateralMoveComponent{Quad: q, Speed: speed}
}
