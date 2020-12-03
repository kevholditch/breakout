package components

import "github.com/liamg/ecs"

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
