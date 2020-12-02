package game

import (
	"github.com/kevholditch/breakout/internal/pkg/ecs"
	"github.com/kevholditch/breakout/internal/pkg/game/components"
)

type PlayerMovementSystem struct {
	playingSpace PlayingSpace
	entities     []laterallyMovableEntity
}

type laterallyMovableEntity struct {
	base       *ecs.Entity
	speed      *components.SpeedControlComponent
	position   *components.PositionedComponent
	dimensions *components.DimensionComponent
}

func NewPlayerMovementSystem(space PlayingSpace) *PlayerMovementSystem {
	return &PlayerMovementSystem{
		playingSpace: space,
		entities:     []laterallyMovableEntity{},
	}
}

func (m *PlayerMovementSystem) Add(entity *ecs.Entity) {
	m.entities = append(m.entities,
		laterallyMovableEntity{
			base:       entity,
			speed:      entity.Component(components.IsSpeedControllable).(*components.SpeedControlComponent),
			position:   entity.Component(components.IsPositioned).(*components.PositionedComponent),
			dimensions: entity.Component(components.HasDimensions).(*components.DimensionComponent),
		})
}

func (m *PlayerMovementSystem) Update(dt float32) {
	for _, e := range m.entities {
		moveAmount := dt * e.speed.Speed
		if e.position.X+moveAmount < 0 {
			e.position.X = 0
			e.speed.Speed = 0
		} else if e.position.X+moveAmount+e.dimensions.Width > m.playingSpace.Width {
			e.position.X = m.playingSpace.Width - e.dimensions.Width
			e.speed.Speed = 0
		} else {
			e.position.X += moveAmount
		}
	}
}

func (m *PlayerMovementSystem) Remove(basic *ecs.Entity) {
	var del = -1
	for index, e := range m.entities {
		if e.base.ID() == basic.ID() {
			del = index
			break
		}
	}
	if del >= 0 {
		m.entities = append(m.entities[:del], m.entities[del+1:]...)
	}

}

func (m *PlayerMovementSystem) RequiredTypes() []interface{} {
	return []interface{}{
		components.IsSpeedControllable,
		components.IsPositioned,
		components.HasDimensions,
	}
}
