package game

import (
	"github.com/kevholditch/breakout/internal/pkg/ecs"
	"github.com/kevholditch/breakout/internal/pkg/game/components"
)

type LateralMovementSystem struct {
	playingSpace PlayingSpace
	maxWidth     float32
	entities     []laterallyMovableEntity
}

type laterallyMovableEntity struct {
	base       *ecs.Entity
	speed      *components.SpeedComponent
	position   *components.PositionedComponent
	dimensions *components.DimensionComponent
}

func NewLateralMovementSystem(space PlayingSpace) *LateralMovementSystem {
	return &LateralMovementSystem{
		maxWidth:     0,
		playingSpace: space,
		entities:     []laterallyMovableEntity{},
	}
}

func (m *LateralMovementSystem) Add(entity *ecs.Entity) {
	dimensions := entity.Component(components.HasDimensions).(*components.DimensionComponent)
	m.entities = append(m.entities,
		laterallyMovableEntity{
			base:       entity,
			speed:      entity.Component(components.HasSpeed).(*components.SpeedComponent),
			position:   entity.Component(components.IsPositioned).(*components.PositionedComponent),
			dimensions: dimensions,
		})

	if dimensions.Width >= m.maxWidth {
		m.maxWidth = dimensions.Width
	}

}

func (m *LateralMovementSystem) Update(dt float32) {
	for _, e := range m.entities {
		moveAmount := e.speed.Speed[0] * dt
		if e.position.X+moveAmount < 0 {
			e.position.X = 0
			e.speed.Speed[0] = 0
		} else if e.position.X+moveAmount+m.maxWidth > m.playingSpace.Width {
			e.position.X = m.playingSpace.Width - m.maxWidth
			e.speed.Speed[0] = 0
		} else {
			e.position.X += moveAmount
		}
	}
}

func (m *LateralMovementSystem) Remove(basic *ecs.Entity) {
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

	m.maxWidth = 0
	for _, e := range m.entities {
		if e.dimensions.Width >= m.maxWidth {
			m.maxWidth = e.dimensions.Width
		}
	}

}

func (m *LateralMovementSystem) RequiredTypes() []interface{} {
	return []interface{}{
		components.HasSpeed,
		components.IsPositioned,
		components.HasDimensions,
		components.IsControllable,
	}
}
