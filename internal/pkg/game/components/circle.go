package components

import (
	"github.com/liamg/ecs"
	"math"
)

var IsCircle *Circle

type Circle interface {
	CircleComponent() *CircleComponent
}

type CircleComponent struct {
	Radius float32
	Buffer []float32
}

func NewCircleComponent(radius float32) *CircleComponent {
	triangleAmount := float32(60)
	twicePi := float32(2.0) * math.Pi

	var buffer []float32
	for i := float32(0); i <= triangleAmount; i++ {
		x1 := float32(0.5) + (radius * float32(math.Cos(float64(i*twicePi/triangleAmount))))
		y1 := float32(0.5) + (radius * float32(math.Sin(float64(i*twicePi/triangleAmount))))
		buffer = append(buffer, x1, y1)
	}

	return &CircleComponent{
		Radius: radius,
		Buffer: buffer,
	}
}

func (c *CircleComponent) CircleComponent() *CircleComponent {
	return c
}

func init() {
	ecs.RegisterComponent(&CircleComponent{})
}
