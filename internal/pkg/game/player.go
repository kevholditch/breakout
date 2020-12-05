package game

import (
	"github.com/kevholditch/breakout/internal/pkg/game/components"
	"github.com/liamg/ecs"
)

func NewPlayer() *ecs.Entity {
	return components.NewEntityBuilder().
		WithDimensions(200, 30).
		WithColour(colourWhite).
		WithPosition(200, 20).
		IsQuad().
		WithSpeed(0, 0).
		MakeControllable().
		Build()
}
