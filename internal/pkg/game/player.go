package game

import (
	"github.com/kevholditch/breakout/internal/pkg/ecs"
	"github.com/kevholditch/breakout/internal/pkg/game/components"
)

//
//type PlayerEntity struct {
//	ecs.BasicEntity
//	RenderComponent *RenderComponent
//	MoveComponent   *LateralMoveComponent
//	StateComponent  *PlayerStateComponent
//}

func NewPlayer() *ecs.Entity {

	playerQuad := components.NewQuadWithColour(200, 30, colourWhite)

	player := ecs.NewEntity()

	player.Add(components.NewRenderComponent(playerQuad))
	player.Add(components.NewPlayerStateComponent())
	player.Add(NewControlComponent(0))
	player.Add(components.NewPositionedComponent(200, 20))

	return player

	//return &PlayerEntity{
	//	BasicEntity:     ecs.NewBasic(),
	//	RenderComponent: NewRenderComponent(playerQuad),
	//	MoveComponent:   NewLateralMoveComponent(playerQuad, 0.0),
	//	StateComponent:  NewPlayerStateComponent(),
	//}

}
