package game

type RenderComponent struct {
	Quad    *Quad
	Circle  *Circle
	TextBox *TextBox
}

func NewRenderComponent(q *Quad) *RenderComponent {
	return &RenderComponent{Quad: q}
}
