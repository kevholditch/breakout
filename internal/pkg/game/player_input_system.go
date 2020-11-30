package game

import (
	"github.com/EngoEngine/ecs"
)

type PlayerInputSystem struct {
	playerMoveComponent  *LateralMoveComponent
	playerStateComponent *PlayerStateComponent
	ballPhysicsComponent *BallPhysicsComponent
	subscribe            func(func(int), func(int))
}

func NewPlayerInputSystem(subscribe func(func(int), func(int)), moveComponent *LateralMoveComponent,
	stateComponent *PlayerStateComponent, ballPhysicsComponent *BallPhysicsComponent) *PlayerInputSystem {
	p := &PlayerInputSystem{
		subscribe:            subscribe,
		playerMoveComponent:  moveComponent,
		playerStateComponent: stateComponent,
		ballPhysicsComponent: ballPhysicsComponent,
	}

	p.subscribe(func(key int) {
		switch key {
		case 32:
			if stateComponent.State == Kickoff {
				stateComponent.State = Playing
				ballPhysicsComponent.Speed = [2]float32{0.5, 0.5}
			}
		case 263:
			if moveComponent.Speed > 0 {
				moveComponent.Speed = 0
			} else {
				moveComponent.Speed -= 1.0
			}
		case 262:
			if moveComponent.Speed < 0 {
				moveComponent.Speed = 0
			} else {
				moveComponent.Speed += 1.0
			}
		}
	}, func(key int) {

	})

	return p
}

func (m *PlayerInputSystem) Update(dt float32) {}

func (m *PlayerInputSystem) Remove(_ ecs.BasicEntity) {}
