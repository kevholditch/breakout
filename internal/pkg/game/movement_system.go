package game

import "github.com/EngoEngine/ecs"

type MovementSystem struct {
	entities map[uint64]struct {
		ecs.BasicEntity
		moveComponent *MoveComponent
	}
}

func NewMovementSystem() *MovementSystem {

	m := &MovementSystem{entities: map[uint64]struct {
		ecs.BasicEntity
		moveComponent *MoveComponent
	}{}}

	return m
}

func (m *MovementSystem) Add(entity ecs.BasicEntity, moveComponent *MoveComponent) *MovementSystem {
	m.entities[entity.ID()] = struct {
		ecs.BasicEntity
		moveComponent *MoveComponent
	}{
		entity, moveComponent,
	}
	return m
}

func (m *MovementSystem) Update(dt float32) {
	for _, e := range m.entities {
		e.moveComponent.Quad.Move(dt*e.moveComponent.Speed, 0)
	}
}

func (m *MovementSystem) Remove(_ ecs.BasicEntity) {}
