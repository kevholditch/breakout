package components

import "github.com/liamg/ecs"

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
