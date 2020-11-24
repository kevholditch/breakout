package game

type RenderComponent struct {
	Quad *Quad
}

func NewRenderComponent(q *Quad) *RenderComponent {
	return &RenderComponent{Quad: q}
}
