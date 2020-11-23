package game

import (
	"github.com/EngoEngine/ecs"
	"github.com/kevholditch/breakout/internal/pkg/render"
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
	leftMargin := float32(20)
	blockWidth := float32(80)
	blockSpacing := float32(10)
	blockHeight := float32(20)
	topMargin := float32(200)

	for i := 0; i < 11; i++ {
		l.currentBlocks = append(l.currentBlocks,
			&BlockEntity{
				BasicEntity: ecs.NewBasic(),
				RenderComponent: NewRenderComponent(render.NewQuad((float32(i)*(blockWidth+blockSpacing))+leftMargin, l.playingSpace.Height-topMargin, blockWidth, blockHeight,
					0.9, 0.3, 0.7, 1.0)),
			})
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
