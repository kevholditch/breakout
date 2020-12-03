package components

import "github.com/go-gl/mathgl/mgl32"

var IsSpeedControllable *SpeedControllable

type SpeedControllable interface {
	SpeedControllableComponent() *SpeedControlComponent
}

type SpeedControlComponent struct {
	Speed mgl32.Vec2
}

func NewControlComponent(initialSpeed mgl32.Vec2) *SpeedControlComponent {
	return &SpeedControlComponent{Speed: initialSpeed}
}

func (c *SpeedControlComponent) SpeedControllableComponent() *SpeedControlComponent {
	return c
}
