package game

import "github.com/EngoEngine/ecs"

type PlayerMovementSystem struct {
	width    float32
	height   float32
	entities map[uint64]struct {
		*ecs.BasicEntity
		lateralMoveComponent *LateralMoveComponent
	}
}

func NewPlayerMovementSystem(space PlayingSpace) *PlayerMovementSystem {

	m := &PlayerMovementSystem{
		width:  space.Width,
		height: space.Height,
		entities: map[uint64]struct {
			*ecs.BasicEntity
			lateralMoveComponent *LateralMoveComponent
		}{}}

	return m
}

func (m *PlayerMovementSystem) Add(entity *ecs.BasicEntity, moveComponent *LateralMoveComponent) *PlayerMovementSystem {
	m.entities[entity.ID()] = struct {
		*ecs.BasicEntity
		lateralMoveComponent *LateralMoveComponent
	}{
		entity, moveComponent,
	}
	return m
}

func (m *PlayerMovementSystem) Update(dt float32) {
	for _, e := range m.entities {
		moveAmount := dt * e.lateralMoveComponent.Speed
		if e.lateralMoveComponent.Quad.Position.X()+moveAmount < 0 {
			e.lateralMoveComponent.Quad.Position = [4]float32{0, e.lateralMoveComponent.Quad.Position.Y(), e.lateralMoveComponent.Quad.Width(), e.lateralMoveComponent.Quad.Position.W()}
			e.lateralMoveComponent.Speed = 0
		} else if e.lateralMoveComponent.Quad.Position.Z()+moveAmount > m.width {
			e.lateralMoveComponent.Quad.Position = [4]float32{m.width - e.lateralMoveComponent.Quad.Width(), e.lateralMoveComponent.Quad.Position.Y(), m.width, e.lateralMoveComponent.Quad.Position.W()}
			e.lateralMoveComponent.Speed = 0
		} else {
			e.lateralMoveComponent.Quad.Position[0] += moveAmount
			e.lateralMoveComponent.Quad.Position[2] += moveAmount
		}
	}
}

func (m *PlayerMovementSystem) Remove(_ ecs.BasicEntity) {}
