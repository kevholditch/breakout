package components

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/kevholditch/breakout/internal/pkg/ecs"
)

var HasSpeed *Speed

type Speed interface {
	SpeedComponent() *SpeedComponent
}

type SpeedComponent struct {
	Speed mgl32.Vec2
}

func NewSpeedComponent(initialSpeed mgl32.Vec2) *SpeedComponent {
	return &SpeedComponent{Speed: initialSpeed}
}

func (c *SpeedComponent) SpeedComponent() *SpeedComponent {
	return c
}

func init() {
	ecs.RegisterComponent(&SpeedComponent{})
}
