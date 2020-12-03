package components

import "github.com/kevholditch/breakout/internal/pkg/ecs"

var HasBallPhysics *BallPhysics

type BallPhysics interface {
	BallPhysicsComponent() *BallPhysicsComponent
}

type BallPhysicsComponent struct{}

func NewBallPhysicsComponent() *BallPhysicsComponent {
	return &BallPhysicsComponent{}
}

func (b *BallPhysicsComponent) BallPhysicsComponent() *BallPhysicsComponent {
	return b
}

func init() {
	ecs.RegisterComponent(&BallPhysicsComponent{})
}
