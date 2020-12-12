package game

import (
	"github.com/kevholditch/breakout/internal/pkg/game/components"
	"github.com/kevholditch/breakout/internal/pkg/game/primitives"
	"github.com/liamg/ecs"
)

type LevelSystem struct {
	playingSpace PlayingSpace
	entities     []blockEntity
	world        *ecs.World
}

type blockEntity struct {
	base       *ecs.Entity
	position   *components.PositionedComponent
	dimensions *components.DimensionComponent
}

func NewBlockEntity(x, y, w, h float32, colour primitives.Colour) *ecs.Entity {

	block := ecs.NewEntity()

	block.Add(components.NewDimensionsComponent(w, h))
	block.Add(components.NewColouredComponent(colour))
	block.Add(components.NewPositionedComponent(x, y))
	block.Add(components.NewQuadComponent())
	block.Add(components.NewBlockComponent())

	return block
}

func NewLevelSystem(space PlayingSpace) *LevelSystem {
	return &LevelSystem{entities: []blockEntity{}, playingSpace: space}
}

func (l *LevelSystem) New(world *ecs.World) {
	l.world = world
}

func (l *LevelSystem) GetBlocks() []blockEntity {
	return l.entities
}

func (l *LevelSystem) Update(float322 float32) {
	if len(l.entities) == 0 {
		blockWidth := float32(80)

		blockHeight := float32(20)
		topMargin := float32(40)
		sideMargin := float32(10)
		blockMargin := float32(10)

		blocksInARow := float32(11)
		spacesInARow := blocksInARow - 1
		blockRowLength := blocksInARow * blockWidth
		blockSpacing := (l.playingSpace.Width - ((sideMargin * 2) + blockRowLength)) / spacesInARow
		numberOfRows := float32(4)

		for j := float32(0); j < numberOfRows; j++ {
			y := (l.playingSpace.Height - topMargin) - (j * (blockHeight + blockMargin))
			alpha := float32(1)
			for i := float32(0); i < blocksInARow; i++ {
				blockColour := colourTeal
				blockColour.A = alpha

				l.world.AddEntity(NewBlockEntity((i*(blockWidth+blockSpacing))+sideMargin, y, blockWidth, blockHeight, blockColour))
				alpha -= 0.05
			}
		}

	}
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
