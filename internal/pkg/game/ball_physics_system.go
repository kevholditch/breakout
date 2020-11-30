package game

import (
	"github.com/EngoEngine/ecs"
)

type BallPhysicsSystem struct {
	playerQuad           *Quad
	playerStateComponent *PlayerStateComponent
	playingSpace         PlayingSpace
	entities             []struct {
		*ecs.BasicEntity
		ballPhysicsComponent *BallPhysicsComponent
	}
}

func NewBallPhysicsSystem(playerQuad *Quad, playerStateComponent *PlayerStateComponent, playingSpace PlayingSpace) *BallPhysicsSystem {
	return &BallPhysicsSystem{
		playerQuad:           playerQuad,
		playerStateComponent: playerStateComponent,
		playingSpace:         playingSpace,
		entities: []struct {
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
		switch b.playerStateComponent.State {
		case Kickoff:
			{
				ball.ballPhysicsComponent.Circle.Position[0] = b.playerQuad.Position.X() + (b.playerQuad.Width() / 4)
				ball.ballPhysicsComponent.Circle.Position[1] = b.playerQuad.Position.W() + 14
			}
		case Playing:
			{
				ballMove := [2]float32{dt * ball.ballPhysicsComponent.Speed[0], dt * ball.ballPhysicsComponent.Speed[1]}
				ball.ballPhysicsComponent.Circle.Position = ball.ballPhysicsComponent.Circle.Position.Add(ballMove)

				circle := ball.ballPhysicsComponent.Circle

				// ball bouncing
				if circle.LeftMost() <= 0 || circle.RightMost() >= b.playingSpace.Width {
					ball.ballPhysicsComponent.Speed[0] = ball.ballPhysicsComponent.Speed[0] * -1
				}
				if circle.UpperMost() >= b.playingSpace.Height {
					ball.ballPhysicsComponent.Speed[1] = ball.ballPhysicsComponent.Speed[1] * -1
				}
				if circle.LowerMost() <= 0 {
					b.playerStateComponent.State = Kickoff
				}

				// check if we hit player if ball is going downwards
				if ball.ballPhysicsComponent.Speed[1] < 0 {

					if circle.LowerMost() <= (b.playerQuad.Position.W()+5) &&
						circle.LowerMost() >= (b.playerQuad.Position.W()-5) &&
						circle.Position.X() >= b.playerQuad.Position.X() &&
						circle.Position.X() <= b.playerQuad.Position.Z() {
						ball.ballPhysicsComponent.Speed[1] = ball.ballPhysicsComponent.Speed[1] * -1
					}
				}
			}
		}
	}

}

func (b *BallPhysicsSystem) Remove(_ ecs.BasicEntity) {}
