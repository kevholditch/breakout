package components

import (
	"github.com/kevholditch/breakout/internal/pkg/game/primitives"
	"github.com/liamg/ecs"
)

type builder struct {
	entity *ecs.Entity
}

func NewEntityBuilder() *builder {
	return &builder{entity: ecs.NewEntity()}
}

func (b *builder) WithBallPhysics() *builder {
	b.entity.Add(NewBallPhysicsComponent())
	return b
}

func (b *builder) IsBlock() *builder {
	b.entity.Add(NewBlockComponent())
	return b
}

func (b *builder) IsQuad() *builder {
	b.entity.Add(NewQuadComponent())
	return b
}

func (b *builder) IsCircle(radius float32) *builder {
	b.entity.Add(NewCircleComponent(radius))
	return b
}

func (b *builder) WithColour(colour primitives.Colour) *builder {
	b.entity.Add(NewColouredComponent(colour))
	return b
}

func (b *builder) MakeControllable() *builder {
	b.entity.Add(NewControlComponent())
	return b
}

func (b *builder) WithDimensions(width, height float32) *builder {
	b.entity.Add(NewDimensionsComponent(width, height))
	return b
}

func (b *builder) WithPosition(x, y float32) *builder {
	b.entity.Add(NewPositionedComponent(x, y))
	return b
}

func (b *builder) WithSpeed(speedX, speedY float32) *builder {
	b.entity.Add(NewSpeedComponent([2]float32{speedX, speedY}))
	return b
}

func (b *builder) WithComponent(component ecs.Component) *builder {
	b.entity.Add(component)
	return b
}

func (b *builder) Build() *ecs.Entity {
	return b.entity
}
