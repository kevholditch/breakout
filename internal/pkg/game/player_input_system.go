package game

import (
	"fmt"
	"github.com/kevholditch/breakout/internal/pkg/ecs"
	"github.com/kevholditch/breakout/internal/pkg/game/components"
)

type PlayerInputSystem struct {
	subscribe func(func(int))
	gameState *GameState
	entities  []controllableEntity
	world     *ecs.World
}

type controllableEntity struct {
	base         *ecs.Entity
	speed        *components.SpeedComponent
	controllable *components.ControlComponent
}

func NewPlayerInputSystem(subscribe func(func(int)), state *GameState) *PlayerInputSystem {
	p := &PlayerInputSystem{
		subscribe: subscribe,
		gameState: state,
		entities:  []controllableEntity{},
	}
	p.subscribe(p.handleKeyPress)

	return p
}

func (m *PlayerInputSystem) New(world *ecs.World) {
	m.world = world
}

func (m *PlayerInputSystem) handleKeyPress(key int) {

	for _, e := range m.entities {

		switch key {
		case 263:
			if e.speed.Speed[0] > 0 {
				e.speed.Speed[0] = 0
			} else {
				e.speed.Speed[0] -= 1.0
			}
		case 262:
			if e.speed.Speed[0] < 0 {
				e.speed.Speed[0] = 0
			} else {
				e.speed.Speed[0] += 1.0
			}
		}
	}

	if key == 32 && m.gameState.State == Kickoff {
		fmt.Printf("in here\n")
		m.gameState.State = Playing
		var ballEntities []struct {
			base         *ecs.Entity
			controllable *components.ControlComponent
		}

		for _, e := range m.entities {
			if e.base.Component(components.IsCircle) != nil {
				e.speed.Speed[0] = 0.5
				e.speed.Speed[1] = 0.5
				ballEntities = append(ballEntities, struct {
					base         *ecs.Entity
					controllable *components.ControlComponent
				}{
					base:         e.base,
					controllable: e.controllable,
				})
			}
		}
		for _, ballEntity := range ballEntities {
			m.world.RemoveComponentFromEntity(ballEntity.controllable, ballEntity.base)
			m.world.AddComponentToEntity(components.NewBallPhysicsComponent(), ballEntity.base)
		}
	}
}

func (m *PlayerInputSystem) Add(entity *ecs.Entity) {

	controlComponent := entity.Component(components.HasSpeed).(*components.SpeedComponent)
	m.entities = append(m.entities, controllableEntity{
		base:         entity,
		speed:        controlComponent,
		controllable: entity.Component(components.IsControllable).(*components.ControlComponent),
	})

}

func (m *PlayerInputSystem) Update(elapsed float32) {}

func (m *PlayerInputSystem) Remove(entity *ecs.Entity) {
	var del = -1
	for index, e := range m.entities {
		if e.base.ID() == entity.ID() {
			del = index
			break
		}
	}
	if del >= 0 {
		m.entities = append(m.entities[:del], m.entities[del+1:]...)
	}
}

func (m *PlayerInputSystem) RequiredTypes() []interface{} {
	return []interface{}{
		components.HasSpeed,
		components.IsControllable,
	}
}
