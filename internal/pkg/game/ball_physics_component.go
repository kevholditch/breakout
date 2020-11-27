package game

import "github.com/go-gl/mathgl/mgl32"

type BallPhysicsComponent struct {
	Speed  mgl32.Vec2
	Circle *Circle
}

func NewBallPhysicsComponent(speed mgl32.Vec2, circle *Circle) *BallPhysicsComponent {
	return &BallPhysicsComponent{
		Speed:  speed,
		Circle: circle,
	}
}
