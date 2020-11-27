package game

import "github.com/EngoEngine/ecs"

type BallEntity struct {
	ecs.BasicEntity
	RenderComponent      *RenderComponent
	BallPhysicsComponent *BallPhysicsComponent
}

func NewBall() *BallEntity {

	ballCircle := NewCircleWithColour(200, 200, 12, colourCoral)

	return &BallEntity{
		BasicEntity:          ecs.NewBasic(),
		RenderComponent:      &RenderComponent{Circle: ballCircle},
		BallPhysicsComponent: NewBallPhysicsComponent([2]float32{0, 0}, ballCircle),
	}
}
