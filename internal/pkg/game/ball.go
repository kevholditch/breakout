package game

import (
	"github.com/kevholditch/breakout/internal/pkg/game/components"
	"github.com/liamg/ecs"
)

func NewBall() *ecs.Entity {

	ball := ecs.NewEntity()

	ball.Add(components.NewDimensionsComponent(12, 12))
	ball.Add(components.NewColouredComponent(colourCoral))
	ball.Add(components.NewPositionedComponent(200, 62))
	ball.Add(components.NewCircleComponent(12))
	ball.Add(components.NewSpeedComponent([2]float32{0, 0}))
	ball.Add(components.NewControlComponent())

	return ball
}
