package game

type RenderComponent struct {
	Quad   *Quad
	Circle *Circle
}

func NewRenderComponent(q *Quad) *RenderComponent {
	return &RenderComponent{Quad: q}
}
