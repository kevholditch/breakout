package game

import (
	"github.com/kevholditch/breakout/internal/pkg/ecs"
	"github.com/kevholditch/breakout/internal/pkg/game/components"
)

//
//import "github.com/kevholditch/breakout/internal/pkg/ecs"
//
//type BallEntity struct {
//	RenderComponent      *RenderComponent
//	BallPhysicsComponent *BallPhysicsComponent
//}
//
//func NewBall() *ecs.Entity {
//
//	ball := ecs.NewEntity()
//	ball.Add()
//
//	ballCircle := NewCircleWithColour(200, 200, 12, colourCoral)
//
//	return &BallEntity{
//		BasicEntity:          ecs.NewBasic(),
//		RenderComponent:      &RenderComponent{Circle: ballCircle},
//		BallPhysicsComponent: NewBallPhysicsComponent([2]float32{0, 0}, ballCircle),
//	}
//}

func NewBall() *ecs.Entity {

	ball := ecs.NewEntity()

	ball.Add(components.NewDimensionsComponent(12, 12))
	ball.Add(components.NewColouredComponent(colourCoral))
	ball.Add(components.NewPositionedComponent(200, 62))
	ball.Add(components.NewCircleComponent(12))
	ball.Add(components.NewControlComponent([2]float32{0, 0}))

	return ball

}
