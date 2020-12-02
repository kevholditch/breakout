package game

import (
	"github.com/kevholditch/breakout/internal/pkg/ecs"
	"github.com/kevholditch/breakout/internal/pkg/game/components"
)

type PlayerInputSystem struct {
	subscribe func(func(int), func(int))
	entities  []controllableEntity
}

type controllableEntity struct {
	base                 *ecs.Entity
	controlComponent     *SpeedControlComponent
	playerStateComponent *components.PlayerStateComponent
}

func NewPlayerInputSystem(subscribe func(func(int), func(int))) *PlayerInputSystem {
	p := &PlayerInputSystem{
		subscribe: subscribe,
		entities:  []controllableEntity{},
	}

	return p
}

func (m *PlayerInputSystem) Add(entity *ecs.Entity) {

	controlComponent := entity.Component(IsSpeedControllable).(*SpeedControlComponent)
	playerStateComponent := entity.Component(components.HasPlayingState).(*components.PlayerStateComponent)
	m.entities = append(m.entities, controllableEntity{
		base:                 entity,
		controlComponent:     controlComponent,
		playerStateComponent: playerStateComponent,
	})
	m.subscribe(func(key int) {
		switch key {
		case 32:
			if playerStateComponent.State == components.Kickoff {
				playerStateComponent.State = components.Playing
				//ballPhysicsComponent.Speed = [2]float32{0.5, 0.5}
			}
		case 263:
			if controlComponent.Speed > 0 {
				controlComponent.Speed = 0
			} else {
				controlComponent.Speed -= 1.0
			}
		case 262:
			if controlComponent.Speed < 0 {
				controlComponent.Speed = 0
			} else {
				controlComponent.Speed += 1.0
			}
		}
	}, func(key int) {

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
	return []interface{}{IsSpeedControllable, components.HasPlayingState}
}
