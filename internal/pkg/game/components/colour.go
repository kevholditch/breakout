package components

import (
	"github.com/kevholditch/breakout/internal/pkg/game/primitives"
	"github.com/liamg/ecs"
)

var IsColoured *Coloured

type Coloured interface {
	ColouredComponent() *ColouredComponent
}

type ColouredComponent struct {
	Colour primitives.Colour
}

func NewColouredComponent(colour primitives.Colour) *ColouredComponent {
	return &ColouredComponent{
		Colour: colour,
	}
}

func (c *ColouredComponent) ColouredComponent() *ColouredComponent {
	return c
}

func init() {
	ecs.RegisterComponent(&ColouredComponent{})
}
