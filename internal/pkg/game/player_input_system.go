package game

import "github.com/EngoEngine/ecs"

type PlayerInputSystem struct {
	PlayerEntity *PlayerEntity
}

func NewPlayerInputSystem(player *PlayerEntity, subscribe func(func(int), func(int))) *PlayerInputSystem {
	p := &PlayerInputSystem{
		PlayerEntity: player,
	}

	inc := float32(1000)

	subscribe(func(key int) {
		switch key {
		case 263:
			if player.MoveComponent.Speed > 0 {
				player.MoveComponent.Speed = 0
			} else {
				player.MoveComponent.Speed -= inc
			}
		case 262:
			if player.MoveComponent.Speed < 0 {
				player.MoveComponent.Speed = 0
			} else {
				player.MoveComponent.Speed += inc
			}
		}
	}, func(key int) {

	})

	return p
}

func (m *PlayerInputSystem) Update(dt float32) {}

func (m *PlayerInputSystem) Remove(_ ecs.BasicEntity) {}
