package game

type MoveComponent struct {
	Quad  *Quad
	Speed float32
}

func NewMoveComponent(q *Quad, speed float32) *MoveComponent {
	return &MoveComponent{Quad: q, Speed: speed}
}
