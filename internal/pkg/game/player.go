package game

import (
	"github.com/EngoEngine/ecs"
)

type PlayerEntity struct {
	ecs.BasicEntity
	RenderComponent *RenderComponent
	MoveComponent   *MoveComponent
}

func NewPlayer() *PlayerEntity {

	playerQuad := NewQuadWithColour(200, 20, 200, 30, colourWhite)

	return &PlayerEntity{
		BasicEntity:     ecs.NewBasic(),
		RenderComponent: NewRenderComponent(playerQuad),
		MoveComponent:   NewMoveComponent(playerQuad, 0.0),
	}

}
