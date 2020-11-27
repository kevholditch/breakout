package game

import (
	"github.com/EngoEngine/ecs"
)

type BallPhysicsSystem struct {
	playerQuad *Quad
	entities   []struct {
		*ecs.BasicEntity
		ballPhysicsComponent *BallPhysicsComponent
	}
}

func NewBallPhysicsSystem(playerQuad *Quad) *BallPhysicsSystem {
	return &BallPhysicsSystem{playerQuad: playerQuad, entities: []struct {
		*ecs.BasicEntity
		ballPhysicsComponent *BallPhysicsComponent
	}{}}
}

func (b *BallPhysicsSystem) Add(entity *ecs.BasicEntity, ballPhysicsComponent *BallPhysicsComponent) *BallPhysicsSystem {
	b.entities = append(b.entities, struct {
		*ecs.BasicEntity
		ballPhysicsComponent *BallPhysicsComponent
	}{entity, ballPhysicsComponent})

	return b
}

func (b *BallPhysicsSystem) Update(dt float32) {
	for _, ball := range b.entities {
		ball.ballPhysicsComponent.Circle.Position[0] = b.playerQuad.Position.X() + (b.playerQuad.Width() / 4)
		ball.ballPhysicsComponent.Circle.Position[1] = b.playerQuad.Position.W() + 14
	}

}

func (b *BallPhysicsSystem) Remove(_ ecs.BasicEntity) {}
