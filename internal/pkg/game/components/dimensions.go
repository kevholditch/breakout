package components

import "github.com/kevholditch/breakout/internal/pkg/ecs"

var HasDimensions *Dimensions

type Dimensions interface {
	DimensionComponent() *DimensionComponent
}

type DimensionComponent struct {
	Width, Height float32
}

func NewDimensionsComponent(width, height float32) *DimensionComponent {
	return &DimensionComponent{
		Width:  width,
		Height: height,
	}
}

func (d *DimensionComponent) DimensionComponent() *DimensionComponent {
	return d
}

func init() {
	ecs.RegisterComponent(&DimensionComponent{})
}
