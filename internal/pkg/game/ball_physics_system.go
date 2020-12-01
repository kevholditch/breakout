package game

import (
	"github.com/EngoEngine/ecs"
)

type BallPhysicsSystem struct {
	playerQuad           *Quad
	playerStateComponent *PlayerStateComponent
	ballPhysicsComponent *BallPhysicsComponent
	playingSpace         PlayingSpace
	world                *ecs.World
	entities             []struct {
		*ecs.BasicEntity
		renderComponent *RenderComponent
	}
}

func NewBallPhysicsSystem(playerQuad *Quad, playerStateComponent *PlayerStateComponent, playingSpace PlayingSpace, ballPhysicsComponent *BallPhysicsComponent) *BallPhysicsSystem {
	return &BallPhysicsSystem{
		playerQuad:           playerQuad,
		playerStateComponent: playerStateComponent,
		ballPhysicsComponent: ballPhysicsComponent,
		playingSpace:         playingSpace,
		entities: []struct {
			*ecs.BasicEntity
			renderComponent *RenderComponent
		}{}}
}

func (b *BallPhysicsSystem) New(world *ecs.World) {
	b.world = world
}

func (b *BallPhysicsSystem) Add(entity *ecs.BasicEntity, renderComponent *RenderComponent) {
	b.entities = append(b.entities, struct {
		*ecs.BasicEntity
		renderComponent *RenderComponent
	}{entity, renderComponent})
}

func (b *BallPhysicsSystem) Update(dt float32) {

	switch b.playerStateComponent.State {
	case Kickoff:
		{
			b.ballPhysicsComponent.Circle.Position[0] = b.playerQuad.Position.X() + (b.playerQuad.Width() / 4)
			b.ballPhysicsComponent.Circle.Position[1] = b.playerQuad.Position.W() + 14
		}
	case Playing:
		{
			ballMove := [2]float32{dt * b.ballPhysicsComponent.Speed[0], dt * b.ballPhysicsComponent.Speed[1]}
			b.ballPhysicsComponent.Circle.Position = b.ballPhysicsComponent.Circle.Position.Add(ballMove)

			circle := b.ballPhysicsComponent.Circle

			// if going left then check left side of screen
			if b.ballPhysicsComponent.Speed[0] < 0 && circle.LeftMost() <= 0 {
				b.ballPhysicsComponent.Speed[0] = b.ballPhysicsComponent.Speed[0] * -1
			}
			if b.ballPhysicsComponent.Speed[0] > 0 && circle.RightMost() >= b.playingSpace.Width {
				b.ballPhysicsComponent.Speed[0] = b.ballPhysicsComponent.Speed[0] * -1
			}
			if b.ballPhysicsComponent.Speed[1] > 0 && circle.UpperMost() >= b.playingSpace.Height {
				b.ballPhysicsComponent.Speed[1] = b.ballPhysicsComponent.Speed[1] * -1
			}
			if circle.LowerMost() <= 0 {
				b.playerStateComponent.State = Kickoff
			}

			// check if we hit player if ball is going downwards
			if b.ballPhysicsComponent.Speed[1] < 0 {

				if circle.LowerMost() <= b.playerQuad.Position.W() &&
					circle.LowerMost() >= b.playerQuad.Position.Y() &&
					circle.Position.X() >= b.playerQuad.Position.X() &&
					circle.Position.X() <= b.playerQuad.Position.Z() {
					b.ballPhysicsComponent.Speed[1] = b.ballPhysicsComponent.Speed[1] * -1
				}
			}

			entitiesToRemove := []ecs.BasicEntity{}

			for _, block := range b.entities {
				q := block.renderComponent.Quad
				blockHit := false

				// if ball going down
				if b.ballPhysicsComponent.Speed[1] < 0 {
					if circle.LowerMost() <= q.Position.W() &&
						circle.LowerMost() >= q.Position.Y() &&
						circle.Position.X() >= q.Position.X() &&
						circle.Position.X() <= q.Position.Z() {
						b.ballPhysicsComponent.Speed[1] = b.ballPhysicsComponent.Speed[1] * -1
						blockHit = true
					}
				}
				// if ball going right
				if b.ballPhysicsComponent.Speed[0] > 0 {
					if circle.RightMost() <= q.Position.Z() &&
						circle.RightMost() >= q.Position.X() &&
						circle.Position.Y() >= q.Position.Y() &&
						circle.Position.Y() <= q.Position.W() {
						b.ballPhysicsComponent.Speed[0] = b.ballPhysicsComponent.Speed[0] * -1
						blockHit = true
					}
				}

				// if ball going up
				if b.ballPhysicsComponent.Speed[1] > 0 {
					if circle.UpperMost() <= q.Position.W() &&
						circle.UpperMost() >= q.Position.Y() &&
						circle.Position.X() >= q.Position.X() &&
						circle.Position.X() <= q.Position.Z() {
						b.ballPhysicsComponent.Speed[1] = b.ballPhysicsComponent.Speed[1] * -1
						blockHit = true
					}
				}

				// if ball going left
				if b.ballPhysicsComponent.Speed[0] < 0 {
					if circle.LeftMost() <= q.Position.Z() &&
						circle.LeftMost() >= q.Position.X() &&
						circle.Position.Y() >= q.Position.Y() &&
						circle.Position.Y() <= q.Position.W() {
						b.ballPhysicsComponent.Speed[0] = b.ballPhysicsComponent.Speed[0] * -1
						blockHit = true
					}
				}

				if blockHit {
					entitiesToRemove = append(entitiesToRemove, *block.GetBasicEntity())
				}
			}

			for _, entity := range entitiesToRemove {
				b.world.RemoveEntity(entity)
			}

		}
	}

}

func (b *BallPhysicsSystem) Remove(basic ecs.BasicEntity) {
	var del = -1
	for index, e := range b.entities {
		if e.GetBasicEntity().ID() == basic.ID() {
			del = index
			break
		}
	}
	if del >= 0 {
		b.entities = append(b.entities[:del], b.entities[del+1:]...)
	}
}
