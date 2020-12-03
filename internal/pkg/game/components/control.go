package components

import "github.com/kevholditch/breakout/internal/pkg/ecs"

var IsControllable *Controllable

type Controllable interface {
	ControlComponent() *ControlComponent
}

type ControlComponent struct{}

func NewControlComponent() *ControlComponent {
	return &ControlComponent{}
}

func (c *ControlComponent) ControlComponent() *ControlComponent {
	return c
}

func init() {
	ecs.RegisterComponent(&ControlComponent{})
}
