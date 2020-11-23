package game

import "github.com/EngoEngine/ecs"

type PlayerInputSystem struct {
	entities map[uint64]struct {
		*ecs.BasicEntity
		moveComponent *MoveComponent
		increment     float32
	}
	subscribe func(func(int), func(int))
}

func NewPlayerInputSystem(subscribe func(func(int), func(int))) *PlayerInputSystem {

	return &PlayerInputSystem{
		entities: map[uint64]struct {
			*ecs.BasicEntity
			moveComponent *MoveComponent
			increment     float32
		}{},
		subscribe: subscribe,
	}

}

func (m *PlayerInputSystem) Add(entity *ecs.BasicEntity, moveComponent *MoveComponent) *PlayerInputSystem {

	m.entities[entity.ID()] = struct {
		*ecs.BasicEntity
		moveComponent *MoveComponent
		increment     float32
	}{
		entity, moveComponent, 1.0,
	}

	m.subscribe(func(key int) {
		switch key {
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

	return m
}

func (m *PlayerInputSystem) Update(dt float32) {}

func (m *PlayerInputSystem) Remove(_ ecs.BasicEntity) {}
