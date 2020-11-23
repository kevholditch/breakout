package game

import (
	"github.com/EngoEngine/ecs"
	"github.com/kevholditch/breakout/internal/pkg/render"
)

type PlayerEntity struct {
	ecs.BasicEntity
	RenderComponent *RenderComponent
	MoveComponent   *MoveComponent
}

func NewPlayer() *PlayerEntity {

	playerQuad := render.NewQuad(200, 20, 200, 30, 0.7, 0.8, 0.2, 1.0)

	return &PlayerEntity{
		BasicEntity:     ecs.NewBasic(),
		RenderComponent: NewRenderComponent(playerQuad),
		MoveComponent:   NewMoveComponent(playerQuad, 0.0),
	}

}
