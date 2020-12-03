package game

import (
	"github.com/kevholditch/breakout/internal/pkg/ecs"
	"github.com/kevholditch/breakout/internal/pkg/game/components"
)

type BallPhysicsSystem struct {
	playerPosition   *components.PositionedComponent
	playerDimensions *components.DimensionComponent
	playingSpace     PlayingSpace
	world            *ecs.World
	entities         []ballPhysicsEntity
	levelSystem      *LevelSystem
	gameState        *GameState
}

type ballPhysicsEntity struct {
	base     *ecs.Entity
	position *components.PositionedComponent
	circle   *components.CircleComponent
	speed    *components.SpeedComponent
}

func NewBallPhysicsSystem(playerPosition *components.PositionedComponent, playerDimensions *components.DimensionComponent, playingSpace PlayingSpace, levelSystem *LevelSystem, state *GameState) *BallPhysicsSystem {
	return &BallPhysicsSystem{
		playerPosition:   playerPosition,
		playerDimensions: playerDimensions,
		playingSpace:     playingSpace,
		levelSystem:      levelSystem,
		gameState:        state,
		entities:         []ballPhysicsEntity{},
	}
}

func (b *BallPhysicsSystem) New(world *ecs.World) {
	b.world = world
}

func (b *BallPhysicsSystem) Add(entity *ecs.Entity) {
	b.entities = append(b.entities, ballPhysicsEntity{
		base:     entity,
		position: entity.Component(components.IsPositioned).(*components.PositionedComponent),
		circle:   entity.Component(components.IsCircle).(*components.CircleComponent),
		speed:    entity.Component(components.HasSpeed).(*components.SpeedComponent),
	})
}

func (b *BallPhysicsSystem) Update(dt float32) {

	entitiesToRemove := []*ecs.Entity{}

	playerW := b.playerPosition.Y + b.playerDimensions.Height
	playerZ := b.playerPosition.X + b.playerDimensions.Width

	for _, ball := range b.entities {
		ballMove := [2]float32{dt * ball.speed.Speed[0], dt * ball.speed.Speed[1]}
		ball.position.X += ballMove[0]
		ball.position.Y += ballMove[1]

		// if going left then check left side of screen
		if ball.speed.Speed[0] < 0 && (ball.position.X-ball.circle.Radius) <= 0 {
			ball.speed.Speed[0] = ball.speed.Speed[0] * -1
		}
		if ball.speed.Speed[0] > 0 && (ball.position.X+ball.circle.Radius) >= b.playingSpace.Width {
			ball.speed.Speed[0] = ball.speed.Speed[0] * -1
		}
		if ball.speed.Speed[1] > 0 && (ball.position.Y+ball.circle.Radius) >= b.playingSpace.Height {
			ball.speed.Speed[1] = ball.speed.Speed[1] * -1
		}
		if (ball.position.Y - ball.circle.Radius) <= 0 {
			b.gameState.State = Kickoff
			return
		}

		// check if we hit player if ball is going downwards
		if ball.speed.Speed[1] < 0 {
			if (ball.position.Y-ball.circle.Radius) <= playerW &&
				(ball.position.Y-ball.circle.Radius) >= b.playerPosition.Y &&
				ball.position.X >= b.playerPosition.X &&
				ball.position.X <= playerZ {
				ball.speed.Speed[1] = ball.speed.Speed[1] * -1
			}
		}

		for _, block := range b.levelSystem.GetBlocks() {

			blockHit := false

			blockW := block.position.Y + block.dimensions.Height
			blockZ := block.position.X + block.dimensions.Width
			// if ball going down
			if ball.speed.Speed[1] < 0 {
				if (ball.position.Y-ball.circle.Radius) <= blockW &&
					(ball.position.Y-ball.circle.Radius) >= blockZ &&
					ball.position.X >= block.position.X &&
					ball.position.X <= blockZ {
					ball.speed.Speed[1] = ball.speed.Speed[1] * -1
					blockHit = true
				}
			}
			// if ball going right
			if ball.speed.Speed[0] > 0 {
				if (ball.position.X+ball.circle.Radius) <= blockZ &&
					(ball.position.X+ball.circle.Radius) >= block.position.X &&
					ball.position.Y >= blockZ &&
					ball.position.Y <= blockW {
					ball.speed.Speed[0] = ball.speed.Speed[0] * -1
					blockHit = true
				}
			}

			// if ball going up
			if ball.speed.Speed[1] > 0 {
				if (ball.position.Y+ball.circle.Radius) <= blockW &&
					(ball.position.Y+ball.circle.Radius) >= blockZ &&
					ball.position.X >= block.position.X &&
					ball.position.X <= blockZ {
					ball.speed.Speed[1] = ball.speed.Speed[1] * -1
					blockHit = true
				}
			}

			// if ball going left
			if ball.speed.Speed[0] < 0 {
				if (ball.position.X-ball.circle.Radius) <= blockZ &&
					(ball.position.X-ball.circle.Radius) >= block.position.X &&
					ball.position.Y >= blockZ &&
					ball.position.Y <= blockW {
					ball.speed.Speed[0] = ball.speed.Speed[0] * -1
					blockHit = true
				}
			}

			if blockHit {
				entitiesToRemove = append(entitiesToRemove, block.base)
			}
		}
	}

	for _, entity := range entitiesToRemove {
		b.world.RemoveEntity(entity)
	}

}

func (b *BallPhysicsSystem) Remove(entity *ecs.Entity) {
	var del = -1
	for index, e := range b.entities {
		if e.base.ID() == entity.ID() {
			del = index
			break
		}
	}
	if del >= 0 {
		b.entities = append(b.entities[:del], b.entities[del+1:]...)
	}

}

func (b *BallPhysicsSystem) RequiredTypes() []interface{} {
	return []interface{}{
		components.IsPositioned,
		components.HasBallPhysics,
		components.HasDimensions,
		components.HasSpeed,
		components.IsCircle,
	}
}
