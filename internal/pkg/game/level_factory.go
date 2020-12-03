package game

import (
	"github.com/kevholditch/breakout/internal/pkg/ecs"
	"github.com/kevholditch/breakout/internal/pkg/game/components"
	"github.com/kevholditch/breakout/internal/pkg/game/primitives"
)

type LevelFactory struct {
	playingSpace PlayingSpace
	blocks       []*ecs.Entity
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

func NewLevelFactory(space PlayingSpace) *LevelFactory {
	return &LevelFactory{
		playingSpace: space,
	}
}

func (l *LevelFactory) NewLevel() []*ecs.Entity {

	blocks := []*ecs.Entity{}
	blockWidth := float32(80)

	blockHeight := float32(20)
	topMargin := float32(100)
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
			blocks = append(blocks, NewBlockEntity((i*(blockWidth+blockSpacing))+sideMargin, y, blockWidth, blockHeight, blockColour))
			alpha -= 0.05
		}
	}

	return blocks
}
