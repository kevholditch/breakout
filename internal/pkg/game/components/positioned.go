package components

import "github.com/kevholditch/breakout/internal/pkg/ecs"

var IsPositioned *Positioned

type Positioned interface {
	PositionedComponent() *PositionedComponent
}

type PositionedComponent struct {
	X, Y float32
}

func NewPositionedComponent(x, y float32) *PositionedComponent {
	return &PositionedComponent{
		X: x,
		Y: y,
	}
}

func (c *PositionedComponent) PositionedComponent() *PositionedComponent {
	return c
}

func init() {
	ecs.RegisterComponent(&PositionedComponent{})
}
