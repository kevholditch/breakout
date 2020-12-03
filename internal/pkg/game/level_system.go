package game

import (
	"github.com/kevholditch/breakout/internal/pkg/game/components"
	"github.com/liamg/ecs"
)

type LevelSystem struct {
	entities []blockEntity
}

type blockEntity struct {
	base       *ecs.Entity
	position   *components.PositionedComponent
	dimensions *components.DimensionComponent
}

func NewLevelSystem() *LevelSystem {
	return &LevelSystem{entities: []blockEntity{}}
}

func (l *LevelSystem) GetBlocks() []blockEntity {
	return l.entities
}

func (l *LevelSystem) Update(float322 float32) {
}

func (l *LevelSystem) Add(entity *ecs.Entity) {
	l.entities = append(l.entities, blockEntity{
		base:       entity,
		position:   entity.Component(components.IsPositioned).(*components.PositionedComponent),
		dimensions: entity.Component(components.HasDimensions).(*components.DimensionComponent),
	})
}

func (l *LevelSystem) Remove(entity *ecs.Entity) {
	var del = -1
	for index, e := range l.entities {
		if e.base.ID() == entity.ID() {
			del = index
			break
		}
	}
	if del >= 0 {
		l.entities = append(l.entities[:del], l.entities[del+1:]...)
	}

}

func (l *LevelSystem) RequiredTypes() []interface{} {
	return []interface{}{
		components.IsPositioned,
		components.HasDimensions,
		components.IsBlock,
	}
}
