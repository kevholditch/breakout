package game

import (
	"github.com/EngoEngine/ecs"
)

type PlayerEntity struct {
	ecs.BasicEntity
	RenderComponent *RenderComponent
	MoveComponent   *LateralMoveComponent
}

func NewPlayer() *PlayerEntity {

	playerQuad := NewQuadWithColour(200, 20, 200, 30, colourWhite)

	return &PlayerEntity{
		BasicEntity:     ecs.NewBasic(),
		RenderComponent: NewRenderComponent(playerQuad),
		MoveComponent:   NewLateralMoveComponent(playerQuad, 0.0),
	}

}
