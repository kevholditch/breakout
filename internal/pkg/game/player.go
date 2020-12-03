package game

import (
	"github.com/kevholditch/breakout/internal/pkg/ecs"
	"github.com/kevholditch/breakout/internal/pkg/game/components"
)

func NewPlayer() *ecs.Entity {

	player := ecs.NewEntity()

	player.Add(components.NewDimensionsComponent(200, 30))
	player.Add(components.NewColouredComponent(colourWhite))
	player.Add(components.NewPositionedComponent(200, 20))
	player.Add(components.NewQuadComponent())
	player.Add(components.NewSpeedComponent([2]float32{0, 0}))
	player.Add(components.NewControlComponent())

	return player

}
