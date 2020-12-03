package game

import (
	"github.com/kevholditch/breakout/internal/pkg/ecs"
	"github.com/kevholditch/breakout/internal/pkg/game/components"
)

type PlayerInputSystem struct {
	subscribe func(func(int))
	entities  []controllableEntity
}

type controllableEntity struct {
	base             *ecs.Entity
	controlComponent *components.SpeedControlComponent
}

func NewPlayerInputSystem(subscribe func(func(int))) *PlayerInputSystem {
	p := &PlayerInputSystem{
		subscribe: subscribe,
		entities:  []controllableEntity{},
	}
	p.subscribe(p.handleKeyPress)

	return p
}

func (m *PlayerInputSystem) handleKeyPress(key int) {

	for _, e := range m.entities {
		switch key {
		//case 32: space

		case 263:
			if e.controlComponent.Speed[0] > 0 {
				e.controlComponent.Speed[0] = 0
			} else {
				e.controlComponent.Speed[0] -= 1.0
			}
		case 262:
			if e.controlComponent.Speed[0] < 0 {
				e.controlComponent.Speed[0] = 0
			} else {
				e.controlComponent.Speed[0] += 1.0
			}
		}
	}

}

func (m *PlayerInputSystem) Add(entity *ecs.Entity) {

	controlComponent := entity.Component(components.IsSpeedControllable).(*components.SpeedControlComponent)
	m.entities = append(m.entities, controllableEntity{
		base:             entity,
		controlComponent: controlComponent,
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
	return []interface{}{components.IsSpeedControllable}
}
