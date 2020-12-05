package game

import (
	"github.com/kevholditch/breakout/internal/pkg/game/components"
	"github.com/liamg/ecs"
)

func NewBall() *ecs.Entity {
	return components.NewEntityBuilder().
		WithDimensions(12, 12).
		WithColour(colourCoral).
		WithPosition(200, 62).
		IsCircle(12).
		WithSpeed(0, 0).
		MakeControllable().
		Build()
}
