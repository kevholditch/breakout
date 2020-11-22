package game

import "github.com/EngoEngine/ecs"

type MovementSystem struct {
	width    float32
	height   float32
	entities map[uint64]struct {
		ecs.BasicEntity
		moveComponent *MoveComponent
	}
}

func NewMovementSystem(width, height float32) *MovementSystem {

	m := &MovementSystem{
		width:  width,
		height: height,
		entities: map[uint64]struct {
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
		moveAmount := dt * e.moveComponent.Speed
		if e.moveComponent.Quad.Position.X()+moveAmount < 0 {
			e.moveComponent.Quad.Position = [4]float32{0, e.moveComponent.Quad.Position.Y(), e.moveComponent.Quad.Width(), e.moveComponent.Quad.Position.W()}
			e.moveComponent.Speed = 0
		} else if e.moveComponent.Quad.Position.Z()+moveAmount > m.width {
			e.moveComponent.Quad.Position = [4]float32{m.width - e.moveComponent.Quad.Width(), e.moveComponent.Quad.Position.Y(), m.width, e.moveComponent.Quad.Position.W()}
			e.moveComponent.Speed = 0
		} else {
			e.moveComponent.Quad.Position[0] += moveAmount
			e.moveComponent.Quad.Position[2] += moveAmount
		}
	}
}

func (m *MovementSystem) Remove(_ ecs.BasicEntity) {}
