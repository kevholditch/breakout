package game

import (
	"github.com/EngoEngine/ecs"
)

type HudEntity struct {
	ecs.BasicEntity
	RenderComponent *RenderComponent
}

type HudParams struct {
	Height       int
	ScreenHeight int
	ScreenWidth  int
}

type WindowDimensions struct {
	Width  int
	Height int
}

func NewHud(height int, dimensions WindowDimensions) *HudEntity {

	renderComponent := NewRenderComponent(NewQuadWithColour(0, float32(dimensions.Height-height), float32(dimensions.Width), 1, colourWhite))
	renderComponent.TextBox = NewTextBox([2]float32{float32(dimensions.Width) - 200, (float32(height) / 4) * 3}, "0000000", colourWhite)

	return &HudEntity{
		BasicEntity:     ecs.NewBasic(),
		RenderComponent: renderComponent,
	}
}
