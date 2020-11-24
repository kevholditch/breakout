package game

import (
	"github.com/EngoEngine/ecs"
)

type LevelSystem struct {
	world         *ecs.World
	playingSpace  PlayingSpace
	currentBlocks []*BlockEntity
}

type BlockEntity struct {
	ecs.BasicEntity
	RenderComponent *RenderComponent
}

type RenderAdd interface {
	Add(entity *ecs.BasicEntity, renderComponent *RenderComponent)
}

func NewLevelSystem(space PlayingSpace) *LevelSystem {
	return &LevelSystem{
		playingSpace: space,
	}
}

func (l *LevelSystem) New(world *ecs.World) {
	l.world = world
}

func (l *LevelSystem) generateLevel() {

	l.currentBlocks = []*BlockEntity{}
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
			l.currentBlocks = append(l.currentBlocks,
				&BlockEntity{
					BasicEntity: ecs.NewBasic(),
					RenderComponent: NewRenderComponent(NewQuadWithColour((i*(blockWidth+blockSpacing))+sideMargin, y, blockWidth, blockHeight,
						blockColour)),
				})
			alpha -= 0.05
		}
	}

}

func (l *LevelSystem) Update(float32) {
	if len(l.currentBlocks) != 0 {
		return
	}

	l.generateLevel()
	for _, system := range l.world.Systems() {
		v, ok := system.(RenderAdd)
		if ok {
			for _, e := range l.currentBlocks {
				v.Add(&e.BasicEntity, e.RenderComponent)
			}
		}
	}

}

func (l *LevelSystem) Remove(basic ecs.BasicEntity) {

}
